# npm安装及卸载依赖包

参考文章

1. [npm install 安装某个指定的版本](https://blog.csdn.net/idomyway/article/details/80601864)

2. [npm全局模块卸载及默认安装目录修改](https://blog.csdn.net/baidu_28479651/article/details/53790550)

## 安装指定版本的包

```
npm install webpack@1.0.0
npm install webpack@1.x.x
```

## 卸载全局包

有的时候我们把node的一些工具命令安装在了全局, 比如webpack, 卸载的时候就不能单纯地删除`node_modules`目录了, 这时可以使用`npm uninstall -g webpack`完成卸载.
