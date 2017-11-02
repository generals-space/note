# nodejs异步IO的实现

原文链接

1. [nodejs异步IO的实现](https://cnodejs.org/topic/4f16442ccae1f4aa2700113b)

nodejs的核心之一就是非阻塞的异步IO，于是想知道它是怎么实现的，挖了下nodejs源码，找到些答案，在此跟大家分享下。首先，我用了一段js代码`test-fs-read.js`做测试，代码如下：

```js
var path = require('path'),
fs = require('fs'),
filepath = path.join(__dirname, 'experiment.log'),
fd = fs.openSync(filepath, 'r');

fs.read(fd, 1210241024, 0, 'utf-8', function(err, str, bytesRead) {
    console.log('[main thread] execute read callback');
});
console.log('[main thread] execute operation after read');
```

这段代码的异步IO操作就在`fs.read`的调用上，读取的`experiment.log`是一个12M的文本文件，所谓的异步，大家大概能想得到运行时会先打印

```
[main thread] execute operation after read
```

然后打印回调函数中的

```
[main thread] execute read callback
```

但大家也许还听说过，nodejs是单线程的，那又是怎么实现异步IO的呢？读文件操作是在哪里执行的呢？读完又是怎么调用回调函数的呢？猜想读文件可能是在另一个线程中完成的，读完后通过事件通知nodejs执行回调。为了一探究竟，debug了一把`nodejs`和`libeio`源码，重新编译后，运行测试代码`node test-fs-read.js`，输出如下：

![](https://gitimg.generals.space/53b297cc15608444606ea3f31e7167ce.jpg)

可以看到，nodejs的IO操作是通过调用`libeio`库完成的，debug从`fs.read`开始，js代码经过v8编译后，`fs.read`会调用`node_file.cc`中的Read方法，测试代码的运行经历了以下步骤：

1. `node_file.cc`中的Read方法调用`libeio（eio.c）`的`eio_read`， read请求被放入请求队列`req_queue`中。

2. 主线程创建了1个eio线程，此时主线程的read调用返回。

3. eio线程从req_queue中取出1个请求，开始执行read IO

4. 主线程继续执行read调用后的其它操作。

5. 主线程`poll eio`，从响应队列`res_queue`取已经完成的请求，此时`res_queue为`空，主线程`stop poll`

6. eio线程完成了read IO，read请求被放入响应队列`res_queue`中，并且向主线程发送`libev`事件`want_poll`（通过主线程初始化eio时提供的回调函数）。

7. eio线程从`req_queue`中取下一个请求，此时已经没有请求。

8. 主线程响应`want_poll`事件，从`res_queue`中取出1个请求，取出请求后`res_queue`变为空，主线程发送`done_poll`事件。

9. 主线程执行请求的callback函数。

还需要说明的是，当同时有多个IO请求时，主线程会创建多个eio线程，以提高IO请求的处理速度。
为了更清晰的看到nodejs的IO执行过程，图示如下，序号仅用来标示流程，与上述步骤序号并无对应关系。

![](https://gitimg.generals.space/dc0674db2bd387a69ffda5c3fd45c859.jpg)

最后总结几条，不当之处还请大家指正。

1. nodejs通过`libev`事件得到IO执行状态，而不是轮询，提高了CPU利用率。

2. **虽然nodejs是单线程的，但它的IO操作是多线程的**，多个IO请求会创建多个libeio线程（最多4个），使通常情况的IO操作性能得到提高。

3. 但是当IO操作情况比较复杂的时候，有可能造成线程竞争状态，导致IO性能降低；而且libeio最多创建4个线程，当同时有大量IO请求时，实际性能有待测量。另外，由于每个IO请求对应一个libeio的数据结构，当同时有大量IO操作驻留在系统中时候，会增加内存开销。

4. `libeio`为了实现异步IO功能，带来了额外的管理，当IO数据量比较小的时候，整体性能不一定比同步IO好。

------

xuancheng 3楼

看了这篇文章很受启发。因此去了解了一下node.js异步文件操作的情况，补充几点：

1. node使用libeio实现非阻塞异步文件操作，但实际上libeio是通过多线程的方式，在标准的阻塞式IO上模拟非阻塞异步。

2. 为什么node没有使用原生态的异步IO API（AIO）？原因可能是为了实现跨平台的兼容性（主要针对Windows，具体参见：http://groups.google.com/group/nodejs/browse_thread/thread/aff97d25c59f6f2d）

3. libeio的4个线程限制是默认配置，可以对此配置进行修改.