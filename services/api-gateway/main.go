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

	f := pb.NewVideoProcessingServiceClient(conn)

	x, err := f.HealthCheck(context.Background(), &pb.HealthCheckRequest{Msg:"Ping"})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", x.Msg)

	video.NewVideoController(mux, &f)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome to the gateway! You probably shouldn't be here.")
	})

	http.ListenAndServe("0.0.0.0:3000", mux)
}
