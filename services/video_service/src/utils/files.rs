use std::path::PathBuf;
use std::process::Command;

pub struct VideoChunk {
    idx: u64,
    path: PathBuf
}

pub fn compress_video<'a>(id: &String, og_path: &PathBuf, output_path: PathBuf) -> Result<PathBuf, &'a str> {
    let mut command = Command::new("ffmpeg");
    let file_name = format!("{}_compressed.mp4", id);
    let output_file_path = output_path.join(file_name);
    command.args([
        "-i",
        og_path.to_str().unwrap(),
        "-vcodec",
        "h264",
        "-acodec",
        "mp2",
        output_file_path.to_str().unwrap()
    ]);

    let output = command
        .status()
        .expect("Failed to compress video");

    if !output.success() {
        Err("Failed to compress video")
    } else {
        Ok(output_file_path)
    }
}

pub fn get_video_duration(path: &PathBuf) -> u64 {
    let output = Command::new("ffprobe")
        .arg("-v")
        .arg("error")
        .arg("-show_entries")
        .arg("format=duration")
        .arg("-of")
        .arg("default=noprint_wrappers=1:nokey=1")
        .arg(path)
        .output()
        .expect("failed probe video");
    let output_str = String::from_utf8_lossy(&output.stdout);
    println!("output_str: {}", output_str);
    let duration = output_str.trim().parse::<f64>().unwrap();
    duration as u64
}

fn get_start_end_times(chunk_idx: u64, num_chunks: u64, duration: u64, chunk_size: u64) -> (u64, u64) {
    let start = chunk_idx * chunk_size;
    let end = match num_chunks.cmp(&(chunk_idx + 1)) {
        std::cmp::Ordering::Equal => match duration % chunk_size {
            0 => start + chunk_size,
            extra => start + extra
        }
        _ => start + chunk_size
    };

    (start, end)
}

fn get_num_chunks(duration: u64, chunk_size: u64) -> u64 {
    match duration % chunk_size {
        0 => duration / chunk_size,
        _ => (duration / chunk_size) + 1
    }
}

fn to_timestamp(seconds: u64) -> String {
    let secs = seconds % 60;
    let minutes = (seconds / 60) % 60;
    let hours = ((seconds / 60) / 60) % 24;
    format!("{:02}:{:02}:{:02}", hours, minutes, secs)
}

pub fn chunk_video_data<'a>(id: &String, video_path: &PathBuf, output_path: PathBuf, duration: u64, chunk_size: u64) -> Result<Vec<VideoChunk>, &'a str> {
    let num_chunks = get_num_chunks(duration, chunk_size);
    let mut chunks = Vec::new();

    for idx in 0..num_chunks {
        let (start, end) = get_start_end_times(idx, num_chunks, duration, chunk_size);
        let mut command = Command::new("ffmpeg");
        let file_name = format!("{}_chunk_{}.mp4", id, idx);
        let output_file_path = output_path.join(file_name);
        command.args([
            "-f",
            "mp4",
            "-ss",
            &to_timestamp(start),
            "-to",
            &to_timestamp(end),
            "-i",
            video_path.to_str().unwrap(),
            "-c",
            "copy",
            "-avoid_negative_ts",
            "make_zero",
            output_file_path.to_str().unwrap(),
        ]);
        let output = command
            .status()
            .expect("Failed to chunk video");
        if !output.success() {
            return Err("Failed to chunk video");
        }

        chunks.push(VideoChunk{ idx, path: output_file_path });
    }

    Ok(chunks)
}

#[cfg(test)]
mod tests{
    use super::{to_timestamp, get_num_chunks, get_start_end_times};

    #[test]
    fn test_to_timestamp() {
        assert_eq!(to_timestamp(0),            "00:00:00");
        assert_eq!(to_timestamp(1),            "00:00:01");
        assert_eq!(to_timestamp(30),           "00:00:30");
        assert_eq!(to_timestamp(59),           "00:00:59");
        assert_eq!(to_timestamp(60),           "00:01:00");
        assert_eq!(to_timestamp(61),           "00:01:01");
        assert_eq!(to_timestamp(60 * 59),      "00:59:00");
        assert_eq!(to_timestamp(60 * 59 + 13), "00:59:13");
        assert_eq!(to_timestamp(60 * 60),      "01:00:00");
        assert_eq!(to_timestamp(60 * 60 + 1),  "01:00:01");
    }
    #[test]
    fn test_get_num_chunks() {
        assert_eq!(get_num_chunks(234, 5), 47);
        assert_eq!(get_num_chunks(233, 5), 47);
        assert_eq!(get_num_chunks(226, 5), 46);
        assert_eq!(get_num_chunks(225, 5), 45);
        assert_eq!(get_num_chunks(224, 5), 45);
    }

    #[test]
    fn test_get_start_end_times() {
        let duration = 23;
        let chunk_size = 5;
        let num_chunks = get_num_chunks(duration, chunk_size);
        let (start, end) = get_start_end_times(0, num_chunks, duration, chunk_size);
        assert_eq!(start, 0);
        assert_eq!(end, 5);
        let (start, end) = get_start_end_times(1, num_chunks, duration, chunk_size);
        assert_eq!(start, 5);
        assert_eq!(end, 10);
        let (start, end) = get_start_end_times(2, num_chunks, duration, chunk_size);
        assert_eq!(start, 10);
        assert_eq!(end, 15);
        let (start, end) = get_start_end_times(3, num_chunks, duration, chunk_size);
        assert_eq!(start, 15);
        assert_eq!(end, 20);
        let (start, end) = get_start_end_times(4, num_chunks, duration, chunk_size);
        assert_eq!(start, 20);
        assert_eq!(end, 23);
    }
}
