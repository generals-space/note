# go build出现out of memory错误

场景描述

在阅读源码kubelet的时候, 希望使用`go run`进行调试, 但是kubelet要求禁用swap分区, 使用`swapoff -a`禁用后再执行, 发现提示内存不足:

```
# go run kubelet.go --kubeconfig=/usr/local/kubernetes/kubelet/kubelet.conf --config=/usr/local/kubernetes/kubelet/config.yaml --network-plugin=cni --pod-infra-container-image=registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.1
# command-line-arguments
fatal error: runtime: out of memory
```

开发机是4G内存, 在kubernetes工程目录下执行make都没有问题, 而且用`free -m`监控内存占用, 还有1.7G左右的空闲.

这个问题没有来的及仔细研究, 网上的文章指写到问题出在`link`阶段. 为了跳过这个问题, 只能先把`swap`打开, 使用kube的make将kubelet单独编译成二进制文件, 但把`swap`关掉来执行.
