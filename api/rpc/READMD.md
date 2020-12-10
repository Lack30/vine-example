# RPC API

## 用法

启动 vine api

```
micro api --handler=rpc
```

启动服务

```
go run rpc.go
```

## 请求服务

```
curl -H 'Content-Type: application/json' -d '{"name": "john"}' "http://localhost:8080/example/call"
```

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```