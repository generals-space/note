参考文章

1. [golang中context包解读](http://www.01happy.com/golang-context-reading/)

其实我们完全可以使用select+chan完全协程的流程控制, 不需要借助context. 但是既然有了这样一个包, 并且使用起来也蛮方便, 就来看一看ta是如何做的.

两个接口

1. Context: Deadline, Done, Err, Value方法
2. canceler: cancel, Done方法

ctx对象:

1. emptyCtx
2. cancelCtx
3. timerCtx
4. valueCtx

只有`cancelCtx`有new方法, 其余都是嵌入cancelCtx以达到可cancel()的目的.

4种ctx对象的组合方式很简单, 稍微费神一点的就是`propagateCancel`和`cancel`的执行流程. ta们之间需要建立/取消父子关系.
