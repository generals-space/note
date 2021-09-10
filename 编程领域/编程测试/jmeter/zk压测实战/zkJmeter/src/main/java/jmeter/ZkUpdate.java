package jmeter;

import java.io.Serializable;
import java.util.concurrent.CountDownLatch;

import org.apache.jmeter.config.Arguments;
import org.apache.jmeter.protocol.java.sampler.AbstractJavaSamplerClient;
import org.apache.jmeter.protocol.java.sampler.JavaSamplerContext;
import org.apache.jmeter.samplers.SampleResult;

import org.apache.zookeeper.CreateMode;
import org.apache.zookeeper.WatchedEvent;
import org.apache.zookeeper.Watcher;
import org.apache.zookeeper.Watcher.Event.EventType;
import org.apache.zookeeper.Watcher.Event.KeeperState;
import org.apache.zookeeper.ZooDefs.Ids;
import org.apache.zookeeper.ZooKeeper.States;
import org.apache.zookeeper.ZooKeeper;

/*
 * public Arguments getDefaultParameters(): 设置可用参数及的默认值;
 * public void setupTest(JavaSamplerContext arg0): 每个线程测试前执行一次, 做一些初始化工作;
 * public SampleResult runTest(JavaSamplerContext arg0): 开始测试, 从arg0参数可以获得参数值;
 * public void teardownTest(JavaSamplerContext arg0): 测试结束时调用;
 */


/*
 * parent_node: /test{1..100}, 每次实验可设置不同值, 避免因为节点已存在而无法创建成功, 删起来也方便.
 * child_node: /${__counter(False, counter)} 注意前面的斜线!!!
 *              False表示counter是所有任务中的全局计数,
 *              如果为True则表示在单个线程中中的全局计数.
 *              counter...还没弄明白, 貌似是函数将计数值赋值给counter变量, Jmeter 可以通过counter得到计数值.
 */
public class ZkUpdate extends AbstractJavaSamplerClient implements Serializable {
    //存储用户输入的zk地址
    private static String conn_addr;
    //设置GUI页面显示的变量名称
    private static final String ConnAddrName="conn_addr";
    //设置GUI页面默认显示的变量值,默认值为空
    private static final String ConnAddrValueDefault="127.0.0.1:2181";

    // session超时时间, 单位ms
    private static String session_timeout;
    //设置GUI页面显示的变量名称
    private static final String SessTimeName="session_timeout";
    //设置GUI页面默认显示的变量值,默认值为空
    private static final String SessTimeValueDefault="5000";

    //zk父节点
    private static String parent_node;
    //设置GUI页面显示的变量名称
    private static final String ParentNodeName="parent_node";
    //设置GUI页面默认显示的变量值,默认值为空
    private static final String ParentNodeNameDefault="/test";

    private static String child_node;
    private static String node_content;

    // 信号量, 阻塞程序执行, 用于等待zookeeper连接成功, 发送成功信号(类似于python的Event, golang的WaitGroup)
    static final CountDownLatch connectedSemaphore = new CountDownLatch(1);

    // zk连接对象
    private ZooKeeper zk;

    // resultData变量用来存储响应的数据, 目的是显示到查看结果树中.
    private static String resultData;

    /*
     * 这个方法用来控制显示在GUI页面的属性, 由用户来进行设置.
     * 此方法不用调用, 是一个与生命周期相关的方法, 类加载则运行.
     * (non-Javadoc)
     * @see org.apache.jmeter.protocol.java.sampler.AbstractJavaSamplerClient#getDefaultParameters()
     */
    @Override
    public Arguments getDefaultParameters() {
        System.out.println("读取属性值");
        Arguments params = new Arguments();
        params.addArgument("conn_addr", String.valueOf(ConnAddrValueDefault));
        params.addArgument("session_timeout", String.valueOf(SessTimeValueDefault));
        params.addArgument("parent_node", String.valueOf(ParentNodeNameDefault));
        params.addArgument("child_node", "");
        params.addArgument("node_content", "");
        return params;
    }

