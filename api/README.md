# API

vine api 使用实例

## vine api 简介

vine api 作为 *vine* 框架的网关组件，提供 HTTP 服务和基于服务发现组件的动态路由

vine api 默认命名空间为 go.vine.api. 可通过 `--namespace=` 参数修改。

vine api 支持不同种类的请求处理器，默认为 rpc。可通过 `--handler=` 参数修改

## 相关

- api - http 处理
- proxy - 反向代理
- rpc - rpc 转 http
- meta - 自定义配置

## 请求映射

### API/RPC 映射

vine api RPC 请求映射表

Path    |    Service    |    Method
----	|	----	|	----
/foo/bar    |    go.vine.api.foo    |    Foo.Bar
/foo/bar/baz    |    go.vine.api.foo    |    Bar.Baz
/foo/bar/baz/cat    |    go.vine.api.foo.bar    |    Baz.Cat

api 请求映射表

Path    |    Service    |    Method
----	|	----	|	----
/foo/bar    |    go.vine.api.foo    |    Foo.Bar
/v1/foo/bar    |    go.vine.api.v1.foo    |    Foo.Bar
/v1/foo/bar/baz    |    go.vine.api.v1.foo    |    Bar.Baz
/v2/foo/bar    |    go.vine.api.v2.foo    |    Foo.Bar
/v2/foo/bar/baz    |    go.vine.api.v2.foo    |    Bar.Baz

### proxy 映射

proxy 映射需要指定为 http 请求处理 `--handler=http`

Path    |    Service    |    Service Path
---	|	---	|	---
/greeter    |    go.vine.api.greeter    |    /greeter
/greeter/:name    |    go.vine.api.greeter    |    /greeter/:name
