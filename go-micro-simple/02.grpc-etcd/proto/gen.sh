#!/usr/bin/env bash

# 第一步： 安装 protoc 插件
# 打开下面URL， 跟据自己的系统选择对应的 protoc-3.x.x-linux|osx|win
# https://github.com/protocolbuffers/protobuf/releases
# 下载完后，加入到环境变量或path中, 保证全局可用。
# 验证: protoc --version

# 安装插件
# install protoc-gen-go
# go get github.com/golang/protobuf/{proto,protoc-gen-go}
# install protoc-gen-micro
# go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master

# protoc 生成 go 代码，加载 go-micro 插件
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. *.proto
