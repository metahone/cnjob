[TOC]

# 课后练习 2.2

* 编写一个 HTTP 服务器，此练习为正式作业需要提交并批改
* 鼓励群里讨论，但不建议学习委员和课代表发满分答案，给大家留一点思考空间
* 大家视个人不同情况决定完成到哪个环节，但尽量把1都做完  
  1.接收客户端 request，并将 request 中带的 header 写入 response header  
  2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header  
  3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出  
  4.当访问 localhost/healthz 时，应返回200  

> 注意： 默认日志级别 warn, 如要看到 http 请求输出，需要设置日志级别 --log_level=info

# 课后练习 3.2

* 构建本地镜像
```bash
  make build
```

* 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
  * 请思考有哪些最佳实践可以引入到 Dockerfile 中来  
    最基本的 Dockerfile 最佳实践，比如：  
    1. 应用进程无状态，无共享和其他依赖, 做到 Dockerfile 单进程模型 
    2. 尽量使用尺寸小的基本镜像, 以免打出来得镜像尺寸过大  
    3. 镜像里最好要有一些基本工具, 例如 ping 等, 便于有时候我们需要进入容器里针对一些问题进行调试和排错, 比如使用 busybox 或根据需求选择其他  
    4. 可以将一些关键环境变量等信息打入镜像, 例如本例在 docker build 镜像之前, 通过 sed 将 ${GIT_TAG} 写入了镜像的环境变量 VERSION 中
    ```bash
    ...
    docker:
        sed "s/ENV VERSION=\"\"/ENV VERSION=${GIT_TAG}/" Dockerfile > Dockerfile.tmp
    ...
    ```
  * 有一些最佳实践不仅可以用于 Dockerfile, 例如本例在 Makefile build 编译时, 把一些通过 shell 获取的信息编译时写入了`GitCommit`, `GitTag` 和 `BuildDate` 程序变量中, 运行时 -V 参数可以获取这些信息作为版本信息输出. 上述这样做的目的是, 便于快速定位代码的具体提交和版本, 如果程序运行有问题, 用于代码迅速检出排查


* 将镜像推送至 docker 官方镜像仓库
```bash
  make docker_push
```

* 通过 docker 命令本地启动 httpserver
```bash
  docker run --rm --name httpserver -d -p 80:80 metazone/httpserver
```

* 通过 nsenter 进入容器查看 IP 配置
```bash
# PID=$(docker inspect --format "{{ .State.Pid }}" httpserver)
# nsenter -t $PID -n ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
28: eth0@if29: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```
