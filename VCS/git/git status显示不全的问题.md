# git status显示不全的问题

参考文章

1. [Showing Untracked Files In Status After Creating New Directories In Git](https://simpleit.rocks/git/tutorials/showing-untracked-files-in-status-after-creating-new-directories-in-git/)

某次向仓库中新增了300多个文件, 但是在命令行中使用`git stauts`查看到的结果列表只有50多个.

按照参考文章1中所说, 可以使用`--untracked-files=all`选项以显示所有发生变动的文件, 如下

```
git status --untracked-files=all
```

注意: `--untracked-files all`不行

其实我之前试过从 man 手册中找相关参数, 也发现了这个选项, 其对应的短选项为`-u`, 如下

```console
$ man git-status
       -u[<mode>], --untracked-files[=<mode>]
           Show untracked files.

           The mode parameter is used to specify the handling of untracked files. It is optional: it defaults to all, and if specified, it must be stuck to the option (e.g.  -uno,
           but not -u no).

           The possible options are:

           o   no - Show no untracked files.

           o   normal - Shows untracked files and directories.

           o   all - Also shows individual files in untracked directories.
```

但是尝试了如下几种都是错的

```
-u all
-u o=all
```

正确的应该是`-uall`, 如下

```
git status -uall
```

...fuck
