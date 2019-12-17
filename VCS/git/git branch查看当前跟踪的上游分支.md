# git branch查看当前跟踪的上游分支

参考文章

1. [git如何查看当前分支是从哪个分支拉的？](https://ask.csdn.net/questions/274046)

```
$ git branch -v
  redis          0bd2c0f 完成调用file-push角色的redis搭建
* master         a04c7ef 小修正, 删除neo.tar.gz, files目录的内容不需要提交
  tomcat         c97f318 空目录添加.gitkeep文件以便提交
$ git branch -vv
  redis          0bd2c0f [origin/redis] 完成调用file-push角色的redis搭建
* master         a04c7ef [origin/master] 小修正, 删除neo.tar.gz, files目录的内容不需要提交
  tomcat         c97f318 [origin/tomcat] 空目录添加.gitkeep文件以便提交
```

`git branch -vv`可以查看本地分支对应的远程分支.

注意, 只能列出本地存在的分支.
