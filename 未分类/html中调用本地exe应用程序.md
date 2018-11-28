# html中调用本地exe应用程序

参考文章

1. [网页中打开exe](https://blog.csdn.net/leftfist/article/details/51785374)

2. [html网页调用本地exe程序的实现方法](https://blog.csdn.net/ilovecr7/article/details/46803711)

类似于网页上调用`QQ`, 打开`迅雷`等操作, 我们也可以自定义打开的程序. 但是目标程序必须要在注册表中有指定记录.

手动操作的方法在参考文章2中描述的比较详细, 就是在注册表`HKEY_CLASSES_ROOT`树中建立指定结构.

参考文章1中给出了对应的`reg`文件, 如下

[]中的为全路径键名, 其下的为字段值. `[HKEY_CLASSES_ROOT\MXIControl]`有两个值.

```
Windows Registry Editor Version 5.00

; 分号开头的为注释
; []中的为全路径键名, 其下的为字段值
; 值不能为中文, 会乱码

[HKEY_CLASSES_ROOT\MXIControl]
@="URL:MXIControl Protocol Handler"
"URL Protocol"=""

[HKEY_CLASSES_ROOT\MXIControl\DefaultIcon]
@="D:\\MxiControl-v0.5.2\\MxiControl.exe"

[HKEY_CLASSES_ROOT\MXIControl\shell]

[HKEY_CLASSES_ROOT\MXIControl\shell\open]

[HKEY_CLASSES_ROOT\MXIControl\shell\open\command]
@="D:\\MxiControl-v0.5.2\\MxiControl.exe %1"
```

然后在html上, 要调用目标`MxiControl`程序, 示例如下

```html
<!DOCTYPE html>
<html lang="en">
<head>
</head>
<body>
    <div>
        <a href="MxiControl://">喊话</a>
    </div>
</body>
</html>
```

------

补充:

事实上, 在调用MxiControl程序时, 程序本身可以打开, 但是在输入账号密码登录时出现未知错误(MxiControl是个人作品, 不是像qq, 迅雷那么靠谱), 但是双击图标打开是没问题的.

这种情况非常类似于在linux下在程序中调用命令行工具出错, 但在命令行直接执行时却完全没问题的场景. 因为在命令行执行时, bash加载了隐藏的环境变量, 而在程序中执行时没有. 

按照这个思路, 考虑写一个bat脚本来启动程序, 作为替代, 把这个bat脚本的路径写到注册表中. 结果如下

`.reg`文件

```
Windows Registry Editor Version 5.00

; 分号开头的为注释
; []中的为全路径键名, 其下的为字段值
; 值不能为中文, 会乱码

[HKEY_CLASSES_ROOT\MXIControl]
@="URL:MXIControl Protocol Handler"
"URL Protocol"=""

[HKEY_CLASSES_ROOT\MXIControl\DefaultIcon]
@="D:\\MxiControl-v0.5.2\\Start MxiControl.bat"

[HKEY_CLASSES_ROOT\MXIControl\shell]

[HKEY_CLASSES_ROOT\MXIControl\shell\open]

[HKEY_CLASSES_ROOT\MXIControl\shell\open\command]
@="D:\\MxiControl-v0.5.2\\Start MxiControl.bat"
```

`Start MxiControl.bat`文件, 和`MxiControl.exe`在同一个目录下.

```bat
echo start 

start /d "D:\MxiControl-v0.5.2" MxiControl.exe

exit
```

这样, 在html页面上点击按钮启动程序时会闪现一个cmd框(不过会马上消失), 但这样程序执行时就完全没有问题了.