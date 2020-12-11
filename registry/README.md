# registry

registry etcd

## 用法

> etcd 和 grpc 版本冲突，需要修改 go.mod 
> go mod 中加入: `replace google.golang.org/grpc => google.golang.org/grpc v1.26.0`

## 启动服务

go run --registry=etcd --registry-address=192.168.2.130:12379

