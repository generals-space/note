# Shell脚本技巧

## 1. 脚本内部实现切换用户并执行命令操作

参考文章

[Shell脚本中实现切换用户并执行命令操作](http://www.jb51.net/article/59255.htm)

### 1.1 情景重现

首先来看一下, 不使用任何特殊方法时, 在shell脚本中进行用户切换并执行命令的情况.

```shell
#!/bin/bash

echo 'i am in' $(pwd) 'at the beginning'
su - general
echo 'i am in' $(pwd) 'now'
exit
echo 'i am in' $(pwd) 'finally'
```

命名为`test1.sh`, 使用root执行.

```
[root@b14e517d408b tmp]# ./test1.sh 
i am in /tmp at the beginning
[general@b14e517d408b ~]$ exit
logout
i am in /tmp now
[root@b14e517d408b tmp]# 
```

执行`test1.sh`, 到`su - general`这一句后会陷入新的shell中, 并脱离原来执行脚本时的shell, 脚本进程被挂起. 在这个shell中执行任何命令都不会影响原来的脚本执行...

**手动执行`exit`**后回到原来的shell, 脚本继续执行. 但是注意, 这里输出的是'i am in /tmp now', 说明`exit`命令结束的是`su - general`这条命令, 之后继续执行的代码已经回到原来的shell, 我们使用`su`命令切换用户身份执行命令的目的并没有实现.

脚本中的`exit`直接使脚本退出了, 所以不会看到'i am in ... finally'.

------

如果再加上使用类似python中的`virtualenv`包的效果, 这种子shell嵌套就变的更多, 更复杂了.

```shell
#!/bin/bash

echo 'i am in' $(pwd) 'at the beginning'
su - general
echo 'i am in' $(pwd) 'now'
source /home/general/virpython/bin/activate
## 在服务器上以general身份, 在virtualenv环境下安装了django, 用以验证是否曾经执行到这一句
pip freeze | grep -i django
exit
echo 'i am in' $(pwd) 'finally'
```

将上面的脚本保存为`test2.sh`, 以root身份执行, 结果如下

```
[root@b14e517d408b tmp]# ./test2.sh 
i am in /tmp at the beginning
[general@b14e517d408b ~]$ exit
logout
i am in /tmp now
Django==1.8
[root@b14e517d408b tmp]# pip freeze | grep -i django
[root@b14e517d408b tmp]# 
```

我们看到, 在root下直接执行`pip freeze`是没有输出的, 即root下没有安装django包. 而输出了'i am in /tmp now'后理论上流程已经回到原shell, 当前用户是root, 但却输出了'Django==1.8'.

原因是, 用root身份执行`source /home/general/virpython/bin/activate`, 也可以查询到该虚拟python环境下的包, 这是没有问题的. 但是, 如果在这样的脚本中执行`pip install`命令, 安装第三方包, 到时general用户就没法卸载了, 因为没有权限. 所以并不是一个好的做法.

------

### 1.2 正确实现方法

```shell
#!/bin/bash

echo 'i am in' $(pwd) 'at the beginning'
su - general << EOF
echo 'i am in' $(pwd) 'now'
EOF
echo 'i am in' $(pwd) 'finally'
```

```
[root@b14e517d408b tmp]# ./test3.sh 
i am root at the beginning
i am root now
i am root finally
[root@b14e517d408b tmp]# 
```