# 构建阶段1: 以golang为基础镜像
FROM golang AS build
# 设置变量location
ENV location /go/src/github.com/code-newbee/geeker
WORKDIR ${location}/server
ADD ./ ${location}/server

# RUN 表示执行后面的命令
# 1. RUN <命令>
# 2. RUN ["可执行文件", "参数1", "参数2", ...]
RUN go get -d ./...
RUN go install ./...
RUN CGO_ENABLED=0 go build -o /bin/geeker

# 构建阶段2:
# scratch 为docker官方的保留关键字, 表示一个空镜像
FROM scratch
# 复制阶段1中编译的二进制文件到scratch
COPY --from=build /bin/geeker /bin/geeker
# 启动入口, docker run 的参数会被当作参数传递给该文件
ENTRYPOINT ["/bin/geeker"]

# 声明端口, EXPOSE 仅是展示作用
EXPOSE 50001