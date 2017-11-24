# docker容器与宿主机之间相互拷贝数据

[参考文章](http://stackoverflow.com/questions/22907231/copying-files-from-host-to-docker-container)

## 1. 从容器内拷贝文件到宿主机

```shell
docker cp <containerID>:/file/path/of/container /host/path/target
```

注意

- 拷贝目录也是一样, 而且不需要`-r`参数.

- 另外目标容器未处于启动状态时也可以如此操作.

## 2. 从宿主机拷贝文件到容器内

### 2.1 直接从宿主机拷贝文件到容器物理存储系统



### 2.2 用`-v`参数挂载主机数据卷到容器内

注意

- 需要在容器启动时指定挂载参数.

- 若容器内目标目录已经存在, 挂载之后宿主机的目录将会将其**覆盖**

```shell
docker run -v /path/to/hostdir:/path/to/container 镜像名:标签
```

[宿主机路径:容器内路径]