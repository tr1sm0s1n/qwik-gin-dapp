FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go build scripts/deploy.go
RUN go build -o server

FROM alpine

COPY --from=builder /app/deploy /app/deploy
COPY --from=builder /app/server /app/server
