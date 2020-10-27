## go-micro web 服务注册与服务发现 实现代码

### 服务端代码
> 代码在 server 文件夹

服务端实现将 web 服务注册到服务里，使用服务发现去调用它们。
服务端启动三个服务(--server_address是 go-micro 自带的参数)
```shell script
go run . --server_address :8000
go run . --server_address :8001
go run . --server_address :8002
```

### 客户端
> 代码在 client 文件夹里

通过服务名获取服务注册的IP:PORT, 调用服务端