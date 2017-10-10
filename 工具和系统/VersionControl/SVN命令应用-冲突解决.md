# SVN命令应用-冲突解决

参考文章

[SVN：冲突解决 合并别人的修改](http://www.letuknowit.com/archives/svn-conflict-resolution/)

## 1.

```
root@localhost:/home/kris/calc/trunk# svn up
Conflict discovered in 'main.c'.
Select: (p) postpone, (df) diff-full, (e) edit,
        (mc) mine-conflict, (tc) theirs-conflict,
        (s) show all options:
```

各个选项的含义

```
(p) postpone          暂时推后处理，我可能要和那个和我冲突的家伙商量一番
(df) diff-full        把所有的修改列出来，比比看
(e) edit              直接编辑冲突的文件
(mc) mine-conflict    如果你很有自信可以只用你的修改，把别人的修改干掉
(tc) theirs-conflict  底气不足，还是用别人修改的吧
(s) show all options  显示其他可用的命令
```