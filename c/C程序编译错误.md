# C语言/C++编译错误

## 1. C++

```
error: 'NULL' was not declared in this scope
```
NULL并不是关键字，在C++中貌似不能直接使用，但它被定义在一些标准头文件中，C++可以使用如下头文件：

```c
#include <cstddef>
```

这样还包括了一些其他基本类型如：`std::size_t`

## 2.

```
error: expected constructor, destructor, or type conversion before '.' token.
```

光从这个错误提示是看不出什么问题的,所以不知道原因会感觉很奇怪. C++中, 全局域(在main函数之外, 与`#include`语句平级的地方)只能声明,初始化变量,不能对变量进行赋值,运算,调用函数等操作,谨记.

## 3.

```
relocation 0 has invalid symbol index
```

问题在于链接的时候找不到.o文件, 是不是不在同一目录又没显式指定路径?

## 4.

```
declare class does not name a type
```

出现这个编译错误主要有四个可能原因，简单来说就是找不到目标类对象, 现总结如下：

1. 引用的类命名空间未包含

2. 引用的类头文件未包含

3. 包含了头文件，或者已经前向声明了，则说明所引用的类名写错。

4. 循环引用头文件

前置声明要素：

1. 前置声明需要注意以上提到的四点

2. 尽可能的采用前置声明（做到只有包含继承类的头文件）

3. 使用前置声明时，cpp文件中include 头文件次序必须先 包含前置声明的类定义头文件，再包含本类头文件。

否则会出现如下编译错误.

```
(expected constructor, destructor, or type conversion before ‘typedef')
```

## 5.

```
field has incomplete type
```

类或结构体的前向声明只能用来定义指针对象或引用，因为编译到这里时还没有发现定义，不知道该类或者结构的内部成员，没有办法具体的构造一个对象，所以会报错。

将类成员改成指针或引用就好了。 程序中使用incomplete type实现前向声明有助与实现数据的隐藏，要求调用对象的程序段只能使用声明对象的引用或者指针。

在显式声明异常规范的时候不能使用incomplete type。

## 6.

```
a label can only be part of a statement and a declaration is not a statement
```

错误原因: 可能是在switch下的case语句中定义了变量并为其赋值;

解决办法: 将声明语句放在switch块外即可, 或者给case语句加上大括号, 不过这个没试过.

## 7. 

```c
#include <stdio.h> 
#include <string.h> 
int main(void) 
{ 
    char *buf; 　　　　　　　　　　　　　　//定义char指针
    char *string = "hello "; 　　　　　　//指向常量数据区的“hello”字符串
    buf = string; 　　　　　　　　　　　　　 //将指向常量的指针赋值

    printf("buf=%sn ", buf); 
    strcpy(buf, "world"); 　　　　　　　　　//试图将常量数据区的 "world"字符串拷贝给指向常量数据区的buf
    printf("buf=%sn", buf);
    return 0;
}
```

于是当运行时，在strcpy处，出现了段错误 `Segmentation fault`
解决办法：

1. buf没有空间, 应该用malloc分配空间

```
buf = malloc(4);
```

2. 改变 string为：

```
char string[ ] = "hello";
```

这样，string是指向数组的指针，赋值后，buf也是指向数组的指针，再次调用strcpy时，就把“world”复制到数组中了！

3. 可以直接赋值：

```
buf = "world";
```

> 小结：指针只存贮了一个地址，想把整个字符串复制给他，必须手动分配内存空间，或存放于数组之中。(注意名词：**常量数据区**)