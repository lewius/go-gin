FROM golang:alpine

ADD ./go-gin.exe app.exe
ENTRYPOINT ["app.exe"]