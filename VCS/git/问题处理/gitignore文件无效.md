# gitignore文件无效

参考文章

1. [.gitignore无效, 不能过滤某些文件](http://blog.csdn.net/zhangjs0322/article/details/37658669)

利用.gitignore过滤文件, 如编译过程中的中间文件, 等等, 这些文件不需要被追踪管理.

问题描述:

在`.gitignore`添加`file1`文件, 以过滤该文件, 但是通过`git status`查看仍显示file1文件的修改状态.

原因分析:

`.gitignore`文件只对还没有加入版本管理的文件起作用, 如果之前已经用git把这些文件纳入了版本库并`commit`过, 就不起作用了.

解决方法:

需要在git库中删除该文件, 并更新(`git commit`). 然后再次`git status`查看状态, file1文件不再显示状态.
