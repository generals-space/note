# 编写参数可变的函数

<!tags!>: <!c!> <!va_start!>

参考文章

1. [va_start和va_end使用详解](http://www.cnblogs.com/hanyonglu/archive/2011/05/07/2039916.html)

2. [linux下利用va_start编写参数可变的函数](http://ju.outofmemory.cn/entry/154069)

3. [C library macro - va_start()](https://www.tutorialspoint.com/c_standard_library/c_macro_va_start.htm)

在编写C程序时，会遇到`printf`这样形参可变的函数，可能会好奇这是怎么做到的。

一般我们自己要新写一个函数时，都是在头文件中提前声明好函数，规定函数参数的个数和类型。然后在`.c`源码文件中，去定义这个函数的源码，在源码中调用这些函数参数。

我们如果有需求要编写参数列表可变的函数时，怎么办呢？

在C中，当我们无法列出传递函数的所有实参的类型和数目时,可以用省略号指定参数表.

```c
#include <stdio.h>

int printf(const char *format, ...);
int fprintf(FILE *stream, const char *format, ...);
int sprintf(char *str, const char *format, ...);
```

这种方式和我们以前认识的不大一样，但我们要记住这是C中一种传参的形式，在后面我们就会用到它。

> ANSI标准形式的声明方式，括号内的省略号表示可选参数.

## 1. 函数参数的传递原理

下面是 <stdarg.h> 里面重要的几个宏定义如下：

```c
typedef char* va_list;
void va_start(va_list ap, prev_param); /* ANSI version */
type va_arg(va_list ap, type);
void va_end(va_list ap);
```

`va_list`是一个字符指针(char *)，可以理解为指向**当前参数**的一个指针，取参必须通过这个指针进行。

以下是`va_*`系函数的使用流程

1. 在调用参数表之前，定义一个`va_list`类型的变量，假设为`ap`；

2. 然后应该对`ap`进行初始化，让它指向可变参数表里面的第一个参数，这是通过`va_start`来实现的，第一个参数是`ap`本身，第二个参数是传入的变参表前面紧挨着的一个变量,即`...`之前的那个确定的参数；

3. 然后是获取参数，调用`va_arg`，它的第一个参数是`ap`，第二个参数是要获取的参数的指定类型，然后返回这个指定类型的值，并且把`ap`的位置指向变参表的下一个变量位置；

4. 获取所有的参数之后，我们有必要将这个`ap`指针关掉，以免发生危险，方法是调用`va_end`，它将输入的参数`ap`置为 NULL. 应该养成获取完参数表之后关闭指针的习惯. 说白了，就是让我们的程序具有健壮性。通常`va_start`和`va_end`是成对出现。

## 2. 示例1

```c
/*va_test1.c*/
#include <stdio.h>
#include <string.h>
#include <stdarg.h>

// 函数声明, 至少需要一个确定的参数, 注意括号内的省略号
int demo(char *msg, ...);

void main(int argc, char *argv[])
{
    demo("DEMO", "This", "is", "a", "demo!", "");
    return ;
}

//ANSI标准形式的声明方式，括号内的省略号表示可选参数
//注意, 使用va_*系函数需要至少第一个参数为确定项
int demo(char *msg, ...)
{
    va_list arg_list;
    int arg_index = 0;
    char *para;
    // 初始化, va_start的第2个参数需要是可变参数前的一个参数, 
    // 所以必需提供至少一个确定的参数.
    va_start(arg_list, msg);

    //while循环处理所有传入参数, 以最后的空参数""为结尾.
    while(1)
    {
        //va_arg第2个参数是下一个要获取的参数的类型.
        para = va_arg(arg_list, char*);
        //遇到空参数表示已到结尾.
        if(strcmp(para, "") == 0) break;
        printf("Parameter #%d is: %s\n", arg_index, para);
        arg_index ++;
    }
    va_end(arg_list);
    return 0;
}

```

编译并执行它.

```
$ gcc -g -o va_test1 va_test1.c 
$ ./va_test1 
Parameter #0 is: This
Parameter #1 is: is
Parameter #2 is: a
Parameter #3 is: demo!
```

参考文章3中有另一个示例, 用的是使用第一个参数指定传入的参数个数, while循环中以这个值作为循环次数, 而且只是简单的加法, 具有特殊性.

## 3. 示例2

```c
/*va_test1*/
#include <stdio.h>
#include <dlfcn.h>
#include <stdlib.h>
#include <stdarg.h>

int display(const char *format, ...);

void main(int argc, char *argv[])
{
    char *name = "general";
    int age = 21;
    display("My name is %s, and age is %d\n", name, age);
}

int display(const char *format, ...)
{
    va_list list;
    char *parg;

    va_start(list, format);
    vasprintf(&parg, format, list);
    va_end(list);

    //...看看vasprintf得到了什么
    printf("%s\n", parg);

    free(parg);
}
```

```
$ gcc -g -o va_test2 va_test2.c
$ ./va_test2 
My name is general, and age is 21
```