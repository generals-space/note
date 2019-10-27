# powershell格式化输出

参考文章 

1. [Powershell使用-F方法带入数据](https://www.pstips.net/using-f-operator-to-combine-string-and-data.html)
2. [格式化输出数字](https://www.pstips.net/formatting-numbers-easily.html)

一个简单示例

```console
$ 'hello {0}, this is {1}' -f 'general','jiangming'
hello general, this is jiangming
```

可以看出, `{}`中可以表示参数索引.

> 目前还没有找到可以引用对象的某个属性的方式(python中有`'{key}'.format(obj)`这种方式的)

貌似powershell中的格式化输出比较弱啊, 参考文章1中只给出了3种方式

1. `d`: 控制前置0补全
2. `n`: 定义小数位精确度
3. `p`: 格式化为百分数(将小于1的小数转换成百分数的形式, 但结果只能是整数, 不会包含小数点, 尾巴将会四舍五入)

```console
$ $number = 68
$ '{0:d7}' -f $number
0000068
```

```console
$ $number = 35553568.67826738
$ '{0:n1}' -f $number
35,553,568.7
```

```console
$ $number = 0.32562176536
$ '{0:p0}' -f $number
33%
```
