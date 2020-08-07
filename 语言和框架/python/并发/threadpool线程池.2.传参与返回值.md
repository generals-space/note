# threadpool线程池.2.传参与返回值

参考文章

1. [python线程池（threadpool）模块使用笔记](http://www.cnblogs.com/xiaozi/p/6182990.html)

2. [threadpool官方文档](https://chrisarndt.de/projects/threadpool/api/)

## 1. 传参

当目标函数有多个参数时, 甚至需要`args`与`kwargs`这样灵活的参数时, 如何传入?

其实`makeRequests`第2个参数的列表成员类型是**元组**, 格式为`([], {})`, 元组的第1个成员为列表, 也是寻常意义的`args`, 第2个成员是一个字典, 也是`kwargs`.

```py
import threadpool
import time
pool = threadpool.ThreadPool(4)

def run(name, age = 18, ginder = 'male'):
    print('My name is %s ' % name)
    time.sleep(10)
    print('and my age is %d, I am a %s ' % (age, ginder))

name_list = [
    (['AA'], {'age': 24, 'ginder': 'male'}),
    (['BB'], {'ginder': 'male'}),
    (['CC'], {'age': 26, 'ginder': 'female'}),
    (['DD'], {'age': 32}),
]
reqs = threadpool.makeRequests(run, name_list)
[pool.putRequest(req) for req in reqs]
pool.wait()
```

```
My name is AA 
My name is BB 
My name is CC 
My name is DD 
## 这里会等待10s
and my age is 24, I am a male 
and my age is 18, I am a male 
and my age is 26, I am a female 
and my age is 32, I am a male 
```

## 2. 关于返回值与回调

返回值嘛...一般用于后台线程运行的代码是不会有返回值的, 不过如果确实有必要都线程的执行结果做处理, 就通过指定回调来完成(...那为啥不直接都写在一个函数里?). `makeRequests`函数的第3个参数可以指定一个回调, 在每个`func`任务执行完毕后调用, 并且会为其传入两个参数: 此线程所属`request`请求对象本身, 及其执行结果`result`. 

```py
import threadpool
import time
pool = threadpool.ThreadPool(4)

## 这里是回调...为什么不写在同一个函数里呢...???
def show(request, result):
    print('get result...')
    print(request)
    print(result)

def run(num):
    print('start...')
    time.sleep(num)
    print(num)
    return num * 2
num_list = [1, 3, 7]
reqs = threadpool.makeRequests(run, num_list, show)
[pool.putRequest(req) for req in reqs]
pool.wait()
```

------

这里我们需要再深究一下`threadpool`的源码, `wait`函数其实是一个不断执行`poll`函数的`while`死循环. 而`poll`函数的作用就是迅速检查一遍当前

## 2. 关于线程数量的弹性变化

`ThreadPool(poolsize)`会创建一个指定数量的线程池, 这个`poolsize`的值是固定的, 就算当前任务数量很少, 多余的线程也不会退出...

```py
import threadpool
import time
pool = threadpool.ThreadPool(50)

def run(name):
    print('start...')
    time.sleep(60)
    print(name)
name_list =['AA','BB','CC','DD']
reqs = threadpool.makeRequests(run, name_list)
[pool.putRequest(req) for req in reqs]
pool.wait()
```

执行这个脚本, `time.sleep(60)`会让线程挂起60s, 然后新开一个终端, 用`ps -efL`命令查看有关这个进程的线程信息, 你会看到包括主线程在内一共有`51`个线程在运行. 而我们实际需要运行的任务只有4个, 多余的46个线程都在空等而已...

`threadpool`的源文件中在`__main__`部分给出了一份示例代码. 其中就有动态更改线程池中数量的操作, 但是需要我们手动完成. 

好像有点不够智能哦...
