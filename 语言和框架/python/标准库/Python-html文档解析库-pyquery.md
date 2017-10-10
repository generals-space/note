# Python-html文档解析库-pyquery

参考文章

1. [官方文档索引](http://pythonhosted.org/pyquery/genindex.html)

2. [Python中的jquery PyQuery库使用小结](http://www.jb51.net/article/50069.htm)

3. [python爬虫神器PyQuery的使用方法](https://segmentfault.com/a/1190000005182997#articleHeader8)

首先, 与jQuery不同的是, pyquery需要一个初始化的过程, 毕竟pyquery不像jquery是从远程下载下来的时候就内嵌在网页中自执行的.

```py
from pyquery import PyQuery as pyQuery

result = urllib2.urlopen(某网址).read()
pyQuery = pyQuery(result) ## 实例化pyQuery
for item in pyQuery('tr').items():
    ele = item.find('td').eq(0).text()
```

pyQuery库大部分API函数都可以返回实例本身, 即可以实现链式调用, 写起来很舒服.

但它的选择器不多, 只能通过标签, id和class三种方式选择.

直接选择: 如通过`py = pyQuery(文档内容)`后, 可以通过`py('div')`, `py('#id名称')`或`py(.class名称)`

或是通过`find()`如`find('div')`, `find('#id')`或`find('.class')`这种.

直接实例多为pyquery对象, 可以通过`item()`得到选择后的列表对象, 或是通过`eq()`方法指定索引
