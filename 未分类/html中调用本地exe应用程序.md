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

[HKEY_CLASSES_ROOT\MXIControl\shell]

[HKEY_CLASSES_ROOT\MXIControl\shell\open]

[HKEY_CLASSES_ROOT\MXIControl\shell\open\command]
@="E:\\MxiControl-v0.5.2\\MxiControl.exe %1"
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