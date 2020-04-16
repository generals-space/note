# git checkout切换到指定tag

参考文章

1. [git切换到某个tag](https://blog.csdn.net/DinnerHowe/article/details/79082769)

git clone 整个仓库后使用, 以下命令就可以取得该 tag 对应的代码了.  

```
git checkout tag名称
```

但是, 这时候 git 可能会提示你当前处于一个`detached HEAD`状态. 

因为 tag 相当于是一个快照, 是不能更改它的代码的. 

如果要在 tag 代码的基础上做修改, 你需要一个分支:  

```
git checkout -b 新branch名称 tag名称
git checkout -b 新branch名称 commitID
```

这样会从 tag 创建一个分支, 然后就和普通的 git 操作一样了. 
