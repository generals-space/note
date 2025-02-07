# Python-指定参数类型与返回值类型

参考文章

1. [Python 强类型编程](https://blog.dreamrounder.com/posts/python/strong-type-coding/)
    - dataclass
2. [用@dataclasses和@dataclasses_json做嵌套类型的序列化和反序列化，并定义属性的对外映射字段](https://blog.csdn.net/yournevermore/article/details/139474398)
3. [Dataclass object property alias](https://stackoverflow.com/questions/67001442/dataclass-object-property-alias)
4. [Add alias as a field() parameter for dataclasses](https://github.com/python/cpython/issues/101192)
5. [Python class instances share the same values](https://stackoverflow.com/questions/67102797/python-class-instances-share-the-same-values)
    - dataclass实例的成员属性在所有实例中发生共享, 变成了类属性
6. [In python how to create multiple dataclasses instances with different objects instance in the fields?](https://stackoverflow.com/questions/75794069/in-python-how-to-create-multiple-dataclasses-instances-with-different-objects-in)
    - 高赞回复描述得比较清楚
7. [Why is dataclass field shared across instances](https://stackoverflow.com/questions/73598938/why-is-dataclass-field-shared-across-instances)

## dataclass

参考文章1中介绍了`dataclass`注解, 对于固定结构的数据, 要比dict类型用起来更方便.

```py
from dataclasses import dataclass

## dataclass 会自动的生成构造函数和默认值, 更贴近其他强类型语言的使用方式.
@dataclass
class NodeMetrics:
    cpu:float
    mem:float

@dataclass
class NodeInfo():
    name:str = None
    timestamp:float = None
    metrics:NodeMetrics = None
```

但是序列化/反序列化很不灵活, 作为私有成员的字段不能出现特殊字符, 比如`name`成员, 在转化成json时希望显示为`name.1`(类似于golang struct中的json tag).

目前 python 原生还不支持, 参考文章9官方issue中还在讨论, 需要引入第三方库, 见参考文章7和8.

## dataclass实例的成员属性都变成了类属性, 所有实例共享?

参考文章5, 6, 7都描述了一个问题(我也遇到了), 最简示例如下

```py
from dataclasses import dataclass

@dataclass
class MyData2():
    b:int = None

@dataclass
class MyData1():
    ## 设置默认值
    a:MyData2 = MyData2() 

my_data11 = MyData1()
my_data12 = MyData1()
my_data11.a.b = 123
print(my_data12.a.b) ## 123
```

由于`MyData1.a`在设置默认值时使用了`MyData2()`这是一个实例, 因此所有的 MyData1 实例的成员a都是同一个 MyData2() 实例对象, 这是不对的, 需要使用`default_factory`, 如下

```py
    a:MyData2 = field(default_factory=MyData2)
```
