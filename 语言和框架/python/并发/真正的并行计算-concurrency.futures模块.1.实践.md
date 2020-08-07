# 真正的并行计算-concurrency.futures模块

参考文章

1. [python异步并发模块concurrent.futures简析](http://lovesoo.org/analysis-of-asynchronous-concurrent-python-module-concurrent-futures.html)

2. [python concurrent.futures](https://www.cnblogs.com/kangoroo/p/7628092.html)

1. `python 3.x`中自带了`concurrent.futures`模块.
2. `python 2.7`需要安装`futures`模块, 使用命令`pip install futures`安装即可.

`concurrent`包只有`futures`一个子模块, 而这个模块下只有两个成员: `ProcessPoolExecutor`, `ThreadPoolExecutor`.

参考文章2提供了一个求最大公约数(`greatest common divisor`)的示例, 这是一个计算密集型函数. 理论上只有多进程模型才可以利用多核优势能够尽快完成指定的任务, 下面可以看一下这个示例.

```py
def gcd(pair):
    '''
    @function: 求最大公约数(greatest common divisor)
    @param: pair, 整型数据二元组, 求这两个值的最大公约数.
    '''
    a, b = pair
    low = min(a, b)
    ## 从low开始倒数到0, 每次减1.
    for i in range(low, 0, -1):
        if a % i == 0 and b % i == 0:
            return i

## 目标任务
numbers = [
    (1963309, 2265973), (1879675, 2493670), (2030677, 3814172),
    (1551645, 2229620), (1988912, 4736670), (2198964, 7876293)
]
```

测试环境: MacBookPro2018, 双核4线程.

```console
$ sysctl hw.physicalcpu
hw.physicalcpu: 2
$ sysctl hw.logicalcpu
hw.logicalcpu: 4
```

## 1. 常规操作

```py
import time

start = time.time()
results = list(map(gcd, numbers)) ## 顺序进行
end = time.time()
print(results) ## [1, 5, 1, 5, 2, 3]
print('Took %.3f seconds.' % (end - start)) 
```

运行3次, 花费的时间分别为: `0.967s`, `0.827s`, `0.882s`.

## 2. 多线程ThreadPoolExecutor

```py
import time
from concurrent.futures import ThreadPoolExecutor

start = time.time()
pool = ThreadPoolExecutor(max_workers=4)
results = list(pool.map(gcd, numbers))
end = time.time()
print(results) ## [1, 5, 1, 5, 2, 3]
print('Took %.3f seconds.' % (end - start)) 
```

运行3次, 花费的时间分别为: `0.916s`, `0.892s`, `0.875s`.

## 3. 多进程ProcessPoolExecutor

```py
import time
from concurrent.futures import ProcessPoolExecutor

start = time.time()
pool = ProcessPoolExecutor(max_workers=4)
results = list(pool.map(gcd, numbers))
end = time.time()
print(results) ## [1, 5, 1, 5, 2, 3]
print('Took %.3f seconds.' % (end - start)) 
```

运行3次, 花费的时间分别为: `0.480s`, `0.475s`, `0.467s`.

------

## 4. 总结

可以看到, 计算密集型应用, 使用多线程模型的效果不大, 使用多进程模型的效果倒是很明显. 但是多线程/多进程并行执行时, 返回结果的顺序竟然都是确定的, 简直不可思议.

`ProcessPoolExecutor`类会利用`multiprocessing`模块所提供的底层机制, 完成下列操作: 

1. 把`numbers`列表中的每一项输入数据都传给`map`. 
2. 用`pickle`模块对数据进行序列化, 将其变成二进制形式. 
3. 通过本地套接字, 将序列化之后的数据从煮解释器所在的进程, 发送到子解释器所在的进程. 
4. 在子进程中, 用`pickle`对二进制数据进行反序列化, 将其还原成python对象. 
5. 引入包含`gcd`函数的python模块. 
6. 各个子进程并行的对各自的输入数据进行计算. 
7. 对运行的结果进行序列化操作, 将其转变成字节. 
8. 将这些字节通过socket复制到主进程之中. 
9. 主进程对这些字节执行反序列化操作, 将其还原成python对象. 
10. 最后, 把每个子进程所求出的计算结果合并到一份列表之中, 并返回给调用者. 

`multiprocessing`开销比较大, 原因就在于: 主进程和子进程之间通信, 必须进行序列化和反序列化的操作. 
