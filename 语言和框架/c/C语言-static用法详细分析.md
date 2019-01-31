# C语言static用法详细分析

C语言代码是以文件为单位来组织的，在一个源程序的所有源文件中，一个`外部变量`(注意不是局部变量)或者函数只能在同一个源程序中定义一次，如果有重复定义的话编译器就会报错。伴随着不同源文件变量和函数之间相互引用与相互独立的关系，产生了`extern`和`static`关键字。

这篇文章详细分析了static关键字在编写程序时的三类用法。

## 1. static全局变量

我们知道一个进程在内存中的布局如下图所示：

![](https://gitee.com/generals-space/gitimg/raw/master/a78d7c54006036f32147f6d23c049968.png)

其中`.text`段保存进程所执行的程序二进制文件，`.data`段保存进程所有已经初始化的全局变量，`.bss`段保存进程所有未经过初始化的全局变量。在进程的整个生命周期中，`.data`和`.bss`段内的数据是跟整个进程同生共死的，也就是在进程结束之后这些数据才会被销毁。

当一个进程的全局变量被声明为`static`类型后(中文名为**静态全局变量**)。静态全局变量和其他的全局变量的存储地点并无不同，都是在`.data`段(已初始化)或`.bss`段(未初始化)，但是这样的全局变量只在定义它的源文件中有效，其他源文件无法访问。所以，普通全局变量穿上static外衣后，它就变成了新娘，已心有所属，只能被定义它的源文件(新郎)中的变量或函数访问。

**示例程序**

`staticVar.h`文件：

```c
#include <stdio.h>
void printStr();
```

`staticVar.c`文件：

```c
#include "staticVar.h"
static char *strHello = "hello world";
void printStr()
{
    printf("%sn", strHello);
}
```

`static.c`文件(主程序文件)：

```
#include <stdio.h>
#include "staticVar.h"
int main()
{
    printStr();
    printf("%sn", strHello);
    return 0;
}
```

编译时报错如下(注意多文件编译方式)：

```
general@kali:~/Code/C$ gcc -g static.c staticVar.c -o static
static.c: In function ‘main’:
static.c:6:17: error: ‘strHello’ undeclared (first use in this function)
static.c:6:17: note: each undeclared identifier is reported only once for each function it appears in
```

即`strHello`变量未声明，可以将`static.c`文件中

```
printf("%sn", strHello);
```

行去掉，可以正常编译并运行，将打印出"hello world"。

**结论：**

上述实例中，`staticVar.c`中的`strHello`为一个静态全局变量，它可以被同一源文件中的`printStr()`函数访问，但不能被其他文件中的函数访问。
当然，你可以将对`strHello`的声明放在头文件中，这样就不会出错，因为编译器预处理时会将头文件中的内容复制到`.c`文件。

## 2. static局部变量

普通的局部变量在栈空间上分配，这个局部变量所在的函数被多次调用时，每次调用这个局部变量在栈上的位置都不一定相同。当然，局部变量也可以在堆上动态分配(这是C++里的东西了)，但是记得使用完这个堆空间后要释放掉。

static局部变量中文名为**静态局部变量**，它与普通的局部变量相比有如下几个区别：

1. 位置：静态局部变量被编译器放在全局存储区`.data`段(注意：不在`.bss`段，原因见3)，所以它虽然是在局部定义的，但是在程序的整个生命周期中存在。

2. 访问权限：**静态局部变量只能被其所在作用域内的变量或函数访问**。即虽然它会在程序的整个生命周期存在，但由于它是static的，所以不能被其他函数和源文件访问。

3. 值：静态局部变量如果没有被用户初始化，则会被编译器自动赋值为0，以后每次调用静态局部变量的时候都用上次调用后保存的值。这个比较好理解，每次函数调用静态局部变量的时候都修改它然后离开，下次读的时候从全局存储区读出的静态局部变量就是上次修改后的值。

**示例程序**

`staticVar.h`中的内容不变

`staticVar.c`文件的内容如下：

```c
#include "staticVar.h"
void printStr()
{
    int normal = 0;
    static int stat = 0;//the initialization will only excute one time
    printf("normal = %d----stat = %dn", normal, stat);
    ++ normal;
    ++ stat;
}
```

`static.c`文件：

```c
#include <stdio.h>
#include "staticVar.h"
int main()
{
    printStr();
    printStr();
    printStr();
    printStr();
    //this line will get a compile error, cause main() can't access the stat variable
    //printf("%dn", stat);
    return 0;
}
```

上述程序运行结果为：

```
normal = 0----stat = 0
normal = 0----stat = 1
normal = 0----stat = 2
normal = 0----stat = 3
```

**结论: **

1. 函数每次调用，普通局部变量都是重新分配，而静态局部变量保持上次调用的值不变。

2. `getStr()`函数中有对stat变量的初始化语句，根据运行结果看来，这条语句**只执行了一次**，称为静态变量的初始化, 以后调用`getStr()`时，不再对stat变量重新初始化。

3. 需要注意的是，由于`static`局部变量的这种特性，使得含静态局部变量的函数变得不可重入，即每次调用可能会产生不同的结果。这在多线程编程时可能成为一种隐患。需要多加注意。

## 3. static函数

C++面向对象编程中的private函数，私有函数只有该类的成员变量或成员函数可以访问。在C语言中，也有"private函数"，它就是接下来要说的static函数，完成面向对象编程中private函数的功能。

在程序中有很多个源文件的时候，你可能会让某个源文件只提供一些外界需要的接口函数，其他函数可能是**为了实现这些接口而编写，并不直接提供**。这些所谓的其他的函数你可能并不希望被外界(非本源文件)所访问，这时候就可以用static修饰这些"其他的函数"。

所以**static函数的有效域是当前源文件**，把它想象为面向对象中的private函数就可以了。

示例程序

`file1.h`文件：

```c
#include <stdio.h>
static int called();
void printStr();
file1.c文件：

#include "file1.h"
int called() //note : whether there is the static kerword or not doesn't matter
{
        return 6;
}
void printStr()
{
        int returnVal = called();
        printf("returnVal = %dn", returnVal);
}
```

`file.c`文件(程序主文件)：

```
#include "file1.h"
int main()
{
        int val;
        //this line will get a compile error, because main() can't access the called() static function
        //val = called(); 
        printStr();
        return 0;
}
```
static函数可以很好地解决不同源文件中函数同名的问题，因为一个源文件对于其他源文件中的static函数是不可见的。