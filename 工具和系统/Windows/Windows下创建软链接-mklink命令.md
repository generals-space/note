# Windows下创建软链接-mklink命令

参考文章

1. [Win7中的软链接详解（mklink命令）](https://blog.csdn.net/zht666/article/details/45917155)

linux中有软链接的概念，可以通过`ln`命令创建到目录或文件的软链接，软链接的好处就是可以让一个目录或文件有多个入口但保持单一物理位置，方便应用和管理。

之前一直苦于windows下没有类似的功能，导致有些地方很不方便，不过进入windows vista和win7时代后，这样的功能也被附带在windows中了，通过win7操作系统中的`mklink`命令就可以创建类似的软链接了。

> 快捷方式和软链接不是一回事, 快捷方式和同名目录或是文件可以同时存在, 但互不干涉, 软链接不可以.

**mklink的使用方法**

```
C:\>mklink /?
Creates a symbolic link.

MKLINK [[/D] | [/H] | [/J]] Link Target

        /D      Creates a directory symbolic link.  Default is a file
                symbolic link.
        /H      Creates a hard link instead of a symbolic link.
        /J      Creates a Directory Junction.
        Link    Specifies the new symbolic link name.
        Target  Specifies the path (relative or absolute) that the new link
                refers to.
```

翻译一下: `mklink 目标软链接路径 源路径`

`/D`: 创建目录的软链接(默认只是创建文件的, 当不加这个选项去创建目录的软链接时会出错)

`/H`: 创建`hard link`而不是`symbolic link`.

`/J`: ...我也不知道干啥用的, 不管它, 够用了.

> 注意: `mklink`是cmd命令, 在powershell中无法使用.