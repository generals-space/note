# Python参数列表args与kwargs

原文链接

[Python函数可变参数args及kwargs释义](http://lovesoo.org/python-han-shu-ke-bian-can-shu-args-ji-kwargs-shi-yi.html)

初学Python看到代码中类似`func(*args, **kwargs)`这样的定义时，经常感到一头雾水。

下面通过一个简单的例子来解释Python函数可变参数args及kwargs的意思：

`*args`表示任何多个无名参数，它是一个`tuple`

`**kwargs`表示关键字参数，它是一个`dict`

同时使用`*args`和`**kwargs`时，`*args`参数列必须要在`**kwargs`前，要是像`foo(1,a=1,b=2,c=3,2,3)`这样调用的话，则会提示语法错误`SyntaxError: non-keyword arg after keyword arg`。

测试代码如下：

```python
def foo(*args,**kwargs):
    print 'args=',args
    print 'kwargs=',kwargs
    print '**********************'
 
if __name__=='__main__':
    foo(1,2,3)
    foo(a=1,b=2,c=3)
    foo(1,2,3,a=1,b=2,c=3)
    foo(1,'b','c',a=1,b='b',c='c')
```

执行结果如下：

```
args= (1, 2, 3)
kwargs= {}
**********************
args= ()
kwargs= {'a': 1, 'c': 3, 'b': 2}
**********************
args= (1, 2, 3)
kwargs= {'a': 1, 'c': 3, 'b': 2}
**********************
args= (1, 'b', 'c')
kwargs= {'a': 1, 'c': 'c', 'b': 'b'}
**********************
````