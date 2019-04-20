# Python-subprocess(二)低级接口Popen()

参考文章

1. [python子进程模块subprocess详解与应用实例 之一](http://blog.csdn.net/fireroll/article/details/39153831)

2. [python子进程模块subprocess详解与应用实例 之二](http://blog.csdn.net/fireroll/article/details/39153947)

3. [python子进程模块subprocess详解与应用实例 之三](http://blog.csdn.net/fireroll/article/details/39153991)

4. [subprocess官方文档](https://docs.python.org/2.7/library/subprocess.html)

5. [Python subprocess模块学习总结](http://www.jb51.net/article/48086.htm)

6. [python subprocess.Popen 监控控制台输出](http://blog.csdn.net/mldxs/article/details/8555819)

## 1. Popen()的使用方法

上面的几个函数都是基于Popen()的封装. 这些封装的目的在于让我们容易使用子进程. 当我们想要更个性化我们的需求的时候, 就要转向Popen类, 该类生成的对象用来代表子进程. 

与上面的封装不同, Popen对象创建后, 主程序不会自动等待子进程完成. 我们必须调用对象的wait()方法, 父进程才会等待 (也就是阻塞block), 举例: 

```py
#!/usr/bin/env python
import subprocess

child = subprocess.call('ping -c 4 172.16.3.206', shell = True)
print('child complete')
```

你会发现, 在`call()`生成的ping子进程执行完成之前, print是无法打印信息的, 即发生了阻塞.

但是同样的调用方法, Popen()函数就不会阻塞, print几乎是立刻就打印出了信息.

```py
#!/usr/bin/env python
import subprocess

child = subprocess.Popen('ping -c 4 172.16.3.206', shell = True)
print('child complete')
```

为了让Popen()生成的子进程能够阻塞, 以便取到目标程序的输出, 我们需要使用`wait()`函数.

```py
#!/usr/bin/env python

import subprocess

child = subprocess.Popen('ping -c 4 172.16.3.206', shell = True)
child.wait()
print('child complete')
```

这下, 程序的运行结果就和call()函数相同了. 我们也发现了, Popen()的返回值是一个对象, 它拥有成员方法`wait()`. 实际上它正是suprocess生成的子进程对象.

我们可以直接操作这个子进程对象.

```
child.poll() # 检查子进程状态, 未结束返回None
child.kill() # 终止子进程
child.send_signal(signal) # 向子进程发送信号signal
child.terminate() # 终止子进程
```

这些操作我想应该是把子进程当成守护进程启动的了吧?

好吧我们来试试

```py
#!/usr/bin/env python

import subprocess
import time
child = subprocess.Popen('ping 172.16.3.206', shell = True, stdout = subprocess.PIPE, stderr = subprocess.PIPE)
time.sleep(10)
child.terminate()
print('child complete')
```

就这样, 执行的时候ping子进程会在后台执行, 并且未绑定标准IO, 不过我们可以使用child子进程对象对它进行操作.

再说一句, ping子进程貌似还脱离了当前session group, 因为当我关闭执行终端时, ping进程依然在运行. 竟然真的是把它当成守护进程运行的.

## 2. Popen获取目标命令输出

上面我们对Popen()的执行方式, 使得目标命令的输出都是打印在标准输出中, 却没有办法赋值到某个变量中. 

其实可以使用`subprocess.PIPE`选项将子进程与父进程的标准输入输出连接起来. 然后通过子进程对象的stdout读取. 示例如下

```py
#!/usr/bin/env python

import subprocess
import time
child = subprocess.Popen('ping -c 4 172.16.3.206', shell = True, stdout = subprocess.PIPE, stderr = subprocess.STDOUT)
child.wait()
print('%s' % child.stdout.read())
print('child complete')
```

由于要读取ping子进程的全部输出, 所以只能使用wait()函数等待它执行完成后再读取. `stderr = subprocess.STDOUT`这个参数非常重要, 它表示标准错误使用与标准输出相同的方法处理, 不然当目标子进程输出到标准错误时就无法捕获其信息了.

但是, 发现没有, ping子进程实际有4条输出, 至少每秒一条, 但是由于`wait()`存在, 我们必须得等到ping子进程完全结束才能读取它的输出. 但是如果目标命令的输出太多, 超过了缓冲区的大小的话, 就不知道会发生什么事情了.

正好我们可以使用poll方法检测子进程的运行状态, 然后实时读取它的输出, 直到结束.

```py
#!/usr/bin/env python

import subprocess
import time
child = subprocess.Popen('ping -c 4 172.16.3.206', shell = True, stdout = subprocess.PIPE, stderr = subprocess.PIPE)
while True:
    line = child.stdout.readline()
    print(line)
    if child.poll() != None:
        break
print('child complete')
```

不过好像好了点东西, 下面的那些没了

```
--- 172.16.3.206 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3000ms
rtt min/avg/max/mdev = 0.318/0.338/0.372/0.028 ms
```

不明觉厉...