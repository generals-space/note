# Python执行shell命令

参考文章

[python中如何调用shell 中OS.SYSTEM等方法](http://blog.csdn.net/gray13/article/details/7044453)

[python执行系统命令的方法 ：os.system()，subprocess.popen()，commands](http://xingyunbaijunwei.blog.163.com/blog/static/76538067201341343433333/)

[python os模块进程函数](http://www.cnblogs.com/nisen/p/6060355.html)

[Python模块整理(三)：子进程模块subprocess](http://ipseek.blog.51cto.com/1041109/807513)

os.system(cmd): 

在一个子shell中运行command命令, 并返回command命令执行完毕后的退出状态. 这实际上是使用C标准库函数system()实现的. 这个函数在执行command命令时需要重新打开一个终端, 并且无法保存command命令的执行结果. 

os.popen()得到的是file read对象, 需要对其进行读取`read()`的操作才可以看到执行的输出, 适合`ps`, `ls`这种命令.