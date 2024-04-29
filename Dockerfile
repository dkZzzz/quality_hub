FROM golang:1.21 AS builder
COPY . .
WORKDIR .
EXPOSE 8080
RUN go mod tidy
RUN go run cmd/api/main.go
