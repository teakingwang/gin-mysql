# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.18-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将项目的 Go 源码复制到 Docker 容器中
COPY . /app

# 构建你的 Go 应用
# 假设 build.sh 脚本负责编译并生成一个名为 gin-mysql 的可执行文件
RUN chmod +x /app/build.sh
RUN /app/build.sh

# 使用一个更小的镜像来运行你的应用
FROM alpine:latest

# 从构建阶段复制可执行文件到运行阶段
# 假设 build.sh 生成的可执行文件在当前目录下（即 /workspace/gin-mysql）
COPY --from=builder /app/gin-mysql .

# 暴露应用运行的端口（根据你的应用实际使用的端口进行修改）
EXPOSE 8080

# 设置容器启动时执行的命令
CMD ["./gin-mysql"]