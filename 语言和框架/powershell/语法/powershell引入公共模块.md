# powershell引入公共模块

参考文章

1. [PowerShell如何使用自定义公共函数](https://blog.csdn.net/flyliuweisky547/article/details/18565705)

网上一些文章说要在指定指定目录(`~/Documents/WindowsPowerShell/Modules`)创建子模块, 而且还要按照指定格式写模块简介, 总之很麻烦.

其实我只是想实现类似bash的`source`命令, 或是python的`import..from..`, 能把自己编写的公共模块独立出来.

参考文章1中给出了两种方法, 都有效.

```ps1
## 第一种是点号 .
. ~\Documents\WindowsPowerShell\toolkit.ps1
## 第二种是用`Import-Module`
Import-Module ~\Documents\WindowsPowerShell\toolkit.ps1
```

> 需要注意的是, 模块路径需要写成绝对的, 相对路径的话会出错.
