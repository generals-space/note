# go mod代理-GOPROXY

参考文章

1. [go mod代理和小技巧](https://www.cnblogs.com/xdao/p/go_mod.html)

```bash
export GOPROXY=https://goproxy.io
export GOPROXY=https://mod.gokit.info
export GOPROXY=https://mirrors.aliyun.com/goproxy/
export GOPROXY=https://goproxy.cn ## (这个有时不行)
export GOPROXY=https://proxy.golang.org
```

## FAQ

### 1. 

```
zip: not a valid zip file
```

**问题描述**

使用了goproxy代理后, 执行`go mod download`时会出现上述问题.

**解决方法**

代理服务器缓存的zip 错误, 可以临时关闭代理, `export GOPROXY=''`
