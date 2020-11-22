# Thread类多线程.2.synchronized同步

参考文章

1. [Java并发编程：synchronized](https://www.cnblogs.com/dolphin0520/p/3923737.html)
	- Java 提供了两种方式来实现同步互斥访问: `synchronized`和`Lock`
	- `synchronized`关键字修饰的方法
	- `synObject`可以是`this`, 代表获取当前对象的锁, 也可以是类中的一个属性, 代表获取该属性的锁.

Java 提供了两种方式来实现同步互斥访问: `synchronized`和`Lock`, 虽然在其他语言中接触的`Lock`操作比较多. 但是在 Java 中, 貌似 `synchronized`先出现.

## 1. `synchronized`关键字修饰的方法

这个就不举例了, 其他高级语言里没见过这么用的, 对整个函数加锁, 感觉锁的粒度会很大, 不如对目标变量加锁灵活.

## 2. `synchronized`代码块

这种方法其实就和 lock 一样, 需要将一个对象(一般是`Object`对象)视作锁, 在这个代码块里执行可能发生并发冲突的操作

```java
package com.example.demo;

import java.time.LocalDateTime;

public class DemoApplication {
	private static int counter = 0;
	// 用作锁
	private static Object counterLock = new Object();
	private static void worker(){
		try {
			synchronized(counterLock){
				counter ++;
			}
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

参考文章1中给出了相关示例, 且有这么一句: **`synObject`可以是`this`, 代表获取当前对象的锁, 也可以是类中的一个属性, 代表获取该属性的锁.** 

这就意味着, 要执行`synchronized(this)`块中的代码, 线程需要拥有整个对象的锁, 其他任何线程都无法再对该对象进行任何修改(即使该块内只修改了对象的其中一个属性).

当然我在实际情况中很少遇到给整个对象加锁的场景, 多数是一个类中有多个可能发生线程冲突的属性, 为了减少锁粒度, 一般为每个属性创建一个独立的 lock 对象, 分别进行保护.

------

还有一点需要注意, 这是 Java 特有的场景.

~~类属性中`static`类型的成员属性与非`static`类型的成员属性~~, 这里ta用来举例的是`synchronized`方法而非成员属性. 

在多线程操作中, 如果两个线程其中一个执行的是`synchronized`修饰的普通方法, 而另一个执行的是`synchronized`修饰的`static`方法, 那么这两个线程将不会发生互斥, 也就不会发生阻塞, 等待的情况.

因为访问`static synchronized`方法占用的是类锁, 而访问非`static synchronized`方法占用的是对象锁, 所以不存在互斥现象.

对此我的理解是, 虽然`synchronized`方法没有显式地创建并使用锁对象, 但底层仍然是使用了锁的. 而且不同于`synchronized`代码块中用到的属性锁, 非`static synchronized`方法使用的应该就是上面我们提到的`synchronized(this)`, 即对整个对象加锁.

而`static synchronized`方法使用了类锁, 与对象锁的概念应该不在同一层. 正如`static`属性与普通属性也不在同一层一样.

这个以后遇到了再说吧.
