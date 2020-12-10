# Proxy API

## 用法

启动 vine api 作为代理

```
micro api --handler=http
```

启动 http 服务

```
go run proxy.go
```

## 请求服务

```
curl "http://localhost:8080/example/call?name=john"
```

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```
