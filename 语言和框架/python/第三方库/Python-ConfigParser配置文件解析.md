# Python-ConfigParser配置文件解析

参考文章

1. [Python 解析配置模块之ConfigParser详解](http://www.pythontab.com/html/2014/pythonhexinbiancheng_1120/919.html)

ConfigParser模块提供了3个类: `RawConfigParser`、`ConfigParser`、`SafeConfigParser`. `RawCnfigParser`是最基础的ini文件读取类, ConfigParser、SafeConfigParser支持对`%(value)s`变量的解析.

示例配置文件如下:

`test.ini`

```ini
[sec_a] 
a_key1 = 20 
a_key2 = 10 
   
[sec_b] 
b_key1 = 121 
b_key2 = b_value2 
b_key3 = $r 
b_key4 = 127.0.0.1
```

示例代码

`parse_test.py`

```py
#!/usr/bin/env python
#!coding:utf-8
import ConfigParser 

configParser = ConfigParser.ConfigParser() 
## 实例化后首先读取配置文件
configParser.read("test.ini") 
## 得到所有块
secs = configParser.sections() 
print 'sections:', secs 
## 得到指定块的所有字段
opts = configParser.options("sec_a") 
print 'options:', opts 
## 得到指定块的所有值
kvs = configParser.items("sec_a") 
print 'sec_a:', kvs 
   
## 得到指定字段的值, 可以指定返回类型
str_val = configParser.get("sec_a", "a_key1") 
int_val = configParser.getint("sec_a", "a_key2") 

print "value for sec_a's a_key1:", str_val 
print "value for sec_a's a_key2:", int_val 

## 更新指定块指定字段的值
configParser.set("sec_b", "b_key3", "new-$r") 
configParser.set("sec_b", "b_newkey", "new-value") 

## 创建新块, 并为其创建新的字段
configParser.add_section('a_new_section') 
configParser.set('a_new_section', 'new_key', 'new_value') 
   
## 实际的写回操作
configParser.write(open("test.ini", "w"))
```

执行它, 终端输出如下

```
sections: ['sec_a', 'sec_b']
options: ['a_key1', 'a_key2']
sec_a: [('a_key1', '20'), ('a_key2', '10')]
value for sec_a's a_key1: 20
value for sec_a's a_key2: 10
```

更新后的`test.ini`如下

```ini
[sec_a]
a_key1 = 20
a_key2 = 10

[sec_b]
b_key1 = 121
b_key2 = b_value2
b_key3 = new-$r
b_key4 = 127.0.0.1
b_newkey = new-value

[a_new_section]
new_key = new_value
```

关于ConfigParser与SafetyConfigParser之于RawConfigParser的区别, 在于前两者对指定字段的写入操作可以使用变量更新目标值的指定部分, 而后者只能单纯地全部覆写.

如下示例配置文件

`test2.ini`

```ini
[portal] 
url = http://%(host)s:%(port)s/Portal 
host = localhost 
port = 8080
```

使用`RawConfigParser`: 

```py
#!/usr/bin/env python
#!coding:utf-8
import ConfigParser 
  
configParser = ConfigParser.RawConfigParser() 
  
print "use RawConfigParser() read" 
configParser.read("test2.ini") 
print configParser.get("portal", "url") 
  
print "use RawConfigParser() write" 
## 这一行没有其实实际写入
configParser.set("portal", "url2", "%(host)s:%(port)s") 
print configParser.get("portal", "url2")
```

执行它, 得到的输出为

```
use RawConfigParser() read
http://%(host)s:%(port)s/Portal
use RawConfigParser() write
%(host)s:%(port)s
```

如果使用ConfigParser类

```py
#!/usr/bin/env python
#!coding:utf-8
import ConfigParser 
## 只有这一行不同
configParser = ConfigParser.ConfigParser() 
  
print "use ConfigParser() read" 
configParser.read("test2.ini") 
print configParser.get("portal", "url") 
  
print "use ConfigParser() write" 
configParser.set("portal", "url2", "%(host)s:%(port)s") 
print configParser.get("portal", "url2")
```

执行它, 得到的终端输出为

```
use ConfigParser() read
http://localhost:8080/Portal
use ConfigParser() write
localhost:8080
```

可以看到, ConfigParser类在获取`url`字段时, 将其中的`$(key)s`部分自动使用同属`[protal]`块字段`key`的值进行替换, 十分方便. 

不过, 类似这种`url`字段中涉及到的变量必须存在, 并且同属同一个`[]`块内, 否则会报错.

> 注: SafeConfigParser与ConfigParser效果相同.

## include文件包含

很多时候我们需要在主配置文件中引用多个子配置文件, 便于管理. 一种简单的解决方法是, 组合使用`readfp`与`read`方法. 

它们的区别在于, `readfp()`接受的参数是使用`open()`函数打开的文件对象, 而`read()`接受的参数是表示文件路径的字符串.

以上述两个`ini`文件为例, 使用如下代码同时读取这两个文件

```py
#!/usr/bin/env python
#!coding:utf-8
import ConfigParser 

configParser = ConfigParser.ConfigParser()
## readfp方法不接受以`w`模式打开的文件...
fp = open('test.ini', 'r')
configParser.readfp(fp)

configParser.read('test2.ini')

## 打印test.ini中的字段
print(configParser.get("sec_a", "a_key1"))
## 打印test2.ini中的字段
print(configParser.get("portal", "url"))

## 更新的时候就有乌龙了, 没有执行`write()`方法时看不出来 
configParser.add_section('a_new_section') 
configParser.set('a_new_section', 'new_key', 'new_value') 
configParser.set('portal', 'url2', 'abc')

print(configParser.get("sec_a", "a_key1"))
print(configParser.get("portal", "url2"))

## 目前为止更新操作还是蛮正常的, 但是write()后就会发现这两个文件内容相同了, 而且是是合并后的结果...
## 这样的话, 多文件读写实际应用中并不可取
configParser.write(open("test.ini", "w"))
configParser.write(open("test2.ini", "w"))
```

其实在实际应用中, 我们一般是需要在使用`readfp`并得到`[include]`块后再根据其下引用的文件继续进行解析的. 这里简化了这个过程, 但道理是一样的.