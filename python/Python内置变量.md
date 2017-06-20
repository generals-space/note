# Python内置变量

## 2. __doc__

```python
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