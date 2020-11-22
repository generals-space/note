# Executors.newCachedThreadPool线程池.3.返回值

<!key!>: {C2F23522-2EC1-4D7B-B2C4-503B3BFADACA}

参考文章

1. [Java并发包——线程安全的Map相关类](https://www.cnblogs.com/shamao/p/11093917.html)
    - `ConcurrentHashMap`线程安全的 Map, 底层源码解析
2. [ExecutorService实现获取线程返回值](https://blog.csdn.net/qq_24448899/article/details/79066717)
    - 实现`Callable`接口
    - `future.get()`直接获取目标任务的执行结果, 但会阻塞, 不适合进行结果汇总
3. [JAVA 线程池之Callable返回结果](https://www.cnblogs.com/hapjin/p/7599189.html)
    - `executor.getCompletedTaskCount() < resultList.size()`判断线程池内任务的完成情况, 直到所有任务都完成再进行汇总
    - 如果有效且可靠的话, 应该是最合适的方法.
4. [【搞代码】Java线程池对多个任务的处理结果进行汇总](https://blog.csdn.net/qq_36819098/article/details/99092697)
    - 遍历`future`结果, 依次检测任务的执行情况, 有点笨, 不过应该有效.

线程池只负责执行任务, 不负责返回结果, 而且一般场景下也没有获取结果的必要(都是把任务放到线程池中后台运行就行了).

但是我遇到的场景, 仅仅是希望通过线程池加快原本顺序的操作, 之后需要根据其中的操作结果进行下一步. 

我首先想到的是, 用一个变量存储结果(就像 golang 那样, 用 channel 进行通信). 加锁当然是可以的, 但我遇到的是场景是, 需要在 http 服务的一个处理接口中, 开线程池去执行一些任务. 

```java
    @GetMapping("/")
	public String home(){
		// 开线程池
    }
```

可能你要问, 为什么不在其所属类中创建一个公用的线程池, 统一执行任务呢? 

这就是我上面说的, 每来一个请求, 我就需要执行一些任务, 如果用共用的线程池对象, 放进去的任务我要怎么区分是属于哪个请求呢? 

就算可以对 worker() 传入参数根据请求id(如果有的话)进行过滤, 但是线程池对象的`awaitTermination()`方法是用来等待池中所有任务完成后才能返回的, 而我只想得到当前请求中任务的结果.

所以每处理一个请求, 就创建一个线程池是合理的, 只要线程池能在保证的时间内执行完毕, 再回收就可以了.

接下来就是锁的问题. 线程池都用局部的了, 锁肯定也没必要在类属性中创建了...吗???

但是 Java 中, 要放进线程池内执行的 worker 是需要单独建一个方法的, 那这个方法该怎么拿到`Lock`对象呢...?

------

网上关于这个问题几乎所有的文章都是, 要实现`Callable`接口的类, 把要执行的任务写在`call()`方法, 直接返回你到返回的值. 然后在将任务对象放入线程池时, 使用`submit()`方法提交, 使用`Future<XXX>`对象承接结果(这里的XXX类型与在`call()`方法中返回的结果类型保持一致).

但是之后需要一直通过`getCompletedTaskCount()`方法查询线程池中的任务有没有完成, 不能再使用`awaitTermination()`了...

以后如果实在没有办法再使用这种, 现在我们尝试使用自己的方式.

## Lock 对象当作参数传入 worker 函数

```java
package com.example.demo;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

import com.google.common.collect.Maps;

public class DemoApplication {
	private static void worker(String workerName, Lock lock, Map<String, Boolean>resultMap){
		try {
			// 就是其他语言中的 sleep, 单位为 ms
			Thread.sleep(5000);
			lock.lock();
			if(workerName.equals("worker5")){
				resultMap.put(workerName, false);
			} else {
				resultMap.put(workerName, true);
			}
		} catch (InterruptedException e){

		} finally {
			lock.unlock();
		}
	}

	public static void main(String[] args) {
		LocalDateTime start = LocalDateTime.now();
		System.out.println("start: " + start);

		ExecutorService pool = Executors.newCachedThreadPool();
		// 不知道为什么只能用 final 修饰
		final Lock counterLock = new ReentrantLock();

		String worker = "worker";
		Map<String, Boolean> resultMap = Maps.newHashMap();
		
		for(int i = 0; i < 100; i ++){
			String workerName = worker + i;
			// runnable 行内定义函数的方式, 类似于 python 中的 lamda 表达式
			Runnable runnable = () -> worker(workerName, counterLock, resultMap);
			pool.execute(runnable);
		}
		// shutdown() 关闭通道, 之后无法再往里添加任务, 并非关闭/回收线程池的意思...
		// 类似的话, 应该像是其他语言中的 pool.start() 
		pool.shutdown();

		try{
			// awaitTermination() 类似其他语言线程池中的 join() 或 wait(),
			// 等待线程池中所有线程运行完成, 还可以设置超时时间.
			if(!pool.awaitTermination(10, TimeUnit.SECONDS)){
				System.out.println("timeout...");
				pool.shutdownNow();
			}else {
				LocalDateTime end = LocalDateTime.now();
				System.out.println("end: " + end);
				// 查看结果
				System.out.println(resultMap.size());
				resultMap.forEach((k, v) -> {
					// 打印出 false 的那个线程任务
					if(!v) System.out.println(k);
				});
			}
		} catch (InterruptedException e){
			System.out.println("exception...");
			pool.shutdownNow();
		}
	}
}

```

总感觉这样有些不对劲, 虽然 Lock 对象本身也是个引用类型吧...

> 其实不加锁这个示例也一般不会跑出异常问题来...

## 线程安全的对象 ConcurrentHashMap

直接用线程安全的对象存储结果, 不用加锁.(主要还是觉得将`Lock`对象当作参数传入还是太少见...)

不过网上关于这个类型的文章, 更多的是底层原理解析, 连个示例代码都没有. 而且在搜索的时候, 竟然出现"ConcurrentHashMap 读为什么不加锁"这种问题.

![](https://gitee.com/generals-space/gitimg/raw/master/4c8452be34225325381961661c2cff34.png)

让人深深怀疑这东西是不是真的线程安全的...

后来我询问了搞 Java 的小伙伴, 他说关于"读操作不加锁"是针对其底层实现的问题, 算是面试时的经典问题...md神经病Java

```java
package com.example.demo;

import java.time.LocalDateTime;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

public class DemoApplication {
	private static void worker(String workerName, ConcurrentHashMap<String, Boolean> resultMap){
		try {
			// 就是其他语言中的 sleep, 单位为 ms
			Thread.sleep(5000);
			if(workerName.equals("worker5")){
				resultMap.put(workerName, false);
			} else {
				resultMap.put(workerName, true);
			}
		} catch (InterruptedException e){} 
	}

	public static void main(String[] args) {
		LocalDateTime start = LocalDateTime.now();
		System.out.println("start: " + start);

		ExecutorService pool = Executors.newCachedThreadPool();

		String worker = "worker";
		ConcurrentHashMap<String, Boolean> resultMap = new ConcurrentHashMap<>();
		for(int i = 0; i < 100; i ++){
			String workerName = worker + i;
			// runnable 行内定义函数的方式, 类似于 python 中的 lamda 表达式
			Runnable runnable = () -> worker(workerName, resultMap);
			pool.execute(runnable);
		}
		// shutdown() 关闭通道, 之后无法再往里添加任务, 并非关闭/回收线程池的意思...
		// 类似的话, 应该像是其他语言中的 pool.start() 
		pool.shutdown();

		try{
			// awaitTermination() 类似其他语言线程池中的 join() 或 wait(),
			// 等待线程池中所有线程运行完成, 还可以设置超时时间.
			if(!pool.awaitTermination(10, TimeUnit.SECONDS)){
				System.out.println("timeout...");
				pool.shutdownNow();
			}else {
				LocalDateTime end = LocalDateTime.now();
				System.out.println("end: " + end);
				// 查看结果
				System.out.println(resultMap.size());
				resultMap.forEach((k, v) -> {
					// 打印出 false 的那个线程任务
					if(!v) System.out.println(k);
				});
			}
		} catch (InterruptedException e){
			System.out.println("exception...");
			pool.shutdownNow();
		}
	}
}
```
