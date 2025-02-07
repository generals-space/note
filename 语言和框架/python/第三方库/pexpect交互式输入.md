# Python第三方库-pexpect

参考文章

1. [探索 Pexpect，第 1 部分：剖析 Pexpect](https://www.ibm.com/developerworks/cn/linux/l-cn-pexpect1/)
    - Pexpect 并不与 shell 的元字符例如重定向符号`>`,`>>` 、管道`|` , 还有通配符`*`等做交互, 所以当想运行一个带有管道的命令时必须另外启动一个 shell.
2. [探索 Pexpect，第 2 部分：Pexpect 的实例分析](https://www.ibm.com/developerworks/cn/linux/l-cn-pexpect2/)
3. [python的pexpect模块的应用](https://blog.csdn.net/aA518189/article/details/84640701)
    - pexpect是一个纯pyth实现的模块(即不依赖系统中的 tcl/tcl-devel 库)
    - 使用`python setup.py install`安装`pexpect`, 缺少`ptyprocess`的解决方法

pexpect 4.x 只支持 python3, 3.x 还支持 python2.

以下示例是 python2 的一个示例.

```py
#! encoding: utf-8
import sys
import pexpect

child = pexpect.spawn('/usr/bin/ssh-keygen', ['-t', 'rsa'])
## 将命令执行结果重定向到标准输出.
## 执行目标命令的初始输出不会立刻出现, 只有遇到第一个expect() 才会打印出来.
child.logfile = sys. stdout

try:
    child.expect('Enter file in which to save the key')
    child.sendline(' /root/myssh/myrsa')
    child.expect(' Enter passphrase')
    child.sendline()

    child.expect('Enter same passphrase again')
    child.sendline()

    ## 如果`/root/myssh`这个目录不存在, 则会报错.
    ## 但是我们需要 expect 成功的标记, 否则不等打印出失败的信息程序就直接退出了, 也无法捕获到异常.
    child.expect('The key fingerprint is')
    child.close()
except Exception as e:
    print('pexpect got an exception: ', e)
    child.close()
```
