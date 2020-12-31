# Python实现令牌桶[限速 ratelimit]

<!key!>: {A03F8671-228A-4BE2-821F-05E983245D7B}

<!link!>: {5E8AD061-7897-4D0F-A0DA-232CD84367FE}

参考文章

1. [15行Python代码，帮你理解令牌桶算法](https://juejin.cn/post/6844903580923609102)
    - 关于突发的网络高峰情况, 提出的对于令牌桶算法的优化思路.

关于令牌桶算法的工作原理, 可以见另一篇 golang 的文章, 我觉得已经够清晰了. 

至于参考文章1, ta提供了一个 Python 实现的简单令牌桶, 目前没找到比较权威的库, 就先拿来用了.

```py
import time

'''
参考文章: [15行Python代码，帮你理解令牌桶算法](https://juejin.cn/post/6844903580923609102)

令牌桶需要以一定的速度生成令牌放入桶中, 当程序要发送数据时, 再从桶中取出令牌. 

这里似乎有点问题, 如果我们使用一个死循环, 来不停地发放令牌, 程序就被阻塞住了, 除非使用多线程, 有没有更好的办法?

我们可以在取令牌的时候, 用现在的时间减去上次取令牌的时间, 乘以令牌的发放速度, 计算出桶里可以取的令牌数量(当然不能超过桶的大小), 从而避免循环发放的逻辑. 

'''
import time

class TokenBucket():
    def __init__(self, rate, capacity):
        '''
        @param rate: 令牌发放速度, 单位: 秒

        @param capacity: 桶的大小
        '''
        self._rate = rate
        self._capacity = capacity
        ## 初始时, 令牌桶就是满的
        self._current_amount = capacity
        self._last_consume_time = int(time.time())

    def consume(self, token_amount):
        '''
        @param token_amount: 需要的令牌数
        '''
        ## 计算从上次发送到这次发送, 新发放的令牌数量
        increment = (int(time.time()) - self._last_consume_time) * self._rate 
        ## 令牌数量不能超过桶的容量, 多余的会溢出
        self._current_amount = min(increment + self._current_amount, self._capacity) 

        ## 如果没有足够的令牌, 则不能发送数据
        if token_amount > self._current_amount: return False
        ## 如果有, 则更新 _last_consume_time 时间
        self._last_consume_time = int(time.time())
        self._current_amount -= token_amount
        return True
```

对应的测试代码

```py
bucket = TokenBucket(1, 3)
duration = 5
start = time.time()

while True:
    now = time.time()
    if now - start > duration: break

    if bucket.consume(1):
        print(time.time())
    else:
        time.sleep(1)
```

结果如下

```
1609331474.296015
1609331474.296076
1609331474.296114
1609331475.298961
1609331476.2993941
1609331477.30401
1609331478.3041239
```

示例程序只运行5秒, 由于初始时令牌桶就是满的, 所以在第1秒立刻取出了3枚令牌, 之后的循环中只能等待每秒从桶中取出一个令牌了.

