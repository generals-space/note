# go mod问题处理malformed file path

在将flannel的依赖管理从glide转换成go mod时, 遇到了一些问题.

`glide`配置yaml文件可以被go mod转换使用, 自动生成`go.mod`文件(与`go dep`一样). flannel工程生成的`go.mod`其中有一句

```
require (
	github.com/docker/distribution v0.0.0-20160419170423-cd27f179f2c1 // indirect
)
```

在执行`go mod download`时出现如下错误(其他依赖正常).

```log
$ go mod download
-> unzip /usr/local/gopath/pkg/mod/cache/download/github.com/docker/distribution/@v/v0.0.0-20160419170423-cd27f179f2c1.zip: malformed file path "contrib/docker-integration/generated_certs.d/localregistry:5440/ca.crt": invalid char ':'
github.com/docker/distribution@v0.0.0-20160419170423-cd27f179f2c1: unzip /usr/local/gopath/pkg/mod/cache/download/github.com/docker/distribution/@v/v0.0.0-20160419170423-cd27f179f2c1.zip: malformed file path "contrib/docker-integration/generated_certs.d/localregistry:5440/ca.crt": invalid char ':'
```

在go的官方issue有相关的问题, 但没有明确的解决方案.

后来我按照错误提示中, 到`gopath/pkg/mod/cache`目录下找到出问题zip包, 该目录有如下文件.

```log
/usr/local/gopath/pkg/mod/cache/download/github.com/docker/distribution/@v $ ll
总用量 2060
-rw------- 1 root root    110 1月  15 14:13 list
-rw-r--r-- 1 root root      0 12月 31 14:25 list.lock
-rw------- 1 root root     78 1月  15 14:13 v0.0.0-20160419170423-cd27f179f2c1.info
-rw-r--r-- 1 root root      0 1月  15 14:14 v0.0.0-20160419170423-cd27f179f2c1.lock
-rw------- 1 root root     38 1月  15 14:13 v0.0.0-20160419170423-cd27f179f2c1.mod
-rw------- 1 root root 730330 1月  15 14:14 v0.0.0-20160419170423-cd27f179f2c1.zip
-rw------- 1 root root     47 1月  15 14:14 v0.0.0-20160419170423-cd27f179f2c1.ziphash
```

但是`gopath/pkg/mod/github.com/docker`下并没有该版本(v0.0.0)的`distribution`工程(flannel工程下`vendor`目录下当前也没有), 但是有其他版本.

```log
/usr/local/gopath/pkg/mod/github.com/docker $ ll
总用量 44
dr-x------ 17 root root 4096 12月 31 14:59 distribution@v2.6.0-rc.1.0.20170726174610-edc3ab29cdff+incompatible
dr-x------ 20 root root 4096 1月   3 12:11 distribution@v2.7.1+incompatible
dr-x------ 35 root root 4096 1月   3 12:10 docker@v0.7.3-0.20190327010347-be7ac8be2ae0
dr-x------ 35 root root 4096 12月 31 14:59 docker@v1.13.1
```

就是说, `go mod download`其实已经把依赖包下载并缓存下来了, 但是解压出错, 所以pkg目录下没有对应的工程包.

出错原因是看来是unzip无法解压路径中带有特殊字符的文件/目录, 不过linux下的文件名称几乎没有非法字符, 所以最开始并不愿意相信这个猜想.

我把`cache`目录下的`v0.0.0-20160419170423-cd27f179f2c1.zip`下载到本地(MacOS)解压, 解压开找到`localregistry:5440`, 发现在MacOS的Finder中, 目标名称变成了`localregistry/5440`, 但在命令行, 还是`localregistry:5440`.

然后把解压出来的`distribution@v0.0.0-20160419170423-cd27f179f2c1`目录放到`gopath/pkg/mod/github.com/docker`目录下, 重新执行`go mod download/vendor`成功.

看来确实是`unzip`的锅.
