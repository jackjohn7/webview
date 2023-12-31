FROM golang:1.21 AS builder

# PHASE 1: BUILD

# set working directory for the container

WORKDIR /build

# copy dep files into wd
COPY go.mod go.sum ./

# install deps
RUN go mod download

# Copy all .go files from the current directory and its subdirectories
COPY *.go **/*.go ./
#COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o /web_client main.go

# PHASE 2: DEPLOY
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /web_client /web_client

EXPOSE 5173

CMD ["./web_client"]
