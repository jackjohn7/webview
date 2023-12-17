use std::fs::{self, File};
use std::io::{Read, Write};
use std::path::{Path, PathBuf};
use std::sync::Arc;
use futures::stream::StreamExt;

use tokio::sync::mpsc;
use tonic::{Request, Response, Status, Streaming};
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

use crate::bucketing::BucketStorageService;
use s3::Bucket;

pub struct MyVideoProcessor {
    storage_bucket: Arc<Bucket>
}

impl MyVideoProcessor {
    pub fn new(bucket: Arc<Bucket>) -> Self {
        Self{ storage_bucket: bucket }
    }
}

#[tonic::async_trait]
impl VideoProcessingService for MyVideoProcessor {
    async fn process_new_video(&self, req: Request<Streaming<ProcessVideoRequest>>)
        -> Result<Response<ProcessedVideoData>, Status> {
        let mut stream = req.into_inner();
        let mut video_id: String = String::new();
        let mut video_data = Vec::new();
        while let Some(chunk_result) = stream.next().await {
            match chunk_result {
                Ok(chunk) => {
                    let bytes = chunk.data;
                    video_data.extend_from_slice(&bytes);
                    video_id = chunk.video_id;
                },
                Err(err) => {
                    return Err(Status::aborted(format!("Error receiving video chunk: {}", err)));
                }
            }
        }
        // upload video to MinIO
        println!("num bytes: {}", video_data.len());

        // create temporary directory
        let storage_dir = format!("./temporary_storage/{}", video_id);
        fs::create_dir_all(&storage_dir).map_err(|e| {
            Status::internal(format!("Failed to create storage directory: {}", e))
        })?;
        // Create a unique filename for the stored file (e.g., using a timestamp)
        let file_name = format!("stored_file_{}.mp4", video_id);

        // Combine the storage directory and filename to get the full path
        let file_path = PathBuf::from(storage_dir).as_path().join(&file_name);

        // Open the file for writing
        let mut file = File::create(&file_path).map_err(|e| {
            Status::internal(format!("Failed to create file: {}", e))
        })?;

        // Write the bytes to the file
        file.write_all(&video_data)?;

        println!("User uploaded video with id: {}", video_id);
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

pub struct MyVideoCDN {
    storage_bucket: Arc<Bucket>
}

impl MyVideoCDN {
    pub fn new(bucket: Arc<Bucket>) -> Self {
        Self{ storage_bucket: bucket }
    }
}

#[tonic::async_trait]
impl VideoStreamingService for MyVideoCDN {

    type GetRecentVideosStream = ReceiverStream<Result<RecentVideos, Status>>;

    async fn stream_video_chunk(&self, request: Request<VideoChunkRequest>) -> Result<Response<VideoChunk>, Status> {
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
