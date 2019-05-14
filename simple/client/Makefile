#!/usr/bin/env bash

all: dev run

fmt:
	gofmt -l -w ./app/

install: fmt clean

clean:
	rm -rf output/

dev: install
	go build -o sh/chat_client ./app

test: install
	go build  -o sh/chat_client ./app

run:dev
	sh/chat_client

