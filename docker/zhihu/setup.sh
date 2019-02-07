#!/bin/bash
cd /go/src/github.com/novioleo/zhihu
echo 'Installing'
go build main.go
echo 'Run zhihu'
zhihu
