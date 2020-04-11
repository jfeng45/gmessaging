##  一个通用的Go消息中间件接口

其他语言：

### **[English](README.md)**

这是通用的Go消息消息中间件接口。 目的是在应用程序中平滑地支持不同的消息中间件，并避免锁定到特定的消息中间件。 它没有包含消息传递服务的所有功能，而仅包括我需要的部分。 你可以轻松地对其进行扩展以添加新的功能。 当前，它含有[“ NATS”](https://docs.nats.io/developing-with-nats/developer)的实现。 您可以为其他消息传递服务（例如Kafka）添加实现。

有关如何在实际项目中使用它的示例，请查看[“ Order”](https://github.com/jfeng45/order)或[“ Payment”](https://github.com/jfeng45/payment)

### 下载程序

```
go get github.com/jfeng45/gmessaging
```

### 授权

[MIT](LICENSE.txt) 授权


