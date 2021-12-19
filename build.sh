#!/bin/bash
# MAC 打包linux可执行文件
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pindex