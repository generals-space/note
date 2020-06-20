# go mod - no dependencies to vendor

## 问题描述

go mod `download`和`tidy`没问题, 但是`go.mod`文件中没有依赖项目, 且无法生成`go.sum`文件, `vendor`命令也没有办法生成`vendor`目录.

```console
$ go mod vendor
go: no dependencies to vendor
```

直接使用`go get -v`希望能重写`go.mod`文件, 但是会执行失败.

```console
$ go get -v
Fetching https://goproxy.io/github.com/golang/protobuf/proto/@v/list
Fetching https://goproxy.io/github.com/go-kit/kit/endpoint/@v/list
...省略
build gokit: cannot load github.com/go-kit/kit/endpoint: cannot find module providing package github.com/go-kit/kit/endpoint
```

直接在工程目录外使用`go get -v`单独下载也不行.

```
$ go get -v github.com/go-kit/kit/endpoint
github.com/go-kit/kit (download)
# cd .; git clone https://github.com/go-kit/kit /usr/local/gopath/src/github.com/go-kit/kit
fatal: unable to access 'https://github.com/go-kit/kit/': Encountered end of file
package github.com/go-kit/kit/endpoint: exit status 128
```

前面的情况还好, 但是`go get`下载出问题肯定是因为代理方面有问题了, 我尝试了下把`GOPROXY`改成`direct`, 也试着改了改 http_proxy 等环境变量的设置, 最后发现好像是因为系统出了点问题, 把虚拟机重启一下, 好了...（⊙.⊙）
