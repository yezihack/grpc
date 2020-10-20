#!/usr/bin/env bash

# -I 添加any proto插件

protoc -I=/usr/include/ -I=./ --go_out=plugins=grpc:. *.proto
