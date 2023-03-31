# EZ

简单开发框架

一个开箱即用的框架，终极缝合怪

## 原因

市面上的各个框架虽然功能很多，但是并没有一款能做到开箱即用，业务零侵入的，需要自己增加数据初始化，kafka初始化，grpc client初始化，业务代码掺杂了太多无关代码。所以这里再次缝合一下，支持开箱即用

## 集成

### 数据库

+ sql
+ sqlx
+ gorm
+ go-redis
+ kafka

### framework

+ kratos
+ kite(todo)
+ grpc
+ gin

## 服务类型

### HTTP

基于`gin`框架进行HTTP服务监听

### 默认使用了这2个中间件

```sh
router.Use(gin.Recovery())
router.Use(trace.InjectHeadersIntoMetadata())
```

### 默认使用了如下配置

```sh
gin.SetMode(gin.ReleaseMode)
gin.DisableConsoleColor()
```

#### JSON替换

`gin`内部支持json库的替换，用于提高性能，支持`jsoniter`和`go-json`, 构建时传入对应的tag就能使用对应的json库

```sh
go build -tags=jsoniter .

go build -tags=go_json .
```
