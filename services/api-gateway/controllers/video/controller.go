package video

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"

	"github.com/jackjohn7/webview/api-gateway/pb"
)

type VideoController struct {
	mux *echo.Echo
	videoProcessingService *pb.VideoProcessingServiceClient
}

func NewVideoController(mux *echo.Echo, videoService *pb.VideoProcessingServiceClient) *VideoController {
	c := &VideoController{
		mux: mux,
		videoProcessingService: videoService,
	}
	c.RegisterRoutes()
	return c
}

func (c *VideoController) RegisterRoutes() {
	c.mux.POST("/upload-video", func(ctx echo.Context) error {
		redirectLink := ctx.QueryParam("redirectTo")
		if redirectLink == "" {
			return ctx.Redirect(http.StatusBadRequest, "/")
		}

		fileAttachment, err := ctx.FormFile("video-attachment")
		if err != nil {
			return ctx.Redirect(http.StatusBadRequest, redirectLink)
		}

		fmt.Printf("file with size: %d", fileAttachment.Size);

		video, err := fileAttachment.Open()
		if err != nil {
			return ctx.Redirect(http.StatusBadRequest, redirectLink)
		}
		//bytes := make([]byte, fileAttachment.Size)
		//video.Read(bytes)

		videoId := uuid.New().String()

		service := *c.videoProcessingService

		stream, err := service.ProcessNewVideo(context.Background())
		if err != nil {
			fmt.Println(err)
			return ctx.Redirect(303, redirectLink)
		}

		buffer := make([]byte, 1024 * 2)
		for {
			n, err := video.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				return fmt.Errorf("Could not read from file: %v", err)
			}

			chunk := &pb.ProcessVideoRequest{
				Data: buffer[:n],
				VideoId: videoId,
			}

			if err := stream.Send(chunk); err != nil {
				return fmt.Errorf("Could not send chunk: %v", err)
			}
		}

		responseData, err := stream.CloseAndRecv()
		if err != nil {
			if grpcStatus, ok := status.FromError(err); ok {
				return fmt.Errorf("upload failed: %v", grpcStatus.Message())
			}
			return fmt.Errorf("Could not receive response: %v", err)
		}

		redirectLink = fmt.Sprintf("%s?videoId=%s&videoThumbnail=%s", redirectLink, videoId, responseData.ThumbnailId)

		return ctx.Redirect(303, redirectLink)
	})
}
