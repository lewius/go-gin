# go-gin
golang web service by gin


### 项目依赖管理
1. go mod init
2. go.mod文件配置依赖
3. go mod tidy

> 使用代理解决依赖下载失败问题：$env:GOPROXY="https://goproxy.io"

go get -v -t -d ./...

go build .

### go函数
1. 不支持重载、默认参数
2. 支持多返回值
3. 函数名前的参数为“方法接收者”


短变量声明 :=