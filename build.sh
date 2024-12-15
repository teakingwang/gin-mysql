#!/bin/bash

# 设置Go环境变量（可选，如果已经在系统级别设置则不需要）
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# 定义项目目录和输出目录
PROJECT_DIR=$(pwd)
OUTPUT_DIR=$PROJECT_DIR/output
BINARY_NAME=my_go_app

# 创建输出目录（如果不存在）
mkdir -p $OUTPUT_DIR

# 编译Go项目
echo "Compiling Go project..."
go build -o $OUTPUT_DIR/$BINARY_NAME $PROJECT_DIR/main.go

# 检查编译是否成功
if [ $? -ne 0 ]; then
  echo "Compilation failed!"
  exit 1
fi

# （可选）打包编译后的二进制文件和资源
echo "Packaging compiled binary..."
TAR_NAME=$BINARY_NAME-$(date +%Y%m%d%H%M%S).tar.gz
tar -czf $OUTPUT_DIR/$TAR_NAME -C $OUTPUT_DIR $BINARY_NAME

# 输出打包结果
echo "Package created: $OUTPUT_DIR/$TAR_NAME"

# （可选）清理编译后的二进制文件（如果需要保留则删除此行）
# rm $OUTPUT_DIR/$BINARY_NAME