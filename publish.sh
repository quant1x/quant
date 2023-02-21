#!/bin/sh
# 获取当前路径, 用于返回
p0=`pwd`
# 获取脚本所在路径, 防止后续操作在非项目路径
p1=$(cd $(dirname $0);pwd)

# windows amd64
env GOOS=windows GOARCH=amd64 go build -o bin/strategy-win-amd64.exe github.com/quant1x/quant/strategy
# darwin amd64
env GOOS=darwin GOARCH=amd64 go build -o bin/strategy-mac-amd64 github.com/quant1x/quant/strategy
cd $p0