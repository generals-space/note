# 跟着例子一步步学习redux+react-redux

原文链接

[跟着例子一步步学习redux+react-redux](https://segmentfault.com/a/1190000012976767)

正如原文中所说, 没有使用`react`, `redux`中的术语去解释, 因为那只会引出更多的术语. 它以react为基础, 探索父组件与多级子组件间数据传递的最佳实践, 然后引用`redux`架构在`react`的实现方法, 值得一看.

`initState`

`reducer`: 根据传入的`action`(包括操作类型和操作数), 对state做相应的修改.

`createStore`: 将`initState`和`reducer`绑定起来, 然后就能得到`store`对象.
