# Linux内核编程(二)添加内核模块

一.最简示例

hello.c文件，需要包含内核代码树中的头文件

```c
//注意这些是内核代码树中的头文件，必要时可以包含绝对路径
#include<linux/init.h>
#include<linux/module.h>
MODULE_LICNESE("GPL");

static int hello_init(void)
{
    printk(KERN_ALERT " hello linux, I am general!n");
    return 0;
}

static void hello_exit(void)
{
    printk(KERN_ALERT "linux, I am gone!n");
}

module_init(hello_init);
module_exit(hello_exit);

```

Makefile文件

```makefile
path=/lib/modules/$(shell uname -r)/build
PWD=$(shell pwd)
obj-m:=hello.o//.o文件名应与模块.c文件相匹配
default:
        make -C $(path) M=$(PWD) modules
```

执行`make`编译之后会在目录下生成`.ko`模块文件。

执行`insmod`，`lsmod`，`rmmod`可完成模块的安装与卸载。

用`dmesg`查看模块安装卸载时的输出信息。
