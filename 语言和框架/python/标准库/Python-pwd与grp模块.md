# Python标准库-pwd,grp模块

参考文章

[Python标准库：pwd模块](http://biancheng.dnbcw.net/python/462172.html)

[python pwd和grp模块](http://www.361way.com/python-pwd-grp/4015.html)

## 1. pwd模块

`pwd`模块功能简单，在linux/unix平台下使用，用来获取`/etc/passwd`文件的相关信息。可以用来作为获取用户信息的接口。下面我们来看看具体的使用细节。

pwd模块有3个函数: `getpwuid`, `getpwnam`, `getpwall`。通过这三个函数结合list表的相关属性，可以获得用户的各类信息。

- `getpwuid`: 根据UID获取用户信息, 

- `getpwnam`: 根据USERNAME获取用户信息

- `getpwall`: 获取所有的用户信息

示例

```py
#!/usr/bin/env python
#!coding:utf-8
import pwd

rootInfo = pwd.getpwnam('root')
print(rootInfo)
print(type(rootInfo))

rootInfo = pwd.getpwuid(0)
print(rootInfo)

print('获取成员变量......')
## 使用方法, 通过属性名
print(rootInfo.pw_dir)
## 通过索引值
print(rootInfo[5])

## 获取所有用户的信息
## print(pwd.getpwall())
```

执行它, 输出如下

```
pwd.struct_passwd(pw_name='root', pw_passwd='x', pw_uid=0, pw_gid=0, pw_gecos='root', pw_dir='/root', pw_shell='/bin/bash')
<type 'pwd.struct_passwd'>
pwd.struct_passwd(pw_name='root', pw_passwd='x', pw_uid=0, pw_gid=0, pw_gecos='root', pw_dir='/root', pw_shell='/bin/bash')
获取成员变量......
/root
/root
```

## 2. grp模块

grp模块，提供了一个Unix 用户组/group（/etc/group）数据库的接口

- `getgrgid(gid)`: 返回对应gid的组信息 
- `getgrname(name)`: 返回对应group name的组信息 
- `getgrall()`: 返回所有组信息 

使用方法与pwd模块类似.
