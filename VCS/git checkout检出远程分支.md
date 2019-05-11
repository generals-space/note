# git checkout检出远程分支

检出本地不存在的远程分支时, 需要使用如下语句

```
git checkout -b 本地分支名称 origin/远程分支名称
```

`checkout`子命令帮助文档

```
usage: git checkout [<options>] <branch>
   or: git checkout [<options>] [<branch>] -- <file>...

    -q, --quiet           suppress progress reporting
    -b <branch>           create and checkout a new branch
```

`-b <branch>`指定本地分支, 为`[<options>]`部分.

之后若不指定`<branch>`, 则在本地创建新分支.