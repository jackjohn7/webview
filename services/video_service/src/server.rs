use tokio::sync::mpsc;
use tonic::{Request, Response, Status};
use tokio_stream::{wrappers::ReceiverStream};
use crate::video_processing::{
    video_processing_service_server::VideoProcessingService,
    ProcessVideoRequest,
    ProcessedVideoData,
    DeleteVideoRequest,
    DeleteVideoResponse,
    HealthCheckRequest,
    HealthCheckResponse
};

#[derive(Default)]
pub struct MyVideoProcessor {}

#[tonic::async_trait]
impl VideoProcessingService for MyVideoProcessor {
    async fn process_new_video(&self, _request: Request<ProcessVideoRequest>)
        -> Result<Response<ProcessedVideoData>, Status> {
        Ok(Response::new(ProcessedVideoData { thumbnail_id: "temp".to_owned(), }))
    }
    async fn delete_video(&self, _request: Request<DeleteVideoRequest>)
        -> Result<Response<DeleteVideoResponse>, Status> {
        Ok(Response::new(DeleteVideoResponse { deleted: true }))
    }
    async fn health_check(&self, _request: Request<HealthCheckRequest>)
        -> Result<Response<HealthCheckResponse>, Status> {
        Ok(Response::new(HealthCheckResponse { msg: "pong".to_owned() }))
    }
}

use crate::video_streaming::{VideoChunkRequest, VideoChunk, RecentVideos, RecentVideosRequest};
use crate::video_streaming::video_streaming_service_server::VideoStreamingService;

#[derive(Default)]
pub struct MyVideoCDN {}

#[tonic::async_trait]
impl VideoStreamingService for MyVideoCDN {

    type GetRecentVideosStream = ReceiverStream<Result<RecentVideos, Status>>;

    async fn stream_video_chunk(&self, request: Request<VideoChunkRequest>) -> Result<Response<VideoChunk>,Status> {
        let req = request.into_inner();
        Ok(Response::new(VideoChunk { data: vec![3, 2], chunk_idx: req.chunk_idx, next_chunk_idx: req.chunk_idx + 1 }))
    }

    async fn get_recent_videos(&self, request: Request<RecentVideosRequest>) ->
        std::result::Result<Response<Self::GetRecentVideosStream>, Status> {
            let req = request.into_inner();
            let (tx, rx) = mpsc::channel(req.range as usize);
            for i in req.offset..(req.offset + req.range) {
                tx.send(Ok(RecentVideos{ video_id: format!("{i}"), video_thumbnail_id: "".to_owned(), upload_date: "".to_owned(), video_title: format!("title-{i}") })).await.unwrap();
            }
            Ok(Response::new(ReceiverStream::new(rx) as Self::GetRecentVideosStream))
        }

}
