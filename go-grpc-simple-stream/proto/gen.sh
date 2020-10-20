#!/usr/bin/env bash
set -x

# 第一步： 安装 protoc 插件
# 打开下面URL， 跟据自己的系统选择对应的 protoc-3.x.x-linux|osx|win
# https://github.com/protocolbuffers/protobuf/releases
# 下载完后，加入到环境变量或path中, 保证全局可用。
# 验证: protoc --version

# 第二步：引用 proto, protoc-gen-go, grpc 共3个工具包
# 安装 golang 的proto工具包
# go get -u github.com/golang/protobuf/proto
# 安装 goalng 的proto编译支持
# go get -u github.com/golang/protobuf/protoc-gen-go
# 安装 GRPC 包
# go get -u google.golang.org/grpc

protoc --go_out=plugins=grpc:. *.proto
