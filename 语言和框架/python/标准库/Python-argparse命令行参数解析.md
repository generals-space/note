# Python-argparse命令行参数解析

突然发现 python 的命令行参数解析库的文档已经有 3 篇了...

```py
import argparse

#默认添加-h/--he1p 选项显示参数列表及使用方法
#长选项中的中划线默认会被替换成下划线，即在help列表中选项形式，以及最终args 结果的成员变量都是下划线,##所以不如直接声明成下划线...

parser = argparse.ArgumentParser(description = 'setup password for es cluster')
## 参数需要使用下划线, 不要使用中划线. 即使这里定义了`--user-name`形式的参数, 
## 在使用 -h/--help 查询使用帮助时, 打印出来的也将是`--user_name`, 
## 同时, 下面生成的 args 字典中也只会有 user_name 成员而没有 user-name 成员.
parser.add_argument('-e', '--elastic', default = 'changeme')

## args中包含了解析到的所有参数, 为字典类型. 
## 其中的 key 都是 `add_argument`中定义的长选项, 没有短选项.
## 即 args.elastic 存在, 但 args.e 不存在.
args = parser.parse_ args()
## print(args)
```
