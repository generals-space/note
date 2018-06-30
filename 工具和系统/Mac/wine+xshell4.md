# wine+xshell4

1. [Mac OS X上安装Xshell](https://www.linuxidc.com/Linux/2015-04/116463.htm)

2. [Installing WineHQ packages](https://wiki.winehq.org/MacOS)

3. [xshell的注册码](http://www.xuebuyuan.com/2126914.html)

Xshell 4 注册码： 690313-111999-999313

Xftp 4 注册码：101210-450789-147200

`WineBottler`安装不成功, 启动不了.

使用wine可以装xshell4, 但是找不到在哪里启动. 后来写了个脚本来执行.

```bash
function xshell(){
    ## 设置wine的环境变量
    test "$?BASH_VERSION" = "0" || eval 'setenv() { export "$1=$2"; }';
    setenv PATH "/Applications/Wine Stable.app/Contents/Resources/start/bin:/Applications/Wine Stable.app/Contents/Resources/wine/bin:$PATH";

    wine C:\\Program\ Files\ \(x86\)\\NetSarang\\Xshell\ 4\\Xshell.exe
}
```

太丑了, 而且没有全局快捷键, 放弃.