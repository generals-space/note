# RPC框架选型

参考文章

1. [gRPC基于Golang和Java的简单实现]

1. [Dubbo 实践，演进及未来规划](https://my.oschina.net/u/3959468/blog/3002301)
    - dubbo框架的组件元素和工作流程(注册与订阅).
    - 以实际案例解释dubbo的不足, 解决方案和演进过程
    - 应用级注册与服务级注册的概念
    - 各服务中间件的对比和各自的优缺点, 一些概念上的实现(CAP理论, 去中心化, 推送机制, 和存储容量等)
    - 集中配置中心概念的必要性和部署方案

dubbo本质上是RPC框架, 但ta默认集成了zk和redis, 作为服务发现的中间件, 因此也实现了服务发现的功能. 不像gRPC, 只有RPC的功能, 要附加服务发现的功能, 只能部署etcd/consul服务, 并额外使用客户端SDK, 手动完成注册, 发现和健康检查的操作.
