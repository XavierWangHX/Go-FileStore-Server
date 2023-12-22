# Golang Based Distributed-FileServer

基于Golang实现的一种分布式云存储服务

## 启动主入口/Rabbitmq生产者(将文件异步转移到阿里云OSS)
go run main.go

## 启动Rabbitmq消费者(将文件异步转移到阿里云OSS)
go run service/transfer/main.go

## Pre-request packages 
关于需要手动安装的库

如下：
```shell
go get github.com/garyburd/redigo/redis
go get github.com/go-sql-driver/mysql
go get github.com/garyburd/redigo/redis
go get github.com/json-iterator/go
go get github.com/aliyun/aliyun-oss-go-sdk/oss
go get github.com/streadway/amqp