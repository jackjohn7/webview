FROM rust:1.72.0 as BUILDER

WORKDIR /usr/src/video_service

COPY . .

RUN cargo install --path .

FROM debian:buster-slim

RUN apt update

# Error is occurring. Cannot find protoc
#  Err msg suggests setting "PROTOC" environment variable
# Also will need to figure out how to replicate the protobufs for the build in Docker

RUN apt install -y ffmpeg protobuf-compiler

RUN apt-get -y extra-runtime-dependencies & rm -rf /var/lib/apt/lists/*

COPY --from=BUILDER /usr/local/cargo/bin/video_service /usr/local/bin/video_service

EXPOSE 3001

CMD ["video_service"]
