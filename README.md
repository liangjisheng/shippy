# golang 微服务

[教程](https://segmentfault.com/u/wuyin/articles)

[github-zh](https://github.com/wuYin/shippy)

[github-en](https://github.com/EwanValentine/shippy)

本节先使用 go-micro 的 NATS 消息代理插件，使 user-service 在创建新用户时发布一个带有用户信息且 topic 为 “user.created” 的消息事件，订阅了此 topic 的 email-service 接收到消息后取出用户信息来发送邮件。之后使用 go-micro 自带的 pubsub 层代替了 NATS 充分发挥 protobuf 通信的优势