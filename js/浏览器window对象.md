# 浏览器window对象

参考文章

1. [Window 对象](http://www.w3school.com.cn/jsref/dom_obj_window.asp)

`window.name`: w3c的解释为'设置或返回窗口的名称'. 在chrome中, 设置这个值后在当前标签页生命周期中一直有效, 并且没有跨域问题, 页面刷新或跳转都存有这个值. 但是在其他标签中无法访问.