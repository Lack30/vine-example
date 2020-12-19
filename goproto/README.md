# goproto

goproto-gen 工具使用实例

## 安装 

```shell
go get github.com/gogo/protobuf
go get github.com/lack-io/vineadm/cmd/protoc-gen-gogofaster
go get github.com/lack-io/goproto/cmd/protoc-gen-gogofaster
```

## 生成 *.proto
```shell
goproto-gen --packages "github.com/lack-io/vine-example/goproto/api"
```