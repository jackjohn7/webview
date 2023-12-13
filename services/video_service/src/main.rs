mod server;
mod video_processing;

use server::MyVideoProcessor;
use video_processing::video_processing_service_server::VideoProcessingServiceServer;
use tonic::transport::Server;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:3001".parse().unwrap();
    let processor_service = MyVideoProcessor::default();

    Server::builder()
        .add_service(VideoProcessingServiceServer::new(processor_service))
        .serve(addr)
        .await?;

    Ok(())
}
