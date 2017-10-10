# Makefile编写(一)

参考文章

1. [阮一峰的网络日志 - Make 命令教程](http://www.ruanyifeng.com/blog/2015/02/make.html)

> 代码变成可执行文件, 叫做**编译**(compile)；先编译这个, 还是先编译那个(即编译的安排), 叫做**构建**(build). 

## 1. Makefile基本结构

### 1. 基本格式

```makefile
target: dependency_files
    command
```

`target`是需要由make工具创建的目标体, 通常是目标文件, 可执行文件, 或是一个标签;

`dependency_files`是要创建目标体所依赖的文件;

创建每个目标体时需要执行的命令. (command前必须要有"Tab"符, 否则会在执行make时出错). 

## 2. 简单示例

两个文件, 分别为"hello.h", "hello.c", 希望创建目标体"hello.o", 需要执行的命令为

```
gcc -c hello.c -o hello.o
```

相应的Makefile可以写为

```makefile
hello.o: hello.c hello.h
    gcc -c hello.c -o hello.c
```

然后就可以使用make了, 使用make的格式为"make target", 如上述代码中, 应执行

```
make hello.o
```

会在当前目录下产生hello.o文件.

## 3. 复杂示例

以下工程包含3个头文件和8个C文件, 其Makefile如下

```makefile
edit: main.o kbd.o command.o display.o \
       insert.o search.o files.o utils.o
    cc -o edit main.o kbd.o command.o display.o \
               insert.o search.o files.o utils.o
main.o: main.c defs.h
    cc -c main.c
kbd.o: kbd.c defs.h command.h
    cc -c kbd.c
command.o: command.c defs.h command.h
    cc -c command.c
display.o: display.c defs.h buffer.h
    cc -c display.c
insert.o: insert.c defs.h buffer.h
    cc -c insert.c
search.o: search.c defs.h buffer.h
    cc -c search.c
files.o: files.c defs.h buffer.h command.h
    cc -c files.c
utils.o: utils.c defs.h
    cc -c utils.c
clean:
    rm edit main.o kbd.o command.o display.o \
            insert.o search.o files.o utils.o
```

解析：
------

1. 反斜线'\'代表换行符;

2. `clean`是一个标签, 它不依赖任何文件, 所以冒号后面没有东西.

3. 如果Make命令运行时没有指定目标, 默认会执行Makefile文件的第一个目标.

## 2. Makefile变量

### 1. 用户自定义变量

上一节的复杂示例中, edit规则中依赖文件出现了两次, 还在clean标签里面出现了一次, 可以使用以下的方式定义变量, 用以简化

```makefile
objects = main.o kbd.o command.o display.o 
             insert.o search.o files.o utils.o
edit : $(objects)
           cc -o edit $(objects)
/*其他部分不变*/
clean :
           rm edit $(objects)
```

2.预定义变量

|命令格式	   |    含义 |
|:-:|:-:|
|CC	              |    C编译器的名称, 默认值为cc |
|CPP	         |    C预编译器的名称, 默认值为$(CC) -E |
|CXX	         |    C++编译器的名称, 默认值为g++ |
|CFLAGS	      |    C编译器的选项, 无默认值 |
|CPPFLAGS	|    C预编译器的选项, 无默认值 |
|CXXFLAGS	|    C++编译器的选项, 无默认值 |
|RM	              |   文件删除程序的名称, 默认值为rm -f |
|ASFLAGS	  |   汇编程序的选项, 无默认值 |

## 3. 自动变量

|命令格式	      |含义 |
|:-:|:-:|
|$*	          |         不包含扩展名的目标文件名称|
|$@	          |         目标文件的完整名称|
|$%	          |         如果目标是归档成员, 则该变量表示目标的归档成员名称|
|$+	          |         所有的依赖文件, 以空格分开, 并以出现的先后为序, 可能包含重复的依赖文件|
|$^	          |         所有不重复的依赖文件, 以空格分开|
|$<	          |         第一个依赖文件的名称|
|$?	          |         所有时间戳比目标文件晚的依赖文件, 并以空格分开|

则上述复杂示例的代码可改写为

```makefile
object = main.o kbd.o command.o display.o 
       insert.o search.o files.o utils.o
CC = gcc
CFLAGS = -Wall -O -g

edit : $(object)
    $(CC) $^ -o $@
main.o : main.c defs.h
    $(CC) $(CFLAGS) -c $< $@
kbd.o : kbd.c defs.h command.h
    $(CC) $(CFLAGS) -c $< $@
command.o : command.c defs.h command.h
    $(CC) $(CFLAGS) -c $< $@
display.o : display.c defs.h buffer.h
    $(CC) $(CFLAGS) -c $< $@
insert.o : insert.c defs.h buffer.h
    $(CC) $(CFLAGS) -c $< $@
search.o : search.c defs.h buffer.h
    $(CC) $(CFLAGS) -c $< $@
files.o : files.c defs.h buffer.h command.h
    $(CC) $(CFLAGS) -c $< $@
utils.o : utils.c defs.h
    $(CC) $(CFLAGS) -c $< $@
clean :
    rm edit $(object)
```

扩展阅读

http://blog.csdn.net/ruglcc/article/details/7814546/

## 4. 变量定义

`var := $(a)`, 如果变量`a`在之前和之后都有定义, 则`var`只取在其前面定义的变量值.

```makefile
a = 123
var := $(a)
a = abc
```

则`var`的值为123