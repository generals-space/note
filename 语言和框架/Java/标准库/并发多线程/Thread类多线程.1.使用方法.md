参考文章

1. [Java并发编程：Thread类的使用](https://www.cnblogs.com/dolphin0520/p/3920357.html)
    - sleep相当于让线程睡眠, 交出CPU, 让CPU去执行其他的任务 - 这句应该是错误的, sleep 不会让出 cpu.
    - `yeild`是主动让出 cpu, 得让线程预判啊. 比如线程将要进行一个 IO 操作, 应该是发起 IO 请求后再 yield, 等待系统回调, 但这不好搞啊.
    - 线程的状态可以看一下.
2. [Java并发编程：synchronized](https://www.cnblogs.com/dolphin0520/p/3923737.html)
	- 创建多线程另一种方法, 不需要创建 Thread 的扩展类
3. [Java锁机制Lock用法示例](https://www.jb51.net/article/146258.htm)
	- `Runnable`接口多线程示例

> `Thread.currentThread().sleep(2000)`与`Thread.sleep(2000)`区别是啥???

与 python 的 `threading`库一样, java 的 Thread 类在使用方法, 参数设置等操作上都大致相同.

不过 java 的线程同步比较...别致, 不像其他高级语言那样显式加锁, 而是用一个`synchronized`块.

Java Thread 与 python 一样, 也有两种使用方法

## 1. 继承 Thread 类

```java
package com.example.demo;

import java.time.LocalDateTime;

public class DemoApplication {
	private static int counter = 0;
	// 用作锁, 不过没加锁 countere 的累加结果也没错.
	private Object lock = new Object();

	class MyWorker extends Thread {
		// 扩展 Thread 类必须重写 run 方法
		public void run(){
			try{
				counter ++;
				Thread.sleep(5000);
			} catch(InterruptedException e){

			}
		}
	}
	public static void main(String[] args) {
		DemoApplication app = new DemoApplication();

		LocalDateTime start = LocalDateTime.now();
		System.out.println("start: " + start);
		for(int i = 0; i < 2000; i ++){
			MyWorker myworker = app.new MyWorker();
			myworker.start();
		}

		try {
            // 等待所有线程结束
            // 笨方法, 本来 thread 对象有 join 方法可以等待其结束, 但是这么多线程不容易搞.
            // 这种场景也不合适.
			Thread.sleep(10000);
	
			LocalDateTime end = LocalDateTime.now();
			System.out.println("end: " + end);
			System.out.println("counter: " + counter);

		} catch(InterruptedException e){

		}
	}
}
```

## 2. 创建 Thread 对象并传入工作函数

```java
package com.example.demo;

import java.time.LocalDateTime;

public class DemoApplication {
	private static int counter = 0;
	// 用作锁
	private Object lock = new Object();
	private static void worker(){
		try {
			counter ++;
			// 就是其他语言中的 sleep, 单位为 ms
			Thread.sleep(5000);
		} catch (InterruptedException e){

		}
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

		} catch(InterruptedException e){

		}
	}
}
```

这里初始化 new Thread 实例, 并在`run()`方法中传入真正的工作函数的方式, 与 python 的 threading 有异曲同工之意.

不过其实实现了`run()`方法的, 其实就相当于实现了`Runaable`接口了.

## 实现 Runaable 接口的worker
