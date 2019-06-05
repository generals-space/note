# git拉取和推送所有分支

参考文章

1. [git拉取和推送所有分支](https://www.jianshu.com/p/c196df31322b)
2. [Git push将本地版本库的分支推送到远程服务器上对应的分支](https://www.cnblogs.com/wuer888/p/7656523.html)
    - git push常用命令

将本地所有分支推送到远程

git push默认只推送当前分支到远程, 如需推送本地所有分支, 需要添加`--all`选项.

```
git push --all
```

拉取远程所有分支代码到本地

```
git branch -r | grep -v '\->' | while read remote; do git branch --track "${remote#origin/}" "$remote"; done
git fetch --all
git pull --all
```

虽然git pull也有`--all`选项, 但是好像不能达到我们想要的效果.

`git branch -r`是显示所有远程分支.

`grep -v '\->'`是过滤`origin/HEAD -> origin/master`这种记录.

git pull和git fetch的`--all`标记可以拉取所有已跟踪的分支, 在远程的最新提交.