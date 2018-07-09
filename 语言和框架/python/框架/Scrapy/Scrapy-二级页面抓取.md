# Scrapy-二级页面抓取

参考文章

1. [Scrapy 学习笔记 -- 解决分页爬取的问题](https://www.jianshu.com/p/0c957c57ae10)

2. [Scrapy官方教程 - Following links](https://doc.scrapy.org/en/latest/intro/tutorial.html#following-links)

网站页面典型结构就是**n页列表页** -> **列表链接内容页(内容也分多级页面)**.

参考文章1对于如何使用`yeild`递归处理二级页面解释的比较清楚, 代码也足够清晰.

> 列表页面分析器处理两件事: 一件是分析页面, 拿数据的链接, 交给`self.parse_content()`处理内容页, 另一个就是拿到**下一页**的地址, 由于和当前页面结构一样的, 只需要交由本身再进行分析, 处理即可.

参考文章2是scapy的官方文档, 也有易于理解的代码示例.