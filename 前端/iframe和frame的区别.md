# iframe和frame的区别

参考文章

1. [<iframe>和<frame>区别](https://www.cnblogs.com/ahudyan-forever/p/5706873.html)

2. [html/css基础篇——iframe和frame的区别【转】](https://www.cnblogs.com/chuaWeb/p/5124368.html)

3. [iframe下内容自适应缩放](https://segmentfault.com/q/1010000000259124)

> html5不再支持frame, 单这一条就不需要再纠结要使用哪一个了.

iframe的宽高貌似不能用`class`, `id`定义, 只能在元素本身上通过`width`与`height`属性定义.

-----

很多时候, 内嵌的`iframe`页面大小不符合我们的要求, 但控制权一般不在我们这边, 我们希望能控制这些子页面的缩放, 但是是很难(css的`zoom`字段无效, width的百分比只能改变内嵌页可视窗口的大小). 

参考文章3提供了一个思路, 通过css3的`transform`的`scale`函数. 这种方法有一个难点是, `scale`是相对于其父容器的缩放. 所以`iframe`如果本身已经铺满父容器(`width: 100%; height: 100%`), 再缩小或放大它的话, 你肯定无法得到一个合适的页面. 所以, 需要在定义`width`与`height`属性时, **反向**定义. 

例如: 如果你有一个`100*100`的容器, 但是内嵌页面为`200*200`, 我们想要把它缩小到`100*100`的父容器内, 那么定义这个`iframe`的`width`与`height`为`200%`, 但是`scale(0.5, 0.5)`. 总之就是, 先让`iframe`能够完全承载内嵌页面的显示, 相对于父容器扩大, 那么其内容就能全部展现, 然后`scale`缩小, 就可以了. 在这期间要掌握到缩放的比例.

参考文章3中还提到了`transform-origin`这个参数, 不过貌似是针对`rotate(旋转)`变换的, `scale`时, 如果`iframe`的`width`与`height`不是等比例变化时, 可能不是那么准.