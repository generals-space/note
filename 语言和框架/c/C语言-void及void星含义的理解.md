# C语言-void及void*含义的理解

参考文章

[void及void指针含义的深刻解析](http://blog.csdn.net/geekcome/article/details/6249151)

在2.6.32版本的内核源码中/fs/nfs/write.c文件中的nfs_writedata_alloc()函数中有这样一段代码：

```c
struct nfs_write_data *p = mempool_alloc(nfs_wdata_mempool, GFP_NOFS);
```

而关于`memool_alloc()`，则在`/include/linux/mmpool.h`中有如下定义声明：

```c
extern void * mempool_alloc(mempool_t *pool, gfp_t gfp_mask);
```

也就是说一个类型为`struct nfs_write_data`的指针p，被赋予一个void*类型的值

 怎么回事？不是说不同类型的变量不能随意赋值么？为什么没有类型转换？

## 1. Void的含义

void的字面意思是"无类型"，而void*则为"无类型指针"。

void本身几乎只有"注释"和限制程序的作用，而void*变量则可以指向任何类型的数据。

## 2. 使用方法

### 2.1 关于void

#### 2.1.1 如果函数无返回值，则应声明为void类型，不可省略

在C语言中，凡不加返回值类型的函数，就会被编译器作为返回值为整型处理(注意，很多人以为是void类型)，如：

```c
#include <stdio.h>
add(int a, int b){
        return a + b;
}
void main(){
        printf("2 + 3 = %dn", add(2, 3));
}
```
程序输出为：

```
2 + 3 = 5
```
说明不指定返回值类型的函数的确被用作int函数。

> 林锐博士《高质量C/C++编程》中提到："C++语言有很严格的类型安全检查，不允许上述情况（指函数不加类型声明）发生"。可是编译器并不一定这么认定，譬如在VisualC++6.0中上述add函数的编译无错也无警告且运行正确，所以不能寄希望于编译器会做严格的类型检查。
因此，为了避免混乱，我们在编写C/C++程序时，对于任何函数都必须一个不漏地指定其类型。如果函数没有返回值，一定要声明为void类型。这既是程序良好可读性的需要，也是编程规范性的要求。另外，加上void类型声明后，也可以发挥代码的"自注释"作用。代码的"自注释"即代码能自己注释自己。

#### 2.1.2 如果函数无参数，则应声明其参数为void

以下代码：

```c
#include <stdio.h>
func(){
        return 1;
}
void main(){
         printf("%dn", func());
         printf("%dn", func(2));
}
```

在C中编译成功（输出都为1），而在C++环境下则会出错.

所以无论在C还是C++中，若函数不接受任何参数，一定要指明参数为void。

#### 2.1.3 void不能代表一个真实的变量

如下面的代码都企图让void代表一个真实的变量，都是错误的：

```
void a;//错误
int function(void b)//错误
```

### 2.2 关于void*

#### 2.2.1 void*类型变量可以接受任何类型指针的赋值，（C语言中甚至是双向的）

因为“无类型”可以包容“有类型”，而“有类型”则不能包容“无类型”。

比如典型的内存操作函数memcpy和memset的函数原型分别为：

```
void *memcpy(void *dest, const void *src, size_t len);
void *memset(void *buffer, int c, size_t num);
```

以下代码执行正确：

```
//示例：memset接受任意类型指针
int intArray[100];
memset(intArray, 0, 100 * sizeof(int));//将intArray清0
//示例：memcpy接受任意类型指针
int intArray1[100], intArray2[100];
memcpy(intArray1, intArray2, 100 * sizeof(int));//将intArray2拷贝给intArray1
```

这样，任何类型的指针都可以传入memcpy和memset中，这也真实地体现了内存操作函数的意义，因为它操作的对象仅仅是一片内存，而不论这片内存是什么类型.

再比如：

```
float * p1;
int* p2;
p1 = p2;
```

其中`p1 = p2`语句会编译出错，必须改为：

```
p1 = (float *)p2;
```
而void*则不同，任何类型的指针都可以直接赋值给它，无需进行强制类型转换：

```
void * p1;
int * p2;
p1 = p2;
```

在C语言中这种转换甚至是双向的(p2 = p1也是正确的) 。无论是在linux还是windows平台下，但是在C++中则不同，void*必须强制类型转换才能赋值给其他类型的指针，所以上述代码在C++中会编译出错。

2.2.2 void*类型不能进行算术操作

按照ANSI标准，不能对void指针进行算术操作，即以下操作都是不合法的：

```
void * pvoid;
pvoid ++;//ANSI：错误
pvoid += 1;//ANSI：错误
```

ANSI标准之所以这样认定，是因为它坚持：进行算术操作的指针必须是确定的知道其指向的数据类型的大小的。例如：

```
int * pint;
pint ++;//ANSI：正确
pint++的结果是使其向后移动sizeof(int)个字节
```

但是大名鼎鼎的GNU则不这么认定，它指定`void*`的算术操作与`char*`一致，如：

```
pvoid ++;//GNU：正确
pvoid += 1;//GNU：正确
pvoid++的结果是使其向后移动了1个字节
```

在实际的程序设计中，为迎合ANSI标准，并提高程序的可移植性，我们可以这样编写实现同样功能的代码：

```
void * pvoid;
(char*)pvoid ++;//ANSI：正确；GNU：正确
(char*)pvoid += 1;//ANSI：错误；GNU：正确
```

GNU和ANSI还有一些区别，总体而言，GNU较ANSI更开放，提供了对更多语法的支持。但是我们在实际应用时，还是应该尽可能地迎合ANSI标准
