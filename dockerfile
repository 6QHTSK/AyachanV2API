FROM golang:1.21 AS builder

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"

# 设置工作目录为/app
WORKDIR /app

# 将Go项目的代码复制到容器中
COPY . .

# 下载依赖并构建项目
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --md .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# 设置容器启动时运行的命令
FROM alpine:latest

MAINTAINER 6QHTSK <psk2019@qq.com>

ENV USE_ENV true
ENV RUN_ADDR 0.0.0.0:9000
ENV BD_API https://proxy.ayachan.fun
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /app

# 从编译阶段复制可执行文件到Alpine镜像中
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

EXPOSE 9000
CMD ["./main"]
