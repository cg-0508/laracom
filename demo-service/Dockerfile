FROM golang:1.13-alpine as builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn


# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add git

WORKDIR /app/laracom-demo-service

# 将当前项目所在目录代码拷贝到镜像中
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o laracom-demo-service

#第二阶段的镜像构建
FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add --no-cache bash supervisor ca-certificates
RUN mkdir /app
WORKDIR /app

# 从上一个阶段构建的 builder 容器中拉取
COPY --from=builder /app/laracom-demo-service/laracom-demo-service .

ADD supervisord.conf /etc/supervisord.conf
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]