# 多线程threading库使用.6.获取线程id[thread id inent]

参考文章

1. [python下使用ctypes获取threading线程id](http://xiaorui.cc/archives/3017)
2. [How to obtain a Thread id in Python?](https://stackoverflow.com/questions/919897/how-to-obtain-a-thread-id-in-python)

`threading.Thread`实例对象的的`name`跟`ident`属性, 只是个标识, 与`ps -efT`命令查到的线程id是两码事.


