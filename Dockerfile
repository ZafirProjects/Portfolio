# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o portfolio-web-server ./cmd

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/portfolio-web-server .
EXPOSE 3000
CMD ["./portfolio-web-server"]
