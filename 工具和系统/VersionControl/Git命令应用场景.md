# Git命令应用场景

## 1. 环境配置

### 1.1 用户名/邮箱配置

### 1.2 删除指定配置项

```
$ git config --global user.name=general
```

## 2. 工作区, 暂存区, 本地版本库, 远程版本库相互覆盖

```
## 使用HEAD分支中的内容刷新暂存区内存, add的文件会丢失, 工作区不变. 其实也隐含着使用HEAD分支最近的提交刷新当前分支的最近提交.
git reset HEAD
## 将当前分支的游标指向上一个提交, 即父提交. 放弃了当前提交. 会刷新当前工作区与暂存区, 已经提交过的修改都会丢失(可以通过reflog找回)
git reset --hard HEAD^
```

```
## 查看对象类型, 是blob, commit还是tree或tag.
git cat-file -t 51519c7
```

```
## 查看提交历史
git log
## 查看修改历史(自从git工程创建或克隆, 每一次操作都会记录, commit, checkout...好像没有add)
git reflog
```

## 3.

追踪分支

```
$ git branch --set-upstream-to=origin/master master
```

