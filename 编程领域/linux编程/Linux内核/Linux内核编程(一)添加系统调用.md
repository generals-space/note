# linux内核编程(一)添加系统调用

## 1. .环境

系统：Ubuntu10.04

内核：2.6.32

## 2. .步骤

### 2.1 添加系统调用函数声明

在`linux-2.6.32.65/include/linux`目录中，找到`syscalls.h`, 在这个文件末尾添加：

```c
//返回值声明为int类型时编译会出错，有待研究
asmlinkage long sys_myadd(int x, int y);
```

### 2.2 添加系统调用函数定义

在`linux-2.6.32.65/kernel`目录下，找到`sys.c`,在里面添加入系统调用的实现程序。

在linux-2.6.32里面，跟以前的版本不一样，这里用到了宏`SYSCALL_DEFINE(n)`(`n`为数字，意为系统调用函数的参数个数，如不需要参数则为0)对系统调用进行了封装。难怪网上的都在说找不到系统调用的定义。

```c
//注意函数名去掉了前缀sys_,并且注意参数传递方式
SYSCALL_DEFINE2(myadd, int, x, int, y)
{
    return x + y;
}
```

### 2.3 在系统调用表中注册

在`linux-2.6.32.2/arch/x86/kernel`目录中，找到`syscall_table_32.S`，在这个文件的最后一行，添加：`.int sys_myadd`

```c
.int sys_myadd
```

### 2.4 添加系统调用入口参数(系统调用号)

在`linux-2.6.32.2/arch/x86/include/asm`目录下，找到`unistd_32.h`，在这个文件的 `#define NR_syscalls 337`前面加：`#define __NR_myadd 337`，同时把`NR_syscalls`的值改成`338`.

```c
#define __NR_myadd 337
#ifdef __KERNEL__

#define __NR_syscalls 338
```

> 说明：NR_syscalls相当于系统调用表边界，所有系统号都得小于它。

### 2.5 编译内核并重启

### 2.6 验证系统调用

在C语言中调用`open`，`write`等linux系统调用需要包含`<sys/types.h>`和`<sys/stat.h>`等，应该是C库本身对系统调用进行了封装，才能直接用函数名称进行调用。而我们自行编写的系统调用没有存在于C标准函数库中，所以只能通过`syscall()`函数传入调用号调用函数。

`syscall()`的第一个参数为系统调用号，后面的参数为系统调用的参数。

```c
#include <unistd.h>
#include <stdio.h>
int main(){
    if(syscall(337))
        printf("system call succeedn");
    long result = syscall(337, 10, 23);//syscall()的第一个参数为系统调用号，后面的参数为系统调用的参数
    printf("the result is %dn", (int)result);//注意类型转换
    return 0;
}
```

提示：可以用`time`命令测试代码的执行时间以验证系统调用与用户函数的调用的效率差异，一般来说系统调用的开销要远大于用户函数调用的开销。

## 3. 一个有用的系统调用

遍历所有**内核进程**，声明如下：

```c
SYSCALL_DEFINE0(mycall)
{
    struct task_struct *p;
    printk("***************************************n");
    printk("--------the output of mycall-----------n");
    printk("***************************************n");
    printk("%-20s %-6s %-6s %-20sn", "Name", "Pid", "State", "ParentName");//加负号'-'可以改变对齐方式
    for(p = &init_task; (p = next_task(p)) != &init_task)
    {
        printk("%-20s %-6s %-6s %-20sn", p->comm, p->pid, p->state, p->parent->comm);
    }
    return 1;
}
```

验证代码

```c
#include <unistd.h>
#include <stdio.h>
int main()
{
    if(syscall(338))//338是mycall的调用号
        printf("okn");
    else
        printf("failedn");
    return 0;
}
```

编译运行，输出ok之后，用`dmesg`命令查看结果