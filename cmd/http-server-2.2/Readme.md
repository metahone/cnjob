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