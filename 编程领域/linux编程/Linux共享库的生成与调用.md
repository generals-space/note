# Linux共享库的生成与调用

参考文章

1. [ Linux-（C/C++）动态链接库生成以及使用（libxxx.so）](http://blog.csdn.net/qq_33850438/article/details/52014399)

Linux: `.so`(Shared Object) 共享库

Windows: `.dll`(Dynamic Linked Library) 动态链接库.

地位等同.

so文件源程序不需要main函数, 就算有也不会被执行.

## 1. 共享库的创建

将下面的一个简单函数封装成为共享库.

```c
/*
 * 文件名: max.c
 * function: 取两者之前较大值
 * 这里没有依赖任何其他库所以就没有声明
 * */
int max(int a, int b)  
{  
    return a > b ? a : b;  
}
```

编译方法.

```
$ gcc -fPIC -c max.c 
$ gcc -shared -o libmax.so max.o
```

## 2. 共享库调用

使用链接库, 需要包含其对应的头文件（很正常, 我们平时使用C库函数也需要包含相关头文件）. 所以这里我们再写一个对应的max的头文件.

```c
/*
 * 文件名: max.h
 * */
#ifndef MAX_H_
#define MAX_H_

int max(int a, int b);

#endif
```

然后, 主调函数如下.

```c
/*
 * 文件名: main.c
 * */
#include<stdio.h>  
#include"max.h"  
int main(void)  
{  
    printf("max函数的调用结果: %d\n",max(1, 2));  
    return 0;  
}
```

main.c的编译方法为

```
$ gcc -L. -lmax -o main main.c
$ ll
-rwxr-xr-x. 1 root root 7853 Jun 25 07:49 libmax.so
-rwxr-xr-x. 1 root root 8548 Jun 25 08:05 main
-rw-r--r--. 1 root root  177 Jun 25 07:57 main.c
-rw-r--r--. 1 root root  125 Jun 25 07:48 max.c
-rw-r--r--. 1 root root   98 Jun 25 07:56 max.h
-rw-r--r--. 1 root root 1240 Jun 25 07:49 max.o
```

原文中这里就可以执行main程序了, 但我实验时, 出现了如下错误.

```
$ ./main 
./main: error while loading shared libraries: libmax.so: cannot open shared object file: No such file or directory
```

在执行main时, 依然找不到libmax.so, 这是因为`libmax.so`没有在系统默认路径中. 想像一下, 当一个第三方的程序main, 作者自己在编译时将依赖的共享库libmax.so放在了系统可以找到的位置, 编译成功, 发布release版. 而用户下载时没有下载libmax.so, 或者下载了但是放在一个非默认目录下(比较/tmp), main执行时肯定是找不到的呀. max()函数又没有被编译到main程序里面.

所以如果我们创建的libmax.so没有在系统共享库指定的默认路径下(默认为/lib与/usr/lib), 就需要手动指定共享库所在的目录. 

这里假设我们编写的代码都在/tmp目录下. 修改`/etc/ld.so.conf`文件(没有则手动创建). 添加`/tmp`为单独一行.

```
include ld.so.conf.d/*.conf
/tmp
```

然后执行`ldconfig`, 这个不会有任何输出. 之后再次执行main程序发现正常输出.

```
$ ./main 
max函数的调用结果: 2
```

成功.

## 3. C++调用

上面的调用方式是纯C语言情况下, 有时候, 第三方程序指定编译器使用g++(一般是Makefile中指定CC=g++吧). 这时编译main程序时会出错. 因为libmax这个库仅适合纯C使用, C++并不适合. 

```
$ g++ -L. -lmax -o main main.c 
/tmp/ccoAFApK.o: In function `main':
main.c:(.text+0xf): undefined reference to `max(int, int)'
collect2: error: ld returned 1 exit status
```

如果想编译一个可以供C++使用的共享库, max.h将要作出改变. 额外增加一句: `extern "C"`

```c
/*
 * 文件名: max.h
 * */
#ifndef MAX_H_
#define MAX_H_

// 这句话可理解为, 告诉编译器, 这个动态库（.so）是用C语言写的,   
// 需要用C语言链接方式来链接这个库, 这样就可以g++来编译了.
extern "C"
int max(int a, int b);

#endif
```

再次编译, 成功

```
$ g++ -L. -lmax -o main main.c 
```

但是问题来了, 纯C的编译方式并不支持`extern "C"`, 使用gcc编译修改后的max.h的main程序时也会得到错误.

```
$ gcc -L. -lmax -o main main.c 
In file included from main.c:5:0:
max.h:9:8: error: expected identifier or ‘(’ before string constant
 extern "C"
        ^
```

难道每次编译都要改来改去, 有没有同时适合C/C++链接库的方法呢?

答案是有的, 只需要改动头文件即可, 使用条件编译. C++有一个宏: `__cpluscplus`, 当用g++编译的时候, 就可以识别这个宏.

```c
/*
 * 文件名: max.h
 * */
#ifndef MAX_H_
#define MAX_H_

// 这句话可理解为, 告诉编译器, 这个动态库（.so）是用C语言写的,   
// 需要用C语言链接方式来链接这个库, 这样就可以g++来编译了.
#ifdef __cplusplus
extern "C"
{
#endif
    int max(int a, int b);
    //这里应该还可以写其他函数
#ifdef __cplusplus
}
#endif

#endif
```

ok, 这下gcc和g++都可以正常编译了.