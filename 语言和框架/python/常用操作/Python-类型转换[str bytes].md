# Python-类型转换[str bytes]

参考文章

1. [python3 字符串str和bytes相互转换](https://www.jb51.net/article/241902.htm)

## 1. 相关基础

python3中有两种字符串类型：`str`和`bytes`

python编码问题可以参考[文章](https://www.jb51.net/article/241895.htm)

`str`以`unicode`编码格式保存在内存

所以使用时，不用管前面要不要加`u`

(python2中需要考虑，不加`u`的话，在一些场合会报错)

```py
#!/usr/bin/python3
str0="i am fine thank you"

print(type(str0))
print(str0)

str0=u"i am fine thank you"

print(type(str0))
print(str0)
```

```
# <class 'str'>
# i am fine thank you
# <class 'str'>
# i am fine thank you
```

定义`byte`类型时，在字符串前加`b`

```py
#!/usr/bin/python3
str0=b"i am fine thank you"

print(type(str0))
print(str0)
```

```
# <class 'bytes'>
# b'i am fine thank you'
```

## 2. str和bytes相互转换

在文件传输过程中，通常使用`bytes`格式的数据流，而代码中通常用`str`类型，因此`str`和`bytes`的相互转换就尤为重要。

### 2.1 bytes->str

```py
#!/usr/bin/python3

bytes_data = b'this is a message'
print(type(bytes_data))
print(bytes_data)

# 方法一：
str_data = str(bytes_data, encoding='utf-8')
print(type(str_data))
print(str_data)

# 方法二：
str_data = bytes_data.decode('utf-8')
print(type(str_data))
print(str_data)
```

```
# <class 'bytes'>
# b'this is a message'
# <class 'str'>
# this is a message
# <class 'str'>
# this is a message
```

### 2.2 str->bytes

```py
#!/usr/bin/python3

str_data = 'this is a message'
print(type(str_data))
print(str_data)
# 方法一：
bytes_data = bytes(str_data, encoding='utf-8')
print(type(bytes_data))
print(bytes_data)
# 方法二：
bytes_data = str_data.encode('utf-8')
print(type(bytes_data))
print(bytes_data)
```

```
# <class 'str'>
# this is a message
# <class 'bytes'>
# b'this is a message'
# <class 'bytes'>
# b'this is a message'
```
