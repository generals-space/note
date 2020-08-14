# Python传值调用与传址调用

python中没有像C++那样有显式的引用/指针传递的运算符, 一般来说, python的传值和传址是根据传入参数的类型来选择的

传值的参数类型：数字，字符串，元组（immutable）

传址的参数类型：列表，字典（mutable）

如果函数传入的参数是字典或者列表，就能修改对象的原始值——相当于传址。如果函数传入的是数字、字符或者元组类型，就不能直接修改原始对象——相当于传值。


```py
#!/usr/bin/python

def called(li):
    '''
    调用函数
    '''
    li[0] = 'a'
    print('%s in called' % li)

def caller():
    '''
    主调函数
    '''
    li = [1,2,3]
    print('%s before caller' % li)
    called(li)
    print('%s after caller' % li)

if __name__ == '__main__':
    caller()
```

上述代码的执行结果如下

```
[1, 2, 3] before caller
['a', 2, 3] in called
['a', 2, 3] after caller
```

但是如果函数传入的是字符串, 数字, 元组等类型, 调用函数内部对它们的修改是无法对主调函数的变量产生影响的.


------

下面测试一下传入自定义类型对象的参数属性.

```py
#!/usr/bin/python

class TestObj():
    name = 'a'
    def getName(self):
        return self.name
    def setName(self, name):
        self.name = name
def called(obj):
    obj.name = 'b'
    print('%s in called' % obj.name)

def caller():
    obj = TestObj()
    print('name %s before caller' % obj.name)
    called(obj)
    print('name %s after caller' % obj.name)

if __name__ == '__main__':
    caller()
```

```
name a before caller
b in called
name b after caller
```

看来也是传址调用.

这又涉及到变量复制, 与深浅拷贝的问题.