FROM golang:alpine

ADD ./go-gin.exe
ENTRYPOINT ["go-gin.exe"]