# go-gin
golang web service by gin


### 项目依赖管理
1. go mod init
2. go.mod文件配置依赖
3. go mod tidy

> 使用代理解决依赖下载失败问题：$env:GOPROXY=https://goproxy.io


go get -v -t -d ./...

go build .