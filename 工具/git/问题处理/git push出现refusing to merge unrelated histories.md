# git push出现refusing to merge unrelated histories

参考文章

1. [git无法pull仓库refusing to merge unrelated histories](http://blog.csdn.net/lindexi_gd/article/details/52554159)


```
fatal: refusing to merge unrelated histories
```

问题描述

OSChina初始化一个git仓库, 本地也使用`git init`初始化了一个工作目录, 并添加了文件, 做了一次提交.

不幸的是, 远程的git仓库有了一个readme文件, 相当于一次提交. 于是不管是push还是pull都提示了上述的错误. 的确, 它们基于不同的历史, 不能合并也是可以理解的...

然而更加不幸的是, 本地的提交没有父级提交(因为它是第一个提交), 所以也没有办法删掉, 尴尬了...

解决办法

这个问题出现在git的`2.9.2`版本, 需要使用`--allow-unrelated-histories`选项强制合并. `pull`与`merge`都有这个选项.
