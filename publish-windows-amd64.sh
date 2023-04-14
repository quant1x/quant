#!/bin/sh
# 获取当前路径, 用于返回
p0=`pwd`
# 获取脚本所在路径, 防止后续操作在非项目路径
p1=$(cd $(dirname $0);pwd)

version=$(git describe --tags `git rev-list --tags --max-count=1`)
version=${version:1}
echo "version: ${version}"

application=quant
# 安装路径
BIN=~/runtime/bin
EXT=

# windows amd64
GOOS=windows
GOARCH=amd64

echo "正在编译应用:${application} => $BIN/${application}..."
env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-X 'main.MinVersion=${version}'" -o $BIN/quant${EXT} github.com/quant1x/quant/strategy
echo "正在编译应用:${application} => $BIN/${application}...OK"
cd $p0