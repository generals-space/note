# Executors.newCachedThreadPool传参[线程池]

我们可以在将 worker 方法放入线程池内执行前将要传入的参数传入.

```java
// 省略...
public class DemoApplication {
	private static void worker(String workerName, String[] argList){
		try {
			System.out.println(workerName);
			if(workerName.equals("worker5")){
				argList[2] = "ls /tmp";
			}
			for (String arg : argList) {
				System.out.println(arg);
			}
			// 就是其他语言中的 sleep, 单位为 ms
			Thread.sleep(5000);
		} catch (InterruptedException e){

		}
	}

	public static void main(String[] args) {
		LocalDateTime start = LocalDateTime.now();
		System.out.println("start: " + start);

		ExecutorService pool = Executors.newCachedThreadPool();
		String[] argList = {
			"/bin/bash",
			"-c",
			"echo yes"
		};
		String worker = "worker";
		for(int i = 0; i < 10; i ++){
			String workerName = worker + i;
			// runnable 行内定义函数的方式, 类似于 python 中的 lamda 表达式
			Runnable runnable = () -> worker(workerName, argList);
			pool.execute(runnable);
		}
		// shutdown() 关闭通道, 之后无法再往里添加任务, 并非关闭/回收线程池的意思...
		// 类似的话, 应该像是其他语言中的 pool.start() 
		pool.shutdown();

		// 省略...
	}
}

```

但是要注意, 传入各线程的变量是共享的, 如果是引用类型, 当一个线程修改其中的值后, 其余线程中该变量的值也会发生变化, 从而导致不可控的结果.

为了避免这种情况最好在传入 worker 方法时使用`argList.clone()`拷贝一份副本再传. 至于其余引用类型, 就要找各自的深/浅拷贝方法了.

