---
title: Python标准库-getopt类
---

参考文章

[Python 命令行参数和getopt模块详解](http://www.tuicool.com/articles/jaqQvq)

`getopt()`的语法为

```python
opts, extra_args = getopt(ori_args, shortopts, longopts = [])
```

其中`ori_args`为待解析的参数列表, 注意是**列表类型**.

`shortopts`必选, 字符串类型, 形式为`ab:c:`, 其后紧跟一个冒号的(如`b`与`c`), 表示其后需要一个参数.

`longopts`可选, 列表类型, 形式为['a_long', 'b_long=', 'c_long='], 同样, `b_long`与`c_long`这样的后面接一个`=`的长选项, 表示之后需要一个参数.

返回值有两个, opts为(选项, [参数])为成员组成的列表, 如果传入的参数列表包含未在`shortopts`与`longopts`出现的选项, 则会将它们放在`extra_args`, 用户可以在代码中自行处理它们.

示例

```python
#!/usr/bin/env python
#!coding:utf-8

import getopt

arg_list = '-i -n logic -v 0.3.12'
arg_list = arg_list.split()
try:
    opts, extra_args = getopt.getopt(arg_list, 'in:v:', ['install', 'name=', 'version='])
except getopt.error, e:
    print(e) 
print(opts)
print(extra_args)

for name, value in opts:
    if name == '-i':
        print value
    if name == '-n':
        print value
    if name == '-v':
        print value
```

执行它, 输出如下

```
[('-i', ''), ('-n', 'logic'), ('-v', '0.3.12')]
[]

logic
0.3.12
```