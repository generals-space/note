# typedef int(funClass)(void) VS int(*funVar)(void)(待深究)

## 1. 引言

linux内核代码

`/fs/proc/base.c`

```
int proc_pid_readdir(struct file *filp, void *dirent, filldir_t filldir)
{
    ...
}
```

其中，filldir_t，定义在`/linux/linux/fs.h`

```
typedef int (*filldir_t)(void *, const char *, int, loff_t, u64, unsigned);
```

而filldir()是一个函数，在/arch/parisc/hpux/fs.c中有定义

```
static int filldir(void * __buf, const char * name, int namlen, loff_t offset, u64 ino, unsigned d_type)
{
   ...
}
```

## 2. 解析

一个小例子：

```
#include <stdio.h>
typedef int(functionClass1)(void);
typedef int(*functionClass2)(void);
int(*functionVar)(void);
int getInt()
{
        return 100;
}
int main()
{
        functionClass1 *functionA = getInt;
        printf("%dn", functionA());

        functionClass2 functionB = getInt;
        printf("%dn", functionB());

        functionVar = getInt;
        printf("%dn", functionVar());

        functionClass1 *functionArr[2] = {getInt, getInt};
        printf("%dn", functionArr[1]());

}
```

上述代码的运行结果为：

```
100
100
100
100
```

解析：

1.typedef int(funClass)(void)

可以看到形如typedef int(funClass)(void)的表达式将funClass声明为一种新的类似数据类型的东西，地位等同于int，char等，可以用funClass去定义其他变量。

funClass代表的是返回值类型为int，且参数为void(当然可以自行定义其他参数列表)的函数类型(就像function()，注意function后面有括号和无括号的区别)。

然而若要对funClass的变量赋值，需要用返回值同为int类型的函数名，函数名是一个指针，就需要将变量声明为funClass *类型了(如functionA)，或者直接将funClass定义为函数指针类型(如functionB)。

而functionArr数组则是这种定义方式更普便一些的应用。

2.int(*funVar)(void)

根据上述代码，int(*funVar)(void)是定义了一个名为funVar，返回值为int，参数为void的一个函数指针变量，已经被指定了类型，可以直接用getInt函数名赋值。

如果这样声明：int(funVar)(void)...暂时还想不到如何赋值，memcpy(&funVar, getInt, sizeof(getInt))应该是不成的，所以还是不要自找麻烦。

结尾：...虽然明白了typedef int(funClass)(void)和int(*funVar)(void)，不过对于引言中的filldir好像还是不理解，因为从头到尾都没有对filldir进行赋值的语句，再留一个疑惑。