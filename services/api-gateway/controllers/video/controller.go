package video

import (
	"fmt"
	"net/http"

	"github.com/jackjohn7/webview/api-gateway/pb"
	"github.com/labstack/echo/v4"
)

type VideoController struct {
	mux *echo.Echo
	videoService *pb.VideoProcessingServiceClient
}

func NewVideoController(mux *echo.Echo, videoService *pb.VideoProcessingServiceClient) *VideoController {
	c := &VideoController{
		mux: mux,
		videoService: videoService,
	}
	c.RegisterRoutes()
	return c
}

func (c *VideoController) RegisterRoutes() {
	c.mux.POST("/upload-video", func(c echo.Context) error {
		fmt.Println("route hit!")
		redirectLink := c.QueryParam("redirectTo")
		if redirectLink == "" {
			return c.Redirect(http.StatusBadRequest, "/")
		}

		form, err := c.MultipartForm()
		if err != nil {
			return c.Redirect(http.StatusBadRequest, redirectLink)
		}

		video := form.File["video-attachment"]

		fmt.Println("we here fr")

		if len(video) != 1 {
			return c.Redirect(http.StatusBadRequest, redirectLink)
		}

		fmt.Println("bro got the files")

		fmt.Println(redirectLink)

		return c.Redirect(200, "/")
	})
}
