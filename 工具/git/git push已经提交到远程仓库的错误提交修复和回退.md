# Git命令应用-已经提交到远程仓库的错误提交修复和回退

参考文章

1. [GIT_已经提交到远程仓库的错误提交怎么修复和回退](https://blog.csdn.net/gzgengzhen/article/details/76446082)

已经推到远程的分支上分某次提交发现其中有错误的操作(或commit信息), 需要回退进行更改(一般是比较愚蠢的, 不想被人抓住把柄的那种).

步骤: 

1. 重置(或是重置到某一个commit), reset的标记根据实际情况选择`hard`, `soft`, 或是`mixed`.

```
git reset HEAD^
```

2. 修改并重新commit

```
git commit -m "New commit message"
```

3. 强制上传

```
git push --force
```
