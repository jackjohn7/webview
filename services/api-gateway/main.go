package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jackjohn7/webview/api-gateway/controllers/video"
	"github.com/jackjohn7/webview/api-gateway/pb"
)

func main() {
	fmt.Println("Starting API-Gateway")
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//AllowOrigins: []string{"http://localhost:3000", "http://localhost:5173"},
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	addr := os.Getenv("VIDEO_SERVICE_ADDR")
	if addr == "" {
		addr = "localhost:3001"
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	videoProcessingClient := pb.NewVideoProcessingServiceClient(conn)
	videoStreamingClient := pb.NewVideoStreamingServiceClient(conn)

	a, err := videoStreamingClient.GetRecentVideos(context.Background(), &pb.RecentVideosRequest{Offset: 0, Range: 4})
	if err != nil {
		log.Fatalln(err)
	}

	x, err := videoProcessingClient.HealthCheck(context.Background(), &pb.HealthCheckRequest{Msg:"Ping"})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", x.Msg)

	for {
		vid, err := a.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetRecentVideos(_) = _, %v", videoStreamingClient, err)
		}
		log.Println(vid)
	}

	video.NewVideoController(app, &videoProcessingClient)

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the gateway! You probably shouldn't be here.")
	})

	app.Start(":3000")
}
