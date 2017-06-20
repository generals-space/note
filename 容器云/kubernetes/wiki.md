
参考文章

1. [kubelet组件操作API - kubernetes 简介： kubelet 和 pod](http://cizixs.com/2016/10/25/kubernetes-intro-kubelet?utm_source=tuicool&utm_medium=referral)

2. [ovs+kube配置1 - CentOS 7实战Kubernetes部署](http://www.infoq.com/cn/articles/centos7-practical-kubernetes-deployment)

3. [ovs+kube配置1 - kube相关服务systemd脚本1](https://github.com/yangzhares/GetStartingKubernetes)

3. [ovs+kube配置2 - 基于OVS的Kubernetes多节点环境搭建](http://bingotree.cn/?p=828)

4. [kube相关服务systemd脚本2 - 轻松了解Kubernetes认证功能](http://www.tuicool.com/articles/byUnQn7)

5. [kubernetes 集群的安装部署](http://www.cnblogs.com/galengao/p/5780938.html)

6. [kubernetes-create-pods](http://www.liuhaihua.cn/archives/416728.html)

从kubelet直接获取pid/container信息: `curl http://172.32.100.70:10255/pods`

kubelet健康检查: `curl http://127.0.0.1:10248/healthz`

------

从apiserver获取节点信息

```
[root@localhost ~]# kubectl -s http://172.32.100.90:8080 get nodes
NAME            STATUS    AGE       VERSION
172.32.100.70   Ready     4m        v1.8.0-alpha.0.367+0613ae5077b280
172.32.100.80   Ready     4m        v1.8.0-alpha.0.367+0613ae5077b280
```

ovs+docker主机间容器通信必要步骤

1. ip_forward=1 每次重启network服务都会重置sysctl配置, 最好写在文件中.

2. route路由, 写配置文件

3. kbr0网卡配置(不可加HWADDR参数)

4. type=gre/vxlan应该都没关系

5. 关闭防火墙/selinux

6. NetworkManager关闭

ovs网络划分规则, 单个主机拥有其独立子网, 还是可以多主机同属同一子网, 容器ip在什么范围内保持唯一?

20170607
------

flannel, Open vSwitch的存在目的是实现容器之间的跨主机通信, 否则容器只能与其宿主机沟通. 跨主机的容器需要路由与ip支持.

Etcd来存储每台机器的上子网地址

FAQ

1. kubernetes的release只有8M, 好像没法用???