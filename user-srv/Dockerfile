# 第一阶段需要go程序的编译和构建环境
FROM golang:1.26 AS builder

# 将当前目录下的所有文件都copy到容器中的/src目录下。相当于src替换user-src，/src是项目根目录
COPY . /src     
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

# 第二阶段只需要运行环境
FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

WORKDIR /app

# 从第一阶段的目录中，将构建好的二进制文件copy到/app目录下
COPY --from=builder /src/bin/user-srv ./user-srv

# 将运行需要的配置文件，从builder中copy过来
COPY --from=builder /src/configs ./configs


EXPOSE 8000
EXPOSE 9000

CMD ["./user-srv", "-conf", "./configs"]
