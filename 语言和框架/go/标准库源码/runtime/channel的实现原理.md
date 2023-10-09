# channel的实现原理

参考文章

1. [Golang-Channel原理解析](https://blog.csdn.net/u010853261/article/details/85231944)
    - 讲解channel的底层数据结构`hchan`, 图形化展示存取数据, 及数据满等情况下channel的工作流程
    - channel源码分析
2. [【我的架构师之路】- golang源码分析之channel的底层实现](https://blog.csdn.net/qq_25870633/article/details/83388952)
    - channel源码分析
    - send/recv/close的执行流程图, 很清晰
3. [深入理解channel：设计+源码](https://www.jianshu.com/p/afe41fe1f672)
    - PPT资料 [Understanding Channels](https://speakerdeck.com/kavya719/understanding-channels)
4. [golang 中 channel 的非阻塞访问方法](https://www.jianshu.com/p/4f51fe7dad62)
    - select + default 可实现非阻塞读取.

作为协程间通信的通道, 每个协程在读/写channel对象时都会与其绑定, 在close()操作时, 内部会遍历所有读写的协程, 依次解除联系.

底层实现(`channel`是协程安全的, 因为其内部使用了mutex)

存储数据使用: 数组 + 锁, 同时配置两个索引(recvx, sendx随着读写操作移动, 借此实现循环队列), 锁用来保护 channel 对象中的所有字段.

读写等待协程(sendq/recvq)模拟了队列, 其实就是链表 + enqueue/dequque操作.

## 读操作

1. 首先判断是否为`select+default`的非阻塞读操作, 如果是, 则判断 channel 是否未关闭, 且为无缓冲且此时无写协程存在, 或是有缓冲但 channel 无数据, 那么直接返回(这种情况根本不可能读到数据), 无需加锁;
2. 加锁做后续操作
3. 判断如果 channel 已关闭且其中没有数据存在, 则返回默认值.
4. 如果 sendq 不为空, 则 channel 已满, 交由 recv() 函数处理.
    1. 如果 channel 无缓冲, recv 会调用 recvDirect() 直接从挂起的写协程中取数据
    2. 如果 channel 有缓冲, recv 需要自行从 buf 中取数据, 并处理 recvx/sendx 的位置关系
    3. 最终 recv() 需要将挂起的写协程唤醒.
5. 如果 channel 未满, 且有数据, 可直接读取
6. 如果 channel 没有数据, 且读操作为非阻塞, 则直接返回默认值.
7. 如果 channel 没有数据, 且读操作为阻塞, 那么挂起此读协程.

## 写操作

1. 首先判断是否为`select+default`的非阻塞写操作, 如果是, 则判断 channel 是否未关闭, 且为无缓冲队列且没有读协程等待, 或有缓冲但 channel 已满, 那么直接返回(这种情况根本不可能成功写入), 无需加锁;
2. 加锁做后续操作.
3. 判断 channel 是否关闭, 如果是则panic
4. recvq 队列非空, 表示有读协程, 调用 send() 函数处理
5. 如果 recvq 为空, 且 channel 未满, 可以写入.
6. 如果 channel 已满, 但写协程为非阻塞, 则直接返回写入失败
7. 如果写协程可阻塞, 则阻塞该写协程, 等待被唤醒
