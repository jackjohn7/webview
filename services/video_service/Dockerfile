FROM rust:bookworm as BUILDER

# Still will need to figure out how to replicate the protobufs for the build in Docker
RUN apt update && apt upgrade -y

RUN apt install -y protobuf-compiler

RUN apt-get -y extra-runtime-dependencies & rm -rf /var/lib/apt/lists/*
WORKDIR /usr/src/video_service

COPY . .

RUN cargo install --path .

FROM debian:bookworm-slim

RUN apt update && apt upgrade -y

# Error is occurring. Cannot find protoc
#  Err msg suggests setting "PROTOC" environment variable
ENV PATH "$PATH:/usr/bin/protoc"
ENV PROTOC=/usr/bin/protoc
# Also will need to figure out how to replicate the protobufs for the build in Docker

RUN apt install -y ffmpeg

RUN apt-get -y extra-runtime-dependencies & rm -rf /var/lib/apt/lists/*

RUN apt install -y openssl

COPY --from=BUILDER /usr/local/cargo/bin/video_service /usr/local/bin/video_service

EXPOSE 3001

CMD ["video_service"]
