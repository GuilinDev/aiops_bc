# Build Stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY main/ .
# 使用静态链接构建
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server tcp_server.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client tcp_client.go
RUN ls -l server client  # 检查构建的文件

# Final Stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/client .
RUN ls -l /app/server /app/client
RUN chmod +x /app/server /app/client
RUN apk add --no-cache file && file /app/server
EXPOSE 3333
CMD ["/app/server"]
