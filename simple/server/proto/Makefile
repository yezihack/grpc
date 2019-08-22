#!/usr/bin/env bash

# 生成golang代码
init:
	protoc --go_out=plugins=grpc:. *.proto
	sed -i "" "s/,omitempty//g" *.pb.go
