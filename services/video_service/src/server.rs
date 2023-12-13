use tonic::{Request, Response, Status};
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
        Ok(Response::new(HealthCheckResponse { msg: "pongCUH".to_owned() }))
    }
}
