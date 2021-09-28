# Python执行shell命令.5.subprocess.Popen读取实时输出

参考文章

1. [python获取实时命令行输出](https://www.cnblogs.com/pfeiliu/p/14584627.html)

通过`Popen()`执行的shell命令, 只有在其执行完成后才可以读取其`stdout`和`stderr`, 否则执行`subproc.stdout.read()`会被阻塞住.

但有些命令就是要实时读取其输出的, 要实现这个目的, 可以使用`iter`迭代器完成, 见参考文章1.

