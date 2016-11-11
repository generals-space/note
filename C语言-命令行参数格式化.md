# C语言-命令行参数格式化

参考文章

[C程序对命令行参数的处理方法及实例](http://smilejay.com/2010/12/c-handle-options/)

[linux系统getopt函数详解](http://blog.csdn.net/wangpengqi/article/details/8182734)

C语言对于参数处理有两个函数: `getopt`与`getopt_long`

## 1. getopt

### 1.1 基本使用方法

```c++
#include <stdio.h>
#include <unistd.h>
int main(int argc,char *argv[])
{
    //选项名称, 虽然是int类型, 但实际上还是可以为字母.
    int ch;

    while((ch = getopt(argc, argv, "a:b::c")) != -1)
    {
        printf("当前选项ch: %c\n", ch);
        printf("当前选项参数optarg: %s\n", optarg);
        printf("下一个待解析的参数索引optind: %d\n", optind);
        switch(ch)
        {
            case 'a':
                printf("option a: '%s'\n", optarg);
                break;
            case 'b':
                printf("option b: '%s'\n", optarg);
                break;
            case 'c':
                printf("option c\n");
                break;
            default:
                printf("other option: %c\n",ch);
        }
    }
}
```

将上述代码保存为`getopt_test1.c`, 编译并执行. 如下.

```shell
$ gcc -o getopt_test1 -g ./getopt_test1.c 
$ ./getopt_test1 -a abc -c -b123
当前选项ch: a
当前选项参数optarg: abc
下一个待解析的参数索引optind: 3
option a: 'abc'
当前选项ch: c
当前选项参数optarg: (null)
下一个待解析的参数索引optind: 4
option c
当前选项ch: b
当前选项参数optarg: 123
下一个待解析的参数索引optind: 5
option b: '123'
```

`optarg`和`optind`是两个最重要的`external`变量。

`optarg`是指向参数的指针（当然, 只针对有参数的选项）；

`optind`是`argv[]`数组的索引, 指向下一个待解析参数的位置. 众所周知，`argv[0]`是函数名称，所有参数从`argv[1]`开始，所以`optind`被初始化设置指为1.

每调用一次`getopt()`函数，返回一个选项(ch)，如果该选项有参数，则`optarg`指向该参数。在传入的命令行参数中再也检查不到`getopt()`中`optstring`(函数第3个参数)中包含的选项时，返回`-1`。

同shell中的`getopts`命令一样, 函数`getopt()`认为`optstring`中，以'-'开头的字符（**注意！不是字符串！！**）就是命令行参数选项.

`optstring`中的格式规范如下：

1. 单个字符，表示选项.

2. 单个字符后接一个冒号':'，表示该选项后必须跟一个参数值。参数紧跟在选项后或者以空格隔开。该参数的指针赋给optarg。

3. 单个字符后跟两个冒号'::'，表示该选项后可以指定一个参数, 也可以不指定。指定参数时参数必须紧跟在选项后**不能以空格隔开**.

### 1.2 长选项

如果`optstring`中含有一个大写的'W'字符，后面带有一个冒号，也就是形如`W:`，则表示该"可选字符"是一个"长选项"，也就是说不是只有一个字符的"可选字符"。比如：`gcc -Wall  hello.c`. 要解析该类型参数，`getopt()`函数中的第三个参数`optstring`应该包含`W:all`，而且当解析到`-Wall`时optarg = all.  这一点也是GNU对getopt()函数的扩展。

```c++
#include <stdio.h>
#include <unistd.h>
int main(int argc,char *argv[])
{
    //选项名称, 虽然是int类型, 但实际上还是可以为字母.
    int ch;

    while((ch = getopt(argc, argv, "a:b::cW:install")) != -1)
    {
        printf("当前选项ch: %c\n", ch);
        printf("当前选项参数optarg: %s\n", optarg);
        printf("下一个待解析的参数索引optind: %d\n", optind);
        switch(ch)
        {
            case 'a':
                printf("option a: '%s'\n", optarg);
                break;
            case 'b':
                printf("option b: '%s'\n", optarg);
                break;
            case 'c':
                printf("option c\n");
                break;
            case 'W':
                printf("option W: '%s'\n", optarg);
                break;
            default:
                printf("other option: %c\n",ch);
        }
    }
}
```

将上述代码保存为`getopt_test2.c`, 编译并执行

```shell
$ gcc -o getopt_test2 -g ./getopt_test2.c 
$ ./getopt_test2 -a abc -c -Winstall
当前选项ch: a
当前选项参数optarg: abc
下一个待解析的参数索引optind: 3
option a: 'abc'
当前选项ch: c
当前选项参数optarg: (null)
下一个待解析的参数索引optind: 4
option c
当前选项ch: W
当前选项参数optarg: install
下一个待解析的参数索引optind: 5
option W: 'install'
```

如果有多个长选项, 可以写成这样`a:b::cW:install:W:clean`. 

但是长选项是没有办法接自己的参数的, 如`-Wclean all`, all无法被取到. 所以还是有限制的.

### 1.3 错误信息

如果`getopt()`函数在`argv[]`中解析到一个没有包含在`optstring`中的"可选字符"，它会打印一个错误信息，并将该"可选字符"保存在变量`optopt`中，并返回字符'?', 赋值给`ch`变量.

```c++
#include <stdio.h>
#include <unistd.h>
int main(int argc,char *argv[])
{
    //选项名称, 虽然是int类型, 但实际上还是可以为字母.
    int ch;

    while((ch = getopt(argc, argv, "a:b::cW:install")) != -1)
    {
        printf("当前选项ch: %c\n", ch);
        printf("当前选项参数optarg: %s\n", optarg);
        printf("下一个待解析的参数索引optind: %d\n", optind);
        switch(ch)
        {
            case 'a':
                printf("option a: '%s'\n", optarg);
                break;
            case 'b':
                printf("option b: '%s'\n", optarg);
                break;
            case 'c':
                printf("option c\n");
                break;
            case 'W':
                printf("option W: '%s'\n", optarg);
                break;
            default:
                printf("other option: %c\n",ch);
        }
        printf("optopt + %c \n", optopt);
    }
}
```

```
$ ./getopt_test -a abc -d  -c
当前选项ch: a
当前选项参数optarg: abc
下一个待解析的参数索引optind: 3
option a: 'abc'
optopt +  
./getopt_test: invalid option -- 'd'
当前选项ch: ?
当前选项参数optarg: (null)
下一个待解析的参数索引optind: 4
other option: ?
optopt + d 
当前选项ch: c
当前选项参数optarg: (null)
下一个待解析的参数索引optind: 5
option c
optopt + d 
```

当然，我们可以将变量`opterr`赋值为0，来阻止getopt()函数输出错误信息, 就是上面的'invalid option'这一句。在while循环之前对其赋值即可. 如

```
//选项名称, 虽然是int类型, 但实际上还是可以为字母.
int ch;
opterr = 0;
```

当`getopt()`函数的第三个参数`optstring`的第一个字符是':'时，很显然，这是由于少写了一个"可选字符"的缘故。此时，`getopt()`函数不返回'?'，而是返回':'来暗示我们漏掉了一个"可选字符".

## 2. getopt_long