FROM golang:alpine
MAINTAINER leyius "leyius@163.com"

#设置工作目录
# WORKDIR /root

ADD ./go-gin app
ENTRYPOINT ["./app"]