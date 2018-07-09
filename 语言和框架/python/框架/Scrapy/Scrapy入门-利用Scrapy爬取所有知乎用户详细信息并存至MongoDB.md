# Scrapy入门-利用Scrapy爬取所有知乎用户详细信息并存至MongoDB

参考文章

1. [利用Scrapy爬取所有知乎用户详细信息并存至MongoDB](https://www.cnblogs.com/qcloud1001/p/6744070.html)

2. [Scrapy官方文档 - Items](http://scrapy-chs.readthedocs.io/zh_CN/0.24/topics/items.html)

3. [Scrapy官方文档 - Item Pipeline](http://scrapy-chs.readthedocs.io/zh_CN/0.24/topics/item-pipeline.html)

参考文章1给出的实例对`item`, `item pipeline`的使用有着非常形象的介绍, 十分容易上手.(话说所谓去重就是在数据库中判断目标记录是否已存在, 有点low啊)