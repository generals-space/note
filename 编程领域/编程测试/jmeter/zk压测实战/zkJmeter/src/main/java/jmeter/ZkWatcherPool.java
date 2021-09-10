package jmeter;

public class ZkWatcherPool {
    public static int watchTime = 1000*600;

    public static void main(String[] args) {
        for(int i = 0; i < 5; i ++){
            final int j = i;
			new Thread(){
				public void run(){
                    ZkWatcher zkWatcher = new ZkWatcher();
					// 将真正的工作函数传入
                    zkWatcher.createWatcher(j);
				}
			}.start();
		}

        try {
            System.out.println("threads create completed");
			// 等待所有线程结束
			Thread.sleep(watchTime+1000*30);
		} catch(InterruptedException e){
		}
    }
}