# Python-bytes与str相互转换

参考文章

1. [Python3中bytes类型转换为str类型](https://www.jb51.net/article/148052.htm)
2. [python2与python3的bytes问题](https://www.cnblogs.com/lshedward/p/9926150.html)
    - `python2`中, `bytes`和`str`是一回事
    - `python3`强制使用`unicode`
3. [python 之路，致那些年，我们依然没搞明白的编码](https://www.cnblogs.com/alex3714/articles/7550940.html)
    - 最全面的讲解, 值得收藏.

```py
import subprocess
result = subprocess.Popen(['ls', '/tmp'], stdout=subprocess.PIPE)
print(result.stdout.read())
## b'kube-ovn\nplayground\nproject\n'
```

`stdout.read()`的结果为`b''`, `bytes`类型. 

如果能确定目标命令的输出结果一定是字符串(也有些命令输出的就是二进制, 比如`etcdctl get...`), 可以考虑将结果转换一下, 毕竟换行太难看了.

```py
src = b'kube-ovn\nplayground\nproject\n'
print(type(src)) ## <class 'bytes'>
```

## python2

python2中, 虽然有`bytes()`和`str()`函数, 但是`bytes`和`str`是一回事.

## python3

### bytes -> str

1. `str(src, encoding='utf-8')`
2. `bytes.decode(src)`

```py
>>> src = b'kube-ovn\nplayground\nproject\n'
>>> print(str(src, encoding='utf-8'))
kube-ovn
playground
project

>>> print(bytes.decode(src))
kube-ovn
playground
project
```

### str -> bytes

这种的使用场景比较少吧...我是照着上面反推出来的

1. `bytes(string, encoding='utf-8')`
2. 不过`bytes`没有`encode()`方法.

```py
bytes('kube-ovn\nplayground\nproject\n', encoding='utf-8')
b'kube-ovn\nplayground\nproject\n'
```
