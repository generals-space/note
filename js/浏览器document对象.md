# 浏览器document对象

参考文章

1. [HTML DOM Document 对象](http://www.w3school.com.cn/jsref/dom_obj_document.asp)

`document.URL`: 可以得到访问当前网站的完整url, 包括请求参数(`?`与`&`)和锚点信息(`#`). 直接设置这个值不会导致网页跳转.

`document.referrer`: 跳转来源(如果直接在浏览器地址栏输入网址, 则此值为空字符串).

`document.cookie`: 设置或返回与当前文档有关的所有cookie.

## 1. location对象

> chrome下, `document.location`与`window.location`都是合法的.

假设访问如下链接`http://www.redisfans.com/?p=68#1夜火`, 输出document的location对象, 其属性如下.


```
hash: "#1夜火"
host: "www.redisfans.com"
hostname: "www.redisfans.com"
href: "http://www.redisfans.com/?p=68#1夜火"
origin: "http://www.redisfans.com"
pathname: "/"
port: ""
protocol: "http:"
```

`location.href`: 可以得到浏览器地址栏网址信息, 直接设置这个值会导致浏览器跳转行为.