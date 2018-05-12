# golang-json解析

参考文章

1. [go json解析Marshal和Unmarshal](http://www.baiyuxiong.com/?p=923&utm_source=tuicool&utm_medium=referral)

2. [JSON与Go](http://rgyq.blog.163.com/blog/static/316125382013934153244/)

json字符串的解析结果一般对应语言中的`dict(python)`, `map(java)`, `object(js)`等相似结构, 在go中为结构体struct.

在go中, json的序列化与反序列化主要涉及结构体与`[]byte`类型的相互转换.

参考文章1中提供了简单易读的示例可以看一下.

> 注意: 结构体的成员首字母大写, 不然`Marshal`得不到其中的成员.