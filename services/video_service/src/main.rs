mod server;
mod video_processing;
mod video_streaming;
mod bucketing;
mod utils;

use std::sync::Arc;

use server::{MyVideoProcessor, MyVideoCDN};
use video_processing::video_processing_service_server::VideoProcessingServiceServer;
use video_streaming::video_streaming_service_server::VideoStreamingServiceServer;
use tonic::transport::Server;
use awscreds::Credentials;
use awsregion::Region;
use s3::{Bucket, BucketConfiguration};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {

    // get bucket storage connection
    let bucket_name = "video";
    let region = Region::Custom {
        region: "us-west-3".to_owned(),
        endpoint: "http://localhost:9000".to_owned(),
    };
    let credentials = Credentials::new(Some("ROOTNAME"), Some("CHANGEME123"), None, None, None)?;
    let mut bucket =
        Bucket::new(bucket_name, region.clone(), credentials.clone())?.with_path_style();

    // if video bucket doesn't exist, create it
    if !bucket.exists().await? {
        bucket = Bucket::create_with_path_style(
            bucket_name,
            region,
            credentials,
            BucketConfiguration::default(),
        )
        .await?
        .bucket;
    }


    println!("getting here");
    let addr = "[::1]:3001".parse().unwrap();
    let processor_service = MyVideoProcessor::new(Arc::new(bucket.clone()));
    let cdn_service = MyVideoCDN::new(Arc::new(bucket));

    Server::builder()
        .add_service(VideoProcessingServiceServer::new(processor_service))
        .add_service(VideoStreamingServiceServer::new(cdn_service))
        .serve(addr)
        .await?;

    Ok(())
}
