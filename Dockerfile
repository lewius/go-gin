FROM golang:alpine
MAINTAINER leyius "leyius@163.com"

# ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/leyius/go-gin
COPY . $GOPATH/src/github.com/leyius/go-gin
RUN go build .

ENTRYPOINT ["./go-gin"]