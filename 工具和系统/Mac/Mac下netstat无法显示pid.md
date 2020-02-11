# Mac下netstat无法显示pid

参考文章

1. [mac下netstat](https://blog.csdn.net/pandafxp/article/details/53748031)

与linux不同, mac下的`netstat`中`-p`选项表示`protocol`, 不表示`pid`, 没有办法查看监听着某个端口的进程是哪一个.

参考文章1中说可以用`losf`命令, 貌似是mac自带的.

```
lsof -nP -iTCP:端口号 -sTCP:LISTEN
```

- `-n`: 不显示主机名
- `-P`: 不显示端口的英文名称比如80是http
- `-i <条件>`: 列出符合条件的进程. (4、6、协议、:端口、 @ip)

```bash
## @function: netstat -nlp | grep 端口号
## $1:        端口号
function netstatL
{
    lsof -nP -iTCP:$1 -sTCP:LISTEN
}
## @function: netstat -nap | grep 端口号
## $1:        端口号
function netstatA
{
    lsof -nP -iTCP:$1
}
## @function: netstat -nap | grep PID
## $1:        目标进程pid
function netstatP
{
    lsof -nP -iTCP -a -p $1
}
```
