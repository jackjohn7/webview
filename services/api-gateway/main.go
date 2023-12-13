package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jackjohn7/webview/api-gateway/controllers/video"
	"github.com/jackjohn7/webview/api-gateway/pb"
)

func main() {
	fmt.Println("Starting API-Gateway")
	mux := http.NewServeMux()

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

	a, err := videoStreamingClient.GetRecentVideos(context.Background(), &pb.RecentVideosRequest{Offset: 0, Range: 10})
	if err != nil {
		log.Fatalln(err)
	}

	x, err := videoProcessingClient.HealthCheck(context.Background(), &pb.HealthCheckRequest{Msg:"Ping"})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", x.Msg)
	fmt.Println(a)

	video.NewVideoController(mux, &videoProcessingClient)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome to the gateway! You probably shouldn't be here.")
	})

	http.ListenAndServe("0.0.0.0:3000", mux)
}
