# GDB调试golang程序

参考文章

1. [Golang源码探索(一) 编译和调试源码](https://www.cnblogs.com/zkweb/p/7777525.html)
    - lldb
    - golang示例编译
2. [GDB调试GO程序](http://blog.studygolang.com/2012/12/gdb%E8%B0%83%E8%AF%95go%E7%A8%8B%E5%BA%8F/)
    - runtime-gdb.py
    - runtime.Breakpoint(): 触发调试器断点

gdb可以通过`list 函数名称`查看函数指定函数前后的代码, 如下

```
(gdb) list main // 或是list 文件名(带后缀):main
1	#include <stdio.h>
2
3	void main()
4	{
5	    printf("hello world");
6	}
```

但是golang拥有package的概念, 函数符号名称的命名规则是`包名.函数名`, 例如主函数的符号名称是`main.main`, 这一点与C语言不同.
