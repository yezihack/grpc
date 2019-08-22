#!/usr/bin/env bash

# 定义命名空间
namespace='app.grpc.'

#自动向proto文件添加命名空间

sed -i '' 's/package /package '"$namespace"'/g' proto_files/*.proto

protoc --php_out=output --grpc_out=output --plugin=protoc-gen-grpc=plugin/grpc_php_plugin proto_files/*.proto

sed -i '' 's/app.grpc.//g' output/*/*/*/*Client.php

sed -i '' 's/'"$namespace"'//g' proto_files/*.proto