## grpc 实现服务注册， web 实现服务发现

### Grpc-server grpc 服务端
1. 第一步 创建 proto 文件
2. 第二步 使用 go-micro 注册服务
3. 第三步 实现 proto 逻辑

### Web-client 客户端
1. 第一步 使用 go-micro 实现服务发现
2. 第二步 调用 grpc 服务端提供的方法
3. 第三步 对外提供 web 服务