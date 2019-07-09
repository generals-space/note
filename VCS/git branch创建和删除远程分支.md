# git branch创建和删除远程分支

参考文章

1. [3.5 Git 分支 - 远程分支](https://git-scm.com/book/zh/v2/Git-%E5%88%86%E6%94%AF-%E8%BF%9C%E7%A8%8B%E5%88%86%E6%94%AF)

1. 创建远程分支

创建远程分支前, 先在本地创建对应的本地分支, 比如`test`

```
git branch test
git checkout test
```

本地创建完成后, 推送到远程, 要使用`--set-upstream`选项. 

```
git push --set-upstream origin test
```

这条命令的意思是把本地的 **当前分支**推送为`origin`的test分支.

2. 删除远程分支

...删除远程分支用的不是`branch`子命令, 还是要用`push`

```
git push origin -d test
```