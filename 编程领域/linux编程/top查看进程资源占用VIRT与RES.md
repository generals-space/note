# top查看进程资源占用VIRT与RES

<!tags!>: <!虚拟内存!> <!物理内存!> <!共享内存!>

参考文章

1. [linux top命令看到的实存(RES)与虚存(VIRT)分析](https://www.cnblogs.com/xudong-bupt/p/8643094.html)
    - 解释了top命令中`VIRT(虚拟内存)`, `RES(常驻内存)`与`SHR(共享内存)`的含义
    - 配有2个c语言示例, 展示了申请内存, 与写入数据后`VIRT`和`RES`的值的变化. 示例简单清晰, 值得一看.
2. [linux top命令VIRT,RES,SHR,DATA的含义](https://javawind.net/p131)
    - 介绍了top命令中`VIRT`, `RES`与`SHR`和`DATA`字段的含义.
    - 介绍了top命令所有可用字段的含义, 相同于top命令的高级使用手册
3. [Linux交换空间（swap space）](https://segmentfault.com/a/1190000008125116)
    - 关于swap的原理和作用讲述得比较清晰易懂
    - 参考文章1,2都有提到swap out操作, 可以阅读本文理解其过程.

按照参考文章2所说, 可使用`-f`选项额外显示`Swap`, `Code`, `Data`, `Used`列, 与本文的目的大致相关. 

- `VIRT`: `virtual memory usage` 虚拟内存, 等同于`ps aux`结果中的`VSZ`列. 
    - 虚拟内存会映射到swap空间(想想win下虚拟内存的设置, linux也是一样的道理)
    - VIRT=SWAP+RES
    - `SWAP`是虚拟内存被换出的大小, 但是虚拟内存可以很大, 远远超过swap的大小. 
    - 所以被swap out的空间应该不会是malloc向系统申请的空间, 而是当前进程不常用其他数据和库.
- `RES`: `resident memory usage` 常驻内存, 等同于`ps aux`结果中的`RSS`列. RES=CODE+DATA
    - 由于被称为常驻内存, 所以其与swap是不相交的, 但ta们的内容并非是不相容的. 
- `SHR`: `shared memory` 共享内存.
- `SWAP`: 进程使用的虚拟内存中, 被换出的大小.
- `CODE`: 可执行代码占用的物理内存大小
- `DATA`: 可执行代码以外的部分(数据段+栈)占用的物理内存大小. 
- `USED`: RES+SWAP.

有点乱, 这里保留疑问, 日后再来研究吧<???>

参考文章1给出的第一个示例使用了C++的语法: `new`和`delete`, 这里我改成纯C语言的实现.

`heap.c`

```c
#include <stdio.h>
#include <string.h>

int main()
{
    int test = 0;
    // 分配512M, 未使用
    char* p;
    p = malloc(1024 * 1024 * 512);
    scanf("%d", &test); //等待输入

    // 使用10M
    memset(p, 0, 1024 * 1024 * 10);
    scanf("%d", &test); //等待输入

    // 使用50M
    memset(p, 0, 1024 * 1024 * 50);
    scanf("%d", &test); //等待输入
    return 0;
}
```

编译并执行

```
gcc -g -o heap heap.c
./heap
```

使用`top`命令查看.

在堆上分配512M空间后

```
  PID USER        VIRT    RES    SHR %CPU %MEM COMMAND SWAP   CODE    DATA   USED
 3340 root      528504    344    272  0.0  0.0 heap       0      4  524476    344
```

使用10M后

```
  PID USER        VIRT    RES    SHR %CPU %MEM COMMAND SWAP   CODE    DATA   USED
 3340 root      528504  10636    408  0.0  1.0 heap       0      4  524476  10636
```

使用50M后

```
  PID USER        VIRT    RES    SHR %CPU %MEM COMMAND SWAP   CODE    DATA   USED
 3340 root      528504  51556    408  0.7  5.1 heap       0      4  524476  51556
```

## 2. 在栈上分配

参考文章1中的第二个示例在栈上申请了20M空间, 但是在CentOS7服务器版上, `ulimit -s`所规定的栈空间为8M, 所以在编译完成然后执行时出现了`Segmentation fault (core dumped)`.

我尝试过为`p`数组分配8M内存, 然后...果然失败了. 因为栈空间不只包含这个数组, 还有其他代码, 所以我试着分配了7M内存, 可以了...

为了数据展示更友好, 这里为`p`数组分配5M内存.

```c
#include <stdio.h>
#include <string.h>

int main()
{
    int test = 0;
    // 20M栈, 未使用
    // char p[1024 * 1024 * 20];
    char p[1024 * 1024 * 5];
    scanf("%d", &test);    //等待输入

    // 使用10M
    // memset(p, 0, 1024 * 1024 * 10);
    memset(p, 0, 1024 * 1024 * 3);
    scanf("%d", &test);    //等待输入
    return 0;
}
```

编译并执行

```
gcc -g -o stack stack.c
./stack
```

在栈上分配5M空间后

```
  PID USER        VIRT    RES    SHR %CPU %MEM COMMAND SWAP   CODE    DATA   USED
12855 root        9216    348    272  0.0  0.0 stack      0      4    5188    348
```

使用3M后

```
  PID USER        VIRT    RES    SHR %CPU %MEM COMMAND SWAP   CODE    DATA   USED
12855 root        9216   3512    384  0.0  0.3 stack      0      4    5188   3512
```
