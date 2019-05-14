#!/usr/bin/env bash

all: dev run

fmt:
	gofmt -l -w ./app/

install: fmt clean

clean:
	rm -rf sh/

dev: install
	go build -o sh/chat_service ./app

test: install
	go build  -o sh/chat_service ./app

run:dev
	sh/chat_service

