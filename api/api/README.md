# API

## 使用

启动 vine api

```
vineadm api --handler=api
```

启动服务

```
go run api.go
```


## 请求服务

```
curl "http://localhost:8080/example/call?name=john"
```

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```

## 设置 Namespace

```
vineadm api --handler=api --namespace=com.foobar.api
```

or
```
VINE_API_NAMESPACE=com.foobar.api vineadm api --handler=api
```

设置服务的 namespace

```
service := vine.NewService(
        vine.Name("com.foobar.api.example"),
)
```  