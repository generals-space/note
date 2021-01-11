# Python-指定参数类型与返回值类型

参考文章

1. [python限定方法参数类型、返回值类型、变量类型等](https://www.cnblogs.com/linkenpark/p/11676297.html)
2. [How can I specify the function type in my type hints?](https://stackoverflow.com/questions/37835179/how-can-i-specify-the-function-type-in-my-type-hints)
    - `typing.Callable`函数类型

python 是弱类型语言, 每个变量的类型是不固定的, 如果要对一个变量做某种类型独有的操作时, 可能需要先对此变量进行类型检查.

不过都这么多年过去了, 大家在写代码的时候都会按照函数约定每个变量的类型, 最多只要判断一下是不是`None`就行了, 就是在代码提示的时候可能一下子没法看出一个变量到底是什么类型.

python 3.5 开始, 引入了类型注解(type hints), 可以在定义函数时写明参数的类型.

但这种语法只是一种约定, 相当于注解, 如果传入的参数不符合也不会影响程序的运行.

```py
def test(a:int, b:str) -> str:
    print(a, b)
    return 1000

if __name__ == '__main__':
    test('test', 'abc')
```

虽然`test()`第一个参数指定了`int`, 但是传入一个字符串也不会报错.
