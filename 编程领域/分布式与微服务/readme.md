参考文章

1. [大家都在说的分布式系统到底是什么？](https://juejin.im/post/5af8ea34f265da0b9f40622a)
    - 与集中式系统对比, 得出分布式系统的优点
    - 分布式系统的特性: 分布性, 透明性, 同一性, 通信性
    - 常用的分布式方案: 应用和服务(go-kit), 静态资源(Nginx, CDN), 数据和存储(Postgres, Ceph), 计算(Hadoop)
2. [构建高性能微服务架构](https://www.infoq.cn/article/building-a-high-performance-micro-service-architecture)
    - 从实际场景分析 传统服务 -> 微服务 演化的必要性, 由此也体现了微服务架构的优势: 快速迭代, 方便测试与运维
    - 提到Kube身为开源系统, 存在的一些不足.
3. [踩坑实践：如何消除微服务架构中的系统耦合？](https://mp.weixin.qq.com/s/-yvNs7Az_bDLvyJlCiSc6g)
4. [微服务架构下的软件测试实践](https://blog.csdn.net/weixin_41978708/article/details/80025231)

## 分布式与集群的关系

- 分布式(distributed): 在多台不同的服务器中部署**不同的服务模块**, 通过远程调用协同工作, 对外提供服务. 
- 集群(cluster): 在多台不同的服务器中部署**相同的应用或服务模块**, 构成一个集群, 通过负载均衡设备对外提供服务. 

## 分布式与微服务的关系

可以说**微服务是分布式服务的一种具体实现**, 但并不是唯一实现...?