    /**
     * 初始化方法, 初始化性能测试时的每个线程.
     * 实际运行时每个线程仅执行一次, 在测试方法运行前执行, 类似于LoadRunner中的init方法.
     */
    @Override
    public void setupTest(JavaSamplerContext jsc) {
        conn_addr = jsc.getParameter(ConnAddrName, ConnAddrValueDefault);
        session_timeout = jsc.getParameter(SessTimeName, SessTimeValueDefault);
        parent_node = jsc.getParameter(ParentNodeName, ParentNodeNameDefault);
        try {
            // 为每个测试线程都创建一个zk连接.
            zk = new ZooKeeper(conn_addr, Integer.valueOf(session_timeout),
                    new Watcher() {
                        public void process(WatchedEvent event) {
                            // 获取事件的状态
                            KeeperState keeperState = event.getState();
                            EventType eventType = event.getType();
                            // 如果是建立连接
                            if (KeeperState.SyncConnected == keeperState) {
                                if (EventType.None == eventType) {
                                    // 如果建立连接成功, 则发送信号量, 让后续阻塞程序向下执行
                                    System.out.println("zk 建立连接");
                                    connectedSemaphore.countDown();
                                }
                            }
                        }
                    });
            if (States.CONNECTING == zk.getState()) {
                connectedSemaphore.await();
            }
            // 如果存在父节点则不创建父节点
            if (null == zk.exists(parent_node, false)) {
                // 创建父节点
                zk.create(parent_node, parent_node.getBytes(), Ids.OPEN_ACL_UNSAFE, CreateMode.PERSISTENT);
            }
        } catch (Exception e) {
            // TODO Auto-generated catch block
            // e.printStackTrace();
            throw new RuntimeException(e);
        }
    }

    @Override
    public SampleResult runTest(JavaSamplerContext arg0) {
        String child_node = arg0.getParameter("child_node");
        String node_content = arg0.getParameter("node_content");
        System.out.println(child_node + ": " + node_content);

        // SampleResult这个类是用来将测试结果输出到查看结果树中的, 并且也是用来控制事务的开始和结束的.
        SampleResult results = new SampleResult();
        // 查看结果树中, 左侧Text栏会显示的任务列表(注意, 不是线程列表)中的名称.
        // 注意: 需要将此值设置为统一值, 因为在聚合分析时, 此值会作为标签统计汇总.
        // 比如, 如果这个值在每个任务中都为唯一值, 生成 100 个任务, 50个成功, 50个失败,
        // 在聚合分析时会出现100条结果, 50条的成功率是100%, 另外50条的成功率是0%;
        // 如果这个标签在每个任务中都相同, 得到的结果只有1条, 成功率50%;
        results.setSampleLabel("zk节点测试: " + parent_node);
        results.setDataType(SampleResult.TEXT);

        try{
            // 事务开始标记
            results.sampleStart();

            zk.setData(parent_node + child_node, node_content.getBytes(), -1);

            // 获取节点信息
            byte[] data = zk.getData(parent_node + child_node, false, null);
            resultData = new String(data);
            // setResponseData() 中写入的数据会在聚合报告中进行统计, 得到接收和发送的数据.
            if(null == resultData){
                results.setSuccessful(false);
                results.setResponseData("zk result is null", null);
            } else {
                results.setSuccessful(true);
                results.setResponseData(resultData, null);
                System.out.println("写入成功");
            }
        }catch(Exception e){
            results.setSuccessful(false);
            results.setResponseData(e.toString(), null);
            e.printStackTrace();
        }finally{
            //标记事务结束
            results.sampleEnd();
            return results;
        }
    }

    /**
     * 测试结束方法, 结束测试中的每个线程
     * 实际运行时, 每个线程仅执行一次, 在测试方法运行结束后执行, 类似于Loadrunner中的End方法
     */
    public void teardownTest(JavaSamplerContext arg0) {
        try {
            zk.close();
            System.out.println("关闭");
        } catch (InterruptedException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
    }

    // 测试代码
    public static void main(String[] args) {
        conn_addr = "172.22.254.179:2181";
        session_timeout = "2000";
        parent_node = "/parent";
        child_node = "/child";
        node_content = "hello world!";
        try {
            ZooKeeper zk = new ZooKeeper(conn_addr, Integer.valueOf(session_timeout),
                new Watcher() {
                    public void process(WatchedEvent event) {
                        // 获取事件的状态
                        KeeperState keeperState = event.getState();
                        EventType eventType = event.getType();
                        // 如果是建立连接
                        if (KeeperState.SyncConnected == keeperState) {
                            if (EventType.None == eventType) {
                                // 如果建立连接成功, 则发送信号量, 让后续阻塞程序向下执行
                                System.out.println("zk 建立连接");
                                connectedSemaphore.countDown();
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
            // 创建子节点
            zk.create(parent_node+child_node, node_content.getBytes(), Ids.OPEN_ACL_UNSAFE, CreateMode.PERSISTENT);
            // 获取节点信息
            byte[] data = zk.getData(parent_node + child_node, false, null);
            resultData = new String(data);
            System.out.println(new String(data));
            System.out.println(zk.getChildren(parent_node, false));
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}
