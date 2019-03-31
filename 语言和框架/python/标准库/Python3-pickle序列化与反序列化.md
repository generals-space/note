# Python3-pickle序列化与反序列化

<!key!>: {664fdd98-537c-11e9-b398-aaaa0008a014}

<!link!>: {c2f8cf90-5387-11e9-ae66-aaaa0008a014}

参考文章

1. [Python数据存储：pickle模块的使用讲解](https://www.cnblogs.com/fmgao-technology/p/9078918.html)

2. [python 序列化之JSON和pickle详解](https://www.cnblogs.com/tkqasn/p/6005025.html)

之前一直不知道pickle有什么用, 因为不明白把对象序列化存储到文件有什么可用场景.

后来在遇到使用redis作分布式session缓存时, 就考虑到, 在session中存储的那些对象, 比如日期, 比如自定义用户对象, 在redis中是怎么存储的呢? 服务重启时又是怎么从redis将数据读取出来的呢?

这种情况下就需要用到pickle.

我们可以将`pickle`与`json`对比一下, 两者几乎作用相同, `json`用于将对象序列化成字符串, `pickle`用于将对象序列化成字节数组. 

两者都有`dump`/`dumps`, `load`/`loads`两对方法. `dump/load`用于序列化/反序列化文件, `dumps/loads`用于序列化/反序列化字节数组/字符串.

参考文章1中对于pickle的两组常用方法简单介绍了一下, 还算清晰(如果你已经清楚地知道了使用场景的话).

不过两者的能力范围是不同的, 参考文章2中说 **JSON只能处理基本数据类型, pickle能处理所有Python的数据类型.**

以如下代码为例(不靠读写文件).

```py
import pickle
import json
from datetime import datetime

obj = {
    'name': 'general',
    'age': 21,
    'birthday': datetime.now(),
}

byteArray = pickle.dumps(obj)
print(byteArray)                ## b'\xxx\xxx\xxx'
obj2 = pickle.loads(byteArray)
print(obj2)                     ## {'name': 'general', 'age': 21, 'birthday': datetime.datetime(2019, 3, 31, 15, 10, 11, 526594)}

print(json.dumps(obj))          ## TypeError: Object of type datetime is not JSON serializable
```

`pickle.dumps()`可以得到字节数组, `pickle.loads()`可以将对象反序列化回来, 且`obj2['birthday']`类型仍为`datetime`.

但是`json`模块就不行了, 因为ta不知道`datetime`对象如何变成字符串, 总得给个格式吧? 并且反序列化读取什么样的字符串才能创建回`datetime`对象呢?

不过也不是一点办法都没有, 毕竟只要给出一个序列化和反序列化都遵循的格式就可以了, 参考文章2给出了这个方法.

`pickle`模块在某种情况下, 也是需要这方法的, 上面也说了, pickle能处理所有Python的数据类型(不过`queue.Queue`貌似不行)), 是因为内置类型都提供了相应的格式. 

但是如果是自定义的对象呢?

说实话, 这种情况我宁愿直接存储到数据库, 序列化方法我自己写...
