# Meta API

## 使用

启动 vine api

```
vineadm api
```

启动服务

```
go run meta.go
```

## 请求服务

```
curl -H 'Content-Type: application/json' -d '{"name": "john"}' "http://localhost:8080/example"
```


```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/foo/bar
```
