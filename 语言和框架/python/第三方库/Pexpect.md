# Python第三方库-pexpect

参考文章

1. [探索 Pexpect，第 1 部分：剖析 Pexpect](https://www.ibm.com/developerworks/cn/linux/l-cn-pexpect1/)
    - Pexpect 并不与 shell 的元字符例如重定向符号`>`,`>>` 、管道`|` , 还有通配符`*`等做交互, 所以当想运行一个带有管道的命令时必须另外启动一个 shell.
2. [探索 Pexpect，第 2 部分：Pexpect 的实例分析](https://www.ibm.com/developerworks/cn/linux/l-cn-pexpect2/)
3. [python的pexpect模块的应用](https://blog.csdn.net/aA518189/article/details/84640701)
    - pexpect是一个纯pyth实现的模块(即不依赖系统中的 tcl/tcl-devel 库)
    - 使用`python setup.py install`安装`pexpect`, 缺少`ptyprocess`的解决方法
