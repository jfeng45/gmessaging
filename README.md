# A Generic Go Messaging Interface

Other language: 
### **[中文](README.zh.md)**
 
This is a generic Go messaging interface. The purpose is to support different messaging libraries and avoid to lock-in into specific messaging implementation. It doesn't necessarily include all functionality for messaging service, but only the ones I need. You can easily expand it to add new ones. Currently, it has implementation for ["NATS"](https://docs.nats.io/developing-with-nats/developer). You can add implementation for other Messaging service like Kafka.

For examples on how to use it in real project, please take a look at ["Order"](https://github.com/jfeng45/order) or ["Payment"](https://github.com/jfeng45/payment)

### Download Code

```
go get github.com/jfeng45/gmessaging
```

### License

[MIT](LICENSE.txt) License


