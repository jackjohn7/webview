package video

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

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

		video, err := fileAttachment.Open()
		if err != nil {
			return ctx.Redirect(http.StatusBadRequest, redirectLink)
		}
		bytes := make([]byte, fileAttachment.Size)
		video.Read(bytes)

		videoId := uuid.New().String()

		fmt.Println("MADE IT HERE")
		service := *c.videoProcessingService

		fmt.Println("MADE IT HERE TOO")
		responseData, err := service.ProcessNewVideo(context.Background(), &pb.ProcessVideoRequest{
			Data: bytes,
			VideoId: videoId,
		})
		if err != nil {
			fmt.Println("Somethign errored in the api call")
			fmt.Println(err)
			return ctx.Redirect(303, redirectLink)
		}

		redirectLink = fmt.Sprintf("%s?videoId=%s&videoThumbnail=%s", redirectLink, videoId, responseData.ThumbnailId)

		return ctx.Redirect(303, redirectLink)
	})
}
