# npm build时Killed

参考文章

1. [`npm install` ends with “Killed”](https://stackoverflow.com/questions/38127667/npm-install-ends-with-killed)

在低配阿里云服务器(1核1G)上执行`npm build`编译打包vue工程, 显示被killed掉了.

参考文章1中指出是因为内存不足, 且未设置swap的原因. 

用`free`查了下, 果然swap空间为0, 后来手动创建了1G的swap, 再次`build`, 还是不行, 而且swap空间占用为0.

然后查看`vm.swappiness`内核参数, 其值为0, linux默认此值为60才对, 修改后再次尝试`build`成功.