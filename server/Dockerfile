FROM golang:1.24-alpine AS builder

ENV GO_BUILDFLAGS="-trimpath -ldflags=-s -ldflags=-w"
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go build $GO_BUILDFLAGS scripts/deploy.go && \
    go build $GO_BUILDFLAGS -o server

FROM alpine

COPY --from=builder /app/deploy /app/deploy
COPY --from=builder /app/server /app/server
