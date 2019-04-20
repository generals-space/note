# Python内置变量

## 2. __doc__

```py
#!/usr/bin/env python
#!coding:utf-8

'''
模块文档
'''

import os
import sys 

class MyClass:
    ''' 
    类文档
    '''
    def say(self):
        print self.__doc__
def main():
    ''' 
    函数文档
    '''
    ## 打印模块文档
    print __doc__
    ## 打印函数文档
    print main.__doc__
    ## 打印类文档
    print MyClass.__doc__

    ## 类对象打印本身文档
    myClass = MyClass()
    myClass.say()
main()
```

执行它, 将得到如下

```
模块文档
    函数文档
    类文档
    类文档
```

## 3. `__name__`变量

python中, 每个py文件都是一个模块, 也同时是一个可执行文件(即包含main方法). 因此, 对每个py文件, 可以单独运行, 也可以`import`它给其他程序使用, 这两种情况不一样. 为了区分这两种情况, 可以使用`__name__`属性.

当py文件是直接运行时, `__name__ = "__main__"`,

当此文件被当作模块导入时, `__name__ = 其本身模块名`
