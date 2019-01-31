# Kubernetes-PersistentVolume持久卷

参考文章

1. [Kubernetes对象之PersistentVolume，StorageClass和PersistentVolumeClaim](https://www.jianshu.com/p/99e610067bc8)

    - 对pv和pvc的概念解释的比较浅显易懂

2. [Configure a Pod to Use a PersistentVolume for Storage](https://kubernetes.io/docs/tasks/configure-pod-container/configure-persistent-volume-storage/)

    - kubernetes官方文档
    
参考文章2是用minikube创建pv和pvc的使用示例, 并且指定要node节点只有一个. 我在实验的时候用的是阿里云的kubernetes集群, 2个node节点, 结果发现最终pod运行的时候只有一个node上有`hostPath`指定的路径. 如果这个路径不存在, kubernetes会自动创建.

创建pv资源

```yml
## postgres pv
kind: PersistentVolume
apiVersion: v1
metadata:
  name: demo-volume
  labels:
    type: demo
spec:
  ## storageClassName用于绑定pv和pvc, 可以自定义.
  storageClassName: demo-pv
  capacity:
    storage: 1Gi
  accessModes:
    - ReadOnlyMany
  hostPath:
    path: "/mnt/demo"
---

## postgresql pvc
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: demo-claim
spec:
  storageClassName: demo-pv
  accessModes:
    - ReadOnlyMany
  resources:
    requests:
      storage: 1Gi
```

使用资源

```yml
volumes:
  - name: demo-vol
    persistentVolumeClaim:
      claimName: demo-claim
containers:
- name: demo
  image: nginx
  volumeMounts:
  - name: demo-vol
    mountPath: "/usr/share/nginx/html"
```