FROM golang:alpine AS development
MAINTAINER leyius "leyius@163.com"

# ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/go-gin
COPY . .
RUN go build -o app

FROM alpine:latest AS production
WORKDIR /root/
COPY --from=development /go/src/go-gin/app .

ENTRYPOINT ["./app"]