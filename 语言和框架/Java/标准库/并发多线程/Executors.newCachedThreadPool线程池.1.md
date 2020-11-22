# Executors.newCachedThreadPool线程池

参考文章

1. [ExecutorService等待线程完成后优雅结束](https://blog.csdn.net/feilang00/article/details/87930662)
2. [java如何使用ExecutorService关闭线程池？](https://www.yisu.com/zixun/87749.html)

`Thread`多线程可以将任务放到后台去处理, 创建多个并发执行, 但是`Thread`对象的`join`方法只能等待单个线程, 如果希望等待所有任务执行完毕, 接下来进行下一步操作的话, 还是开线程池比较好.

```java
package com.example.demo;

import java.time.LocalDateTime;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

public class DemoApplication {
	private static void worker(){
		try {
			// 就是其他语言中的 sleep, 单位为 ms
			Thread.sleep(5000);
		} catch (InterruptedException e){

		}
	}

	public static void main(String[] args) {
		LocalDateTime start = LocalDateTime.now();
		System.out.println("start: " + start);
		ExecutorService pool = Executors.newCachedThreadPool();
		String name = "123";
		for(int i = 0; i < 10; i ++){
			// runnable 行内定义函数的方式, 类似于 python 中的 lamda 表达式
			// 如果想要向`worker`中传参, 只要不使用 for 循环中的 i 变量, 其他变量都是可以的.
			Runnable runnable = () -> worker();
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
			}
		} catch (InterruptedException e){
			System.out.println("exception...");
			pool.shutdownNow();
		}
	}
}

```

输出如下

```
start: 2020-11-19T14:23:28.748
end: 2020-11-19T14:23:33.758
```
