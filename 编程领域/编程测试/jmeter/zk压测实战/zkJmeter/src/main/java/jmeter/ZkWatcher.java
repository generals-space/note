package jmeter;

import java.util.List;
import java.util.concurrent.CountDownLatch;

import org.apache.zookeeper.CreateMode;
import org.apache.zookeeper.WatchedEvent;
import org.apache.zookeeper.Watcher;
import org.apache.zookeeper.Watcher.Event.EventType;
import org.apache.zookeeper.Watcher.Event.KeeperState;
import org.apache.zookeeper.ZooDefs.Ids;
import org.apache.zookeeper.ZooKeeper;

public class ZkWatcher {
    //存储用户输入的zk地址
    private static String conn_addr = "172.22.254.59:2181,172.22.254.64:2181,172.22.254.119:2181";
    // session超时时间, 单位ms
    private static String session_timeout = "30000";
    //zk父节点
    private static String parent_node = "/test16";
    private static String child_node = "/child";
    private static String node_content = "hello world!";

    // 信号量, 阻塞程序执行, 用于等待zookeeper连接成功, 发送成功信号(类似于python的Event, golang的WaitGroup)
    final CountDownLatch connectedSemaphore = new CountDownLatch(1);

    // zk连接对象
    private ZooKeeper zk;

    public void createWatcher(int threadNum){
        try {
            zk = new ZooKeeper(conn_addr, Integer.valueOf(session_timeout), new Watcher() {
                public void process(WatchedEvent event) {
                    // 获取事件的状态
                    KeeperState keeperState = event.getState();
                    EventType eventType = event.getType();
                    // 如果是建立连接
                    if (KeeperState.SyncConnected == keeperState) {
                        if (EventType.None == eventType) {
                            // 如果建立连接成功, 则发送信号量, 让后续阻塞程序向下执行
                            // System.out.println("zk 建立连接");
                            connectedSemaphore.countDown();
                            return;
                        } else if(EventType.NodeChildrenChanged == eventType) {
                            System.out.println("NodeChildrenChanged: " + event.getPath());
                        } else if(EventType.NodeDataChanged == eventType) {
                            System.out.println("NodeDataChanged: " + event.getPath());
                        }
                        try {
                            // 开启 watch, 为什么 getData() 不行? 非得是 getChildren() ???
                            // zk.getData(parent_node, true, null);
                            zk.getChildren(parent_node, true, null);
                        } catch (Exception e){
                            e.printStackTrace();
                        }
                    }
                }
            });

            connectedSemaphore.await();

            // 如果存在父节点则不创建父节点
            if (null == zk.exists(parent_node, false)) {
                // 创建父节点
                zk.create(parent_node, parent_node.getBytes(), Ids.OPEN_ACL_UNSAFE, CreateMode.PERSISTENT);
            }
            // 监听子节点
            // 开启 watch, 为什么 getData() 不行? 非得是 getChildren() ???
            // byte[] data = zk.getData(parent_node, true, null);
            // System.out.println(new String(data));
            List<String> list = zk.getChildren(parent_node, true);
            // for(int i = 0; i < list.size(); i ++){
            //     System.out.print(list.get(i) + " ");
            // }
            // System.out.println();
            Thread.sleep(ZkWatcherPool.watchTime);
            System.out.println(threadNum + ": exit 0");
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}
