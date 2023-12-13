mod server;
mod video_processing;
mod video_streaming;

use server::{MyVideoProcessor, MyVideoCDN};
use video_processing::video_processing_service_server::VideoProcessingServiceServer;
use video_streaming::video_streaming_service_server::VideoStreamingServiceServer;
use tonic::transport::Server;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:3001".parse().unwrap();
    let processor_service = MyVideoProcessor::default();
    let cdn_service = MyVideoCDN::default();

    Server::builder()
        .add_service(VideoProcessingServiceServer::new(processor_service))
        .add_service(VideoStreamingServiceServer::new(cdn_service))
        .serve(addr)
        .await?;

    Ok(())
}
