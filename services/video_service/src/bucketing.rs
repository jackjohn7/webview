use s3::request::ResponseData;

pub enum BucketError {
    BucketNotFound,
    FileUploadFailure,
    FileNotFound,
}

pub trait BucketStorageService {
    fn upload_file(&self, data: &[u8], bucket_name: &str) -> Result<(), BucketError>;
    fn delete_file(&self, file_name: String, bucket_name: &str) -> Result<(), BucketError>;
    fn get_file(&self, file_name: String, bucket_name: &str) -> Result<ResponseData, BucketError>;
}

