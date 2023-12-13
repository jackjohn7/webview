package video

import (
	"net/http"

	"github.com/jackjohn7/webview/api-gateway/pb"
)

type VideoController struct {
	mux *http.ServeMux
	videoService *pb.VideoProcessingServiceClient
}

func NewVideoController(mux *http.ServeMux, videoService *pb.VideoProcessingServiceClient) *VideoController {
	return &VideoController{
		mux: mux,
		videoService: videoService,
	}
}

func (c *VideoController) RegisterRoutes() {

}
