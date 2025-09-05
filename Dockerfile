# 多阶段构建Docker镜像
FROM golang:1.18-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的工具
RUN apk add --no-cache git ca-certificates tzdata

# 复制go模块文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o main ./examples/basic

# 最终运行镜像
FROM scratch

# 从builder镜像复制必要文件
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/main /app/main

# 设置环境变量
ENV TZ=Asia/Shanghai

# 暴露端口（如果有web服务的话）
# EXPOSE 8080

# 设置用户
USER 65534:65534

# 运行应用
ENTRYPOINT ["/app/main"]