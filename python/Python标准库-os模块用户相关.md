# Python标准库-os模块 用户相关

## 1. 用户相关


os.getuid(): 获取执行当前程序的用户uid

os.setuid(uid): 设置当前程序的启动用户为uid所表示的用户, uid只能是整型数值


示例:

```python
#!/usr/bin/env python
#!coding:utf-8
import os
current_uid = os.getuid()
print('当前启动用户uid为: %s' % current_uid)

```

## 2. 组相关

os.setgroups(grouplist): 设置进程的访问组列表, grouplist必须是一个列表类型变量

os.setgid(gid): 设置当前程序的启动用户组, gid为整型数值