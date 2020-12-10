# noproto

在没有 protobuf 的情况下使用 vine 服务

vine 支持 json 格式解析请求数据，客户端发送 `Content-Type=application/json` 类型数据时，服务端使用 json 解析并返回

## 内容

- main.go - vine 服务端
- client - 发送 json 数据的客户端

## 启动

启动服务端

```shell
go run main.go
```

启动客户端

```shell
go run client/main.go
```