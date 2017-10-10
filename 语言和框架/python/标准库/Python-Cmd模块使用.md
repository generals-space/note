# Python-Cmd模块使用

[Cmd源码](https://hg.python.org/cpython/file/2.7/Lib/cmd.py)

## 1. 交互模式

```python
#!/bin/python
#!coding:utf-8
from cmd import Cmd
import sys

class CmdInterface(Cmd):
    def __init__(self):
        Cmd.__init__(self)
    ## 以'help_'为前缀的成员方法, 如help_test, 在执行'help test'时会被调用
    def help_exit(self):
        print('这是help_exit方法, 执行help exit命令时输出, 这里写的应该是对do_exit()方法的描述')
    def help_hello(self):
        print('这是help_hello方法, 执行help hello命令时输出, 这里写的应该是对do_hello()方法的描述. ok, hello方法接收一个参数, 是你想对他说hello的那个人')

    ## 以'do_'为前缀的成员方法, 如do_test, 在执行test命令时会被调用, args可以是多个, 名称随意
    def do_hello(self, args):
        print('this is hello: hello %s' % args)
    def do_exit(self, args):
        print('this is exit: good bye %s' % args)
        sys.exit()

if __name__ == '__main__':
    cmdObj = CmdInterface()
    cmdObj.cmdloop()
```

使用示例:

```
$ python cmder.py 
(Cmd) help hello
这是help_hello方法, 执行help hello命令时输出, 这里写的应该是对do_hello()方法的描述. ok, hello方法接收一个参数, 是你想对他说hello的那个人
(Cmd) hello boy  
this is hello: hello boy
(Cmd) help exit
这是help_exit方法, 执行help exit命令时输出, 这里写的应该是对do_exit()方法的描述
(Cmd) exit
this is exit: good bye 
```

### 改进1 

默认继承cmd库的Cmd类的对象, 在**交互模式**中不输入任何命令直接回车相当于重复输入上一条命令.

```
$ python cmdInterface.py 
(Cmd) help exit
这是help_exit方法, 执行help exit命令时输出, 这里写的应该是对do_exit()方法的描述
(Cmd) ## 这里直接输入回车
这是help_exit方法, 执行help exit命令时输出, 这里写的应该是对do_exit()方法的描述
(Cmd) exit
this is exit: good bye 
```

某些时候我们不想这么做, 那么我们可以在我们自定义的派生类中定义一个`emptyline`方法, 里面只写一句`pass`, 这样可以在空行回车的时候什么也不做.

```python
class CmdInterface(Cmd):
    def emptyline(self):
        pass
```

```
[root@localhost tmp]# python cmdInterface.py 
(Cmd) help exit
这是help_exit方法, 执行help exit命令时输出, 这里写的应该是对do_exit()方法的描述
(Cmd) ## 这里直接回车
(Cmd) exit
this is exit: good bye
```

好了, 现在认识一下cmd类中原本的`emptyline`方法. 如下

```python
    def emptyline(self):
        """Called when an empty line is entered in response to the prompt.

        If this method is not overridden, it repeats the last nonempty
        command entered.

        """
        if self.lastcmd:
            return self.onecmd(self.lastcmd)
```

说明`lastcmd`成员变量中保存有上一条执行的命令, 而`onecmd()`方法就是执行指定命令的方法.

那么再看看`onecmd()`方法的源码, 可以发现, 实际有效的就是代为执行`do_所输命令`的函数而已, 当然还包括选项参数的判断与缺省设置. 它里面还调用了解析传入的命令所用到的方法`parseline()`