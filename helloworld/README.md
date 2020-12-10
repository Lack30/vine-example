# HelloWorld

## 启动服务端

```bash
go run server/server.go
```

## 启动服务端

```bash
go run client/client.go

# output: 
#   result: reply:"hello world"
```

## 启动 vine api

```bash
vineadm api --handler=rpc
```

```bash
curl -H 'Content-Type: application/json' -d '{"name": "vine1"}' "http://localhost:8080/api/v1/call"
# output: 
#  {"reply":"hello vine1"}
```