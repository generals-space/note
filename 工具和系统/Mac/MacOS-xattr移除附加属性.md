# MacOS-xattr移除附加属性

xattr -c 文件名称

```
$ ll
drwxrwxr-x@  3 general  staff         96  8  6 05:30 Appium.app
$ xattr -c ./Appium.app
$ ll
drwxrwxr-x   3 general  staff         96  8  6 05:30 Appium.app
```

> 表示附加属性的`@`符号已经被移除.

不过这只能针对单个文件, 如果目标是一个目录, 且需要将目录下所有子文件都移除附加属性, 需要使用`-r`递归选项.

```
xattr -rc ./Appium.app
```
