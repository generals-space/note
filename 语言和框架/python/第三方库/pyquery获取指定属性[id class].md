# pyquery获取指定属性[id class]

参考文章

1. [通过pyquery获取具有特定属性的元素](https://www.cnpython.com/qa/834847)
2. [pyquery基础](https://www.jianshu.com/p/5def029dbdf8)

当一个元素对象拥有的不是像`id`, `class`这种典型的属性, 无法通过`#xxx`和`.xxx`进行选择时, 而是拥有像`data-name=general`这种自定义的属性时, 使用pyquery如何选取目标元素?

```py
dom = PyQuery(html_text)
items = dom.find('div[data-name]')
```
