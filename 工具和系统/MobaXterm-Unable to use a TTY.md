# MobaXterm-Unable to use a TTY

参考文章

1. [解决windows终端执行“kubectl exec -it”显示"unable to use TTY"的问题](https://segmentfault.com/a/1190000021307781)
    - winpty
2. [kubectl exec -it fails with "Unable to use a TTY" and never get a prompt ](https://github.com/kubernetes/kubernetes/issues/37471)
    - xargs, winpty

- Win10: 企业版 1803 (操作系统版本 17134.228)
- MobaXterm: Professional Edition v12.4 Build 4248
- kubectl.exe: 1.16.2
- winpty: 0.4.3 cygwin 2.8.0 x64

## 问题描述

用了 MobaXterm, 在使用`edit`和`exec`子命令时无法执行...

```console
$ kubectl edit deploy xxx
error: there was a problem with the editor "/bin/vim"
```

```console
$ kubectl exec -it xxx /bin/bash
Unable to use a TTY - input is not a terminal or the right kind of file
```

按照参考文章1所说, 下载了`winpty`, 并在`kubectl`前面加上`winpty`.由于我的 MobaXTerm 设置了共享 windows 的 PATH 环境变量, 所以 winpty 解压后就放在D盘了, 然后把ta的bin目录路径加到了`Path`中, MobaXterm 需要重启一下才能使用 winpty.

但是仍然不行, 完全没反应...

```
$ winpty kubectl exec -it xxx /bin/bash
## 没输出
$ echo $?
127
```

有点萌b, 于是直接双击打开`winpty.exe`, 弹出如下窗口.

![图1-2 kubelet开放接口](https://gitee.com/generals-space/gitimg/raw/master/books/3a72d8c25a54027d97e32fc67660b8d8.png)

我看参考文章1中说到, 将`winpty`放到`/usr/bin`下面了. 我用`find`查了查, 发现在 MobaXterm 的`/bin/cygwin1.dll`有这个东西.

于是我真的把`winpty.exe`拷贝到`/usr/bin`下, 甚至把`winty`整个目录拷贝到`/usr/local`下(然后配置`PATH`), 结果还是不行.

## 解决方法

官网下载新新的正版...

不过 vim 的问题还是没解决, 我尝试过下载了`gvim`替换掉 MobaXTerm 自带的 vim , 但是无效, 放弃了.
