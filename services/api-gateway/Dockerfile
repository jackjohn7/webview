FROM golang:1.21 AS builder

# PHASE 1: BUILD

# set working directory for the container

WORKDIR /build

# copy service files into wd

COPY go.mod go.sum ./

RUN go mod download


# copy the entirety of the project into the working directory (technically not ideal)
COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o /api-gateway main.go

# PHASE 2: DEPLOY

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /api-gateway /api-gateway

EXPOSE 3000

CMD ["./api-gateway"]
