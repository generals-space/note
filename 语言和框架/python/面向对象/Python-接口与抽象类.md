# Python-接口与抽象类

参考文章

1. [python中的接口和依赖注入](http://www.cnblogs.com/xinsiwei18/p/5937952.html)

## 1. 接口

python里无接口类型，定义接口只是一个人为规定，在编程过程 **自我约束**, 实现接口的过程与普通继承没什么区别. 在其他的语言里，比如Java，继承类没有重写接口方法是会报错的，而在python里不会，就是因为python没这个类型，所以只是在我们编程过程的一个规定，以`I`开头的类视为接口. 

```py
#定义接口Interface类来模仿接口的概念，但python中压根就没有interface关键字来定义一个接口。
class Interface:
    def read(self): #定接口函数read
        pass
    def write(self): #定义接口函数write
        pass

class Txt(Interface): #文本，具体实现read和write
    def read(self):
        print('文本数据的读取方法')
    def write(self):
        print('文本数据的读取方法')

class Sata(Interface): #磁盘，具体实现read和write
    def read(self):
        print('硬盘数据的读取方法')
    def write(self):
        print('硬盘数据的读取方法')
```

## 2. 抽象类

python中实现抽象类要借助标准库`abc`.

```py
#!encoding:utf-8
import abc

## class FileType(metaclass = abc.ABCMeta):
class FileType():
    ## 指定FileType的元类为ABCMeta
    __metaclass__ = abc.ABCMeta
    ## 抽象类除了约束子类必须实现抽象方法外, 还约束了子类的成员属性.    
    generalType = 'file'

    #定义抽象方法，无需实现功能
    @abc.abstractmethod 
    def read(self):
        '子类必须定义读功能'
        pass
    @abc.abstractmethod
    def write(self):
        '子类必须定义写功能'
        pass

class Memory(FileType):
    generalType = "memory"    
    def read(self):
        print('内存数据的读取方法')
    def write(self):
        print('内存数据的写入方法')

class Disk(FileType): 
    def read(self):
        print('硬盘数据的读取方法')

    def write(self):
        print('硬盘数据的写入方法')

memory = Memory()
disk = Disk()

#这样大家都是被归一化了,也就是一切皆文件的思想
memory.read()
disk.write()

print(memory.generalType) 
print(disk.generalType) 
```

上述代码的执行结果为

```
内存数据的读取方法
硬盘数据的写入方法
memory
file
```

通过对上述代码进行些许修改以验证**抽象类**的特性.

1. 继承抽象类的子类必须实现该抽象类定义的所有抽象方法. (删掉`Memory`或`Disk`的`read`或`write`函数, 执行会报错.)

2. 抽象类中可以定义非抽象的函数(只要不以`abstractmethod`修饰即可), 子类不需要实现非抽象函数.