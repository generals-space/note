---
title: npm镜像源及代理设置
---


使用淘宝npm镜像, 配合使用其提供的`cnpm`命令.

1. [淘宝npm](https://npm.taobao.org/)

2. [NPM设置代理](https://my.oschina.net/deathdealer/blog/208919)

首先, 使用npm下载nodejs的依赖包时的确有速度慢的问题. 根据参考文章1, 安装依赖可以使用淘宝提供的`cnpm`作为替代. 

```
$ npm install -g cnpm --registry=https://registry.npm.taobao.org
```

之后可以使用`cnpm`执行install/list等操作(支持 npm 除了 publish 之外的所有命令)

```
$ cnpm install [name]
```

或者写一个命令别名到`~/.bashrc`, 如下

```bash
alias cnpm="npm --registry=https://registry.npm.taobao.org \
--cache=$HOME/.npm/.cache/cnpm \
--disturl=https://npm.taobao.org/dist \
--userconfig=$HOME/.cnpmrc"
```

之后使用npm安装依赖包时走的就是淘宝的镜像仓库了.

记得`source`生效哦.

------

但有些时候下载依赖包的操作不是由我们手动执行而是某些工具初始化步骤中自动执行的. 这时我们需要手动为npm设置仓库地址.

查看npm配置文件路径

```
$ npm config get userconfig
/root/.npmrc
```

查看全局配置文件路径

```
$ npm config get globalconfig
/usr/etc/npmrc
```

为npm设置代理

```
$ npm config set proxy http://server:port
$ npm config set https-proxy http://server:port
```

如果代理需要认证的话可以这样来设置.

```
$ npm config set proxy http://username:password@server:port
$ npm config set https-proxy http://username:pawword@server:port
```

如果代理不支持https的话需要修改npm存放package的仓库地址.

```
$ npm config set registry "https://registry.npm.taobao.org"
```

这样, 在npm配置文件中就会出现如下内容(我想可以直接手动设置)

```
registry=https://registry.npm.taobao.org
```