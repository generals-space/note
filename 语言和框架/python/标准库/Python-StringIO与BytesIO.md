# Python-StringIO与BytesIO

参考文章

1. [skeyes/proxy-pool](https://gitee.com/skeyes/proxy-pool/blob/4a5dad66f299d433b58aadd48eb19afc27e5b55d/pspider/docs/%E5%9C%A8%E7%BA%BF%E5%9B%BE%E7%89%87%E8%AF%86%E5%88%AB.md)
    - 自己的小项目
2. [Python3 IO](https://www.cnblogs.com/284628487a/p/5590692.html)
3. [StringIO和BytesIO](https://www.liaoxuefeng.com/wiki/0014316089557264a6b348958f449949df42a6d3a2e542c000/001431918785710e86a1a120ce04925bae155012c7fc71e000)
    - 廖雪峰老师的博客

在`python2`中, `StringIO`与`io`同属于标准库, `BytesIO`则在`io`库中; 

在`python3`中, `StringIO`与`BytesIO`都被划分到`io`标准库中了;

`StringIO`与`BytesIO`就类似 golang 中的`Buffer`一样, 使用场景的话...参考文章1中, `PIL.Image()`需要使用这样的 Buffer 对象.

其他的我还真没见过.
