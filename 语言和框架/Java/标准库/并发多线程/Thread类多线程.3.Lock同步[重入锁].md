# Thread类多线程.3.Lock同步[重入锁]

参考文章

1. [Java并发编程：Lock](https://www.cnblogs.com/dolphin0520/p/3923167.html)
	- java 5 开始在`java.util.concurrent.locks`包下提供了`Lock`
	- `synchronized`的缺陷
	- 对`Lock`下各种的获取锁的方法给出了示例代码, 包括为了释放锁的常用方式.
	- `ReentrantLock`可重入锁
	- `ReadWriteLock`读写锁
2. [Java锁机制Lock用法示例](https://www.jb51.net/article/146258.htm)
	- `Runnable`接口多线程示例

关于`synchronized`缺陷的描述, 我的理解是

1. `synchronized`没办法实现类似读写锁的机制, ta会把目标对象所有的访问都加上锁.
2. 当一个线程获得了`synchronized`锁, 其他线程尝试获取锁时会一直等待, 调度到这些等待着的CPU相当于浪费了. 
	- 而`Lock`机制中, 其他线程尝试获取锁时发现锁被占用, 会先在一个队列里进行排队(应该...这是我在 golang 里看的), 等到`Lock`被释放, 会通知这些线程重新竞争锁.

> 可以说, `synchronized`是隐式锁, `Lock`为显式锁.

`Lock`提供了比`synchronized`更多的功能. 但是要注意以下几点: 

1. `Lock`不是Java语言内置的, `synchronized`是Java语言的关键字, 因此是内置特性. `Lock`是一个类, 通过这个类可以实现同步访问; 
2. `Lock`和`synchronized`有一点非常大的不同, 采用`synchronized`不需要用户去手动释放锁, 当`synchronized`方法或者`synchronized`代码块执行完之后, 系统会自动让线程释放对锁的占用; 而Lock则必须要用户去手动释放锁, 如果没有主动释放锁, 就有可能导致出现死锁现象. 

> 如果不需要高级特性的话, `synchronized`就行了, 性能没啥大的差别.

------

关于 Java 中`Lock`的使用, 其实和其他高级语言中的一样了.

## 1. Lock

1. `lock()`: 使用最多的方法, 尝试获得锁, 没什么特殊的; 
2. `unlock()`: 不解释
3. `tryLock()`: 尝试获取锁, 如果获取失败则返回`false`, 不会阻塞, 类似于 python `asyncio.Queue()`队列的`get_nowait()`
4. `lockInterruptibly()`: 不常用, 通过此方法获取锁时, 可以通过`threadXXX.interrupt()`触发处于等待过程中的异常, 中断XXX线程的等待过程. 应该是用于对线程的精确控制.

> `lock`的获取与释放需要使用`try..catch..finally`中进行, 确保锁一定被释放.

...fuck, `Lock`只是个接口, 平常使用时还是直接用`ReentrantLock`或者`ReadWriteLock`吧, 真是长见识了.

```java
package com.example.demo;

import java.time.LocalDateTime;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class DemoApplication {
	private static int counter = 0;
	// 重入锁(Lock本身只是一个接口, 无法直接实例化)
	private static Lock counterLock = new ReentrantLock();
	private static void worker(){
		counterLock.lock();
		try {
			counter ++;
			// 就是其他语言中的 sleep, 单位为 ms
		} catch (Exception e){

		} finally{
			counterLock.unlock();
		}
		try {
			// 注意: 由于加了锁, 就不要加在 sleep 上了, 否则...
			Thread.sleep(5000);
		} catch(InterruptedException e){}
	}

	public static void main(String[] args) {
		LocalDateTime start = LocalDateTime.now();
		System.out.println("start: " + start);
		for(int i = 0; i < 2000; i ++){
			new Thread(){
				public void run(){
					// 将真正的工作函数传入
					worker();
				}
			}.start();
		}
		try {
			// 等待所有线程结束
			Thread.sleep(10000);
	
			LocalDateTime end = LocalDateTime.now();
			System.out.println("end: " + end);
			System.out.println("counter: " + counter);

		} catch(InterruptedException e){}
	}
}
```

## 2. ReentrantLock 可重入锁

`ReentrantLock`: re entrant lock

## 3. ReadWriteLock 读写锁
