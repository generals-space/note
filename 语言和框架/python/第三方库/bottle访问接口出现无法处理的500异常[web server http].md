# bottle访问接口出现无法处理的500异常[web server http]

参考文章

1. [Unhandled exception (http error 500) when using Bottle with AJAX](https://stackoverflow.com/questions/34498162/unhandled-exception-http-error-500-when-using-bottle-with-ajax)

## 问题描述

用bottle写了一个接口, 但是访问时总是出现500问题, 我已经尝试过在整个接口函数体中进行`try..except..`进行异常捕获了, 但竟然没有捕获成功...

## 解决方法

按照参考文章1中所说, bottle要求接口返回的数据必须是`str`类型, 我是直接返回了数值, 所以才会报这个异常, 将返回值转换成`str`类型, 就没问题了.
