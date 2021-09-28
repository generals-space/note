# Python格式化读入[scanf re 正则]

参考文章

1. [Python Read Formatted String](https://stackoverflow.com/questions/7111690/python-read-formatted-string)
    - `struct.unpack()`
2. [Python有类似于C语言的格式化输入scanf吗？](https://www.zhihu.com/question/49840816)
    - python 让用户尽量不关注数据的类型，而`scanf()`对 type 的要求有失大道至简原则
3. [python输入，格式化输入，以及scanf的替代方案](https://blog.csdn.net/u010138758/article/details/70163944/)
    - 从字符串中提取格式化信息的正则表达式示例
    - 调用c标准库
4. [Python使用struct处理二进制(pack和unpack用法)](https://blog.csdn.net/jackyzhousales/article/details/78030847)

## 引言

有如下字符串

```
Physical size: 1080x2340
```

我想从中提取出宽(1080)及高(2340)的值, 想了想, 没找到合适的方法.

在C语言中, 可以使用类似`scanf()`的函数完成, 但是在 python 中, 貌似只能用`split()`等自行切割字符串, 感觉很不优雅, 于是想找找看有没有可用的方式.

## re 正则

参考文章2中提到, python 让用户尽量不关注数据的类型，而`scanf()`对 type 的要求有失大道至简原则, 因此基本上可以确定没有什么官方的办法来实现我的目的, 只能用`re`正则完成.

不过用正则的话, 首先是各数据类型的对应关系要写的话比较麻烦, 另外就是结果的获取, 参考文章3给出了示例.

| `scanf()`类型          | 正则表达式类型                            | 含义                                   |
| :--------------------- | :---------------------------------------- | :------------------------------------- |
| `%c`                   | `.`                                       | 单个任意字符                           |
| `%d`                   | `[-+]?\d+`                                | 整型数值(包括负值)                     |
| `%u`                   | `\d+`                                     | 整型数值(无符号, 即正整数)             |
| `%s`                   | `\S+`                                     | 字符串(不含空白字符)                   |
| `%e`, `%E`, `%f`, `%g` | `[-+]?(\d+(\.\d*)?|\.\d+)([eE][-+]?\d+)?` | 浮点数(包括负值)                       |
| `%o`                   | `[-+]?[0-7]+`                             | 八进制(包括负值)                       |
| `%x`, `%X`             | `[-+]?(0[xX])?[\dA-Fa-f]+`                | 十六进制(包括负值)                     |
| `%i`                   | `[-+]?(0[xX][\dA-Fa-f]+|0[0-7]*|\d+)`     | 十进制, 八进制, 十六进制整数(包括负值) |
| `\`                    | `.{5}`                                    | 这个在`scanf()`中从来没遇到过...       |

## 示例

给定一个字符串

```
/usr/sbin/sendmail - 0 errors, 4 warnings
```

假设使用如下格式的`scanf()`函数捕获

```c
scanf("%s - %d errors, %d warnings", &cmd, &errorNum, &warnNum);
```

那么对应的正则方法为

```py
import re
pattern = re.compile(r"(\S+) - (\d+) errors, (\d+) warnings")
match = pattern.match('/usr/sbin/sendmail - 0 errors, 4 warnings')
## 打印结果为: ('/usr/sbin/sendmail', '0', '4'), 需要对元组中的内容分别捕获.
if match: print(match.groups())
```

大致理解了吧?

## 关于 struct.unpack()

对这个函数, 捕获字符串中信息时必须精确到每个字符, 太麻烦了, 试验没做完, 先放着吧...
