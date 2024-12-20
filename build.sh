#!/bin/bash

# 设置 Go 应用的入口文件（通常是包含 main 函数的 Go 文件）
MAIN_FILE=./cmd/main.go

# 设置输出可执行文件的名称和路径
OUTPUT_FILE=./gin-mysql

# 确保输出目录存在
mkdir -p $(dirname "$OUTPUT_FILE")

# 编译 Go 应用
go build -o "$OUTPUT_FILE" "$MAIN_FILE"

# 检查编译是否成功
if [ $? -ne 0 ]; then
  echo "Error: Go build failed."
  exit 1
fi

echo "Success: Go app built and output to $OUTPUT_FILE"