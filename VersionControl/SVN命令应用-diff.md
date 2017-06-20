# SVN命令应用-diff

参考文章

[使用svn diff的-r参数的来比较任意两个版本的差异](http://blog.csdn.net/feliciafay/article/details/8962515)

- `svn diff`: 对比当前本地的工作拷贝文件(working copy)和缓存在.svn下的版本库文件的区别(即距离上次update/checkout做了哪些修改)

- `svn diff -rA`: 对比当前本地的工作拷贝文件(working copy)和任意版本A的差异(如`svn diff -r94239`对比本地的工作拷贝文件(working copy)和版本94239的差异)

- `svn diff -rA:B`: 对比任意历史版本A和任意历史版本B的差异(如`svn diff -r94239:94127`显示版本94127相对于版本94239的差异)

------

`svn diff`输出格式理解

```
$ svn diff  目标文件路径
Index: 目标文件路径
===================================================================
--- 目标文件路径	(revision 17190)
+++ 目标文件路径	(working copy)
@@ -1,7 +1,8 @@
 [dev:children]
+basic
 login
 logic
 chat
 edge
 callback
-config
+
+config
```

一对`@@`之间的文字表示发生变化的行的信息, 示例中表示修改前版本的`1-7`行变为了现在的`1-8`行, 修改前后不同的地方通过`-`与`+`作为前缀标识.

`+basic`是修改后多出来的, 前后没有相似的行, 所以可以确定这是新增行.

`-config`与`+config`则表示修改前后的区别了, 可以看出修改后`config`行前多了一个空行.

