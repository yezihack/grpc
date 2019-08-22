#!/usr/bin/env bash

fmt:
	gofmt -l -w ./

build:fmt clean
	go build -o output/go_service ./

clean:
	rm -rf output/

run:build
	output/go_service
