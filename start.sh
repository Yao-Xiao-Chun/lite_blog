#!/bin/bash
#开始构建程序
echo "build:tidy"
#设置国内源
export GOPROXY=https://proxy.golang.com.cn,direct
#同步包
/usr/local/go/bin/ go mod tidy
pwd
echo "start build"
/usr/local/go/bin go build -o lite_blog ./main.go
BUILD_ID=DONTKILLME
echo "build:success"
#检测服务是否存活
if pgrep lite_blog > /dev/null
then
  killall lite_blog # kill lite_blog
fi
#启动服务
chmod +x ./lite_blog
nohup ./lite_blog >> lite_blog.log 2>&1 &
echo "run lite_blog success"
