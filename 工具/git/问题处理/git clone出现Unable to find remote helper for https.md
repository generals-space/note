# git clone出现Unable to find remote helper for https

参考文章

1. [“Unable to find remote helper for 'https'” during git clone](http://stackoverflow.com/questions/8329485/unable-to-find-remote-helper-for-https-during-git-clone)

情境描述:

源码安装的git, 2.9.3, 第一次执行`git clone`时出现如下报错

```
git clone https://git.oschina.net/generals-space/work-grapher.git ./grapher/
fatal: Unable to find remote helper for 'https'
```

原因分析与解决方法

这是因为缺少`curl-devel`库的缘故, 需要首先`yum install curl-devel`, 然后重新`configure`并编译安装.
