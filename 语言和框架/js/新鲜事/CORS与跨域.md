# CORS与跨域

参考文章

1. [ajax跨域，这应该是最全的解决方案了](https://segmentfault.com/a/1190000012469713#articleHeader2)

2. [AJAX POST&跨域 解决方案 - CORS](http://www.cnblogs.com/Darren_code/p/cors.html)

3. [前端AJAX请求跨域时遇到的一些坑](https://icewing.cc/post/about-cross-origin.html)

4. [Nginx 跨域设置 Access-Control-Allow-Origin 无效的解决办法](http://blog.csdn.net/frank_passion/article/details/53898769)

5. [Node+Express的跨域访问控制问题：Access-Control-Allow-Origin](http://blog.csdn.net/shelly1072/article/details/55003178)

认知:

1. 跨域限制一般是针对js脚本内的ajax请求的限制, 是客户端作的限制;

2. jsonp和cors都是解决跨域问题的方法, 前者的本质是妥协, 用普通http请求拿到加工过的数据再处理得到实际想要的数据, 现在已经不推荐使用; 后者是一种标准, 表现为在基础http协议上添加额外的头信息与服务器交互并互相确认.

参考文章1中有如下解释

> CORS是一个W3C标准, 全称是"跨域资源共享"（Cross-origin resource sharing）. 它允许浏览器向跨源服务器, 发出XMLHttpRequest请求, 从而克服了AJAX只能同源使用的限制. 

> 基本上目前所有的浏览器都实现了CORS标准, 其实目前几乎所有的浏览器ajax请求都是基于CORS机制的, 只不过可能平时前端开发人员并不关心而已, 所以说其实现在CORS解决方案 **主要是考虑后台该如何实现的问题**.

如何实现? 

一个典型场景: 客户端发出了一个跨域GET请求, 浏览器会在这之前先向目标域发出一个`OPTIONS`请求(与客户端程序无关, 是浏览器自发的行为), 而服务器端要在此接口确认, 能够接受哪个来源(`Origin`)的哪种请求(`Method`).

![](https://gitimg.generals.space/3c904fad63a2982bdb0d6c753d8db9ef.png)

![](https://gitimg.generals.space/243207c202cdb020926ca662fe6961b9.png)

上述两张图片显示了在本地创建的html请求第三方接口时的两次请求, 由于接口是nodejs的express写的, 所以这是在程序层面做的处理.

```js
// app.all('*', function(req, res, next){
app.all('/napi/bootstrap-table-addrbar', function(req, res, next){
    // 允许跨域请求
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin");
    res.header("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS");
    res.header("Content-Type", "application/json;charset=utf-8");
    if (req.method == 'OPTIONS')  {
        res.status(200).send('{"test": "options ok"}');
        return;
    }
    next();
});
```

这其实是参考文章5中给出的解决方法, 只是它只有代码, 没有解释. 

本来浏览器发出的`OPTIONS`请求, 只要服务端在响应头中加上上述几个字段即可, 当然, 可以根据实际情况对`Access-Control-Allow-Origin`做一些调整, 可以写多个值, 只要其中包含本次请求头的`Origin`的值, 浏览器就会再次发出真正的请求. 另外, `Content-Type`字段的值是`application/json`还是`text/plain`应该也需要做相应调整, 这个没有深究.

下面是nginx配置中的解决办法

```ini
if ($request_method = 'OPTIONS') {
    add_header 'Access-Control-Allow-Origin' '*';
    add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS';
    add_header 'Access-Control-Allow-Headers' 'DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type';
    return 204;
}
```

~~ 这是参考文章4中提到的, 可能没必要将各种请求都这个区分开, 不过是最全的情况了. 另外, ~~ 这个条件判断可以精确到接口级别(写到`location`块里), 没必要作为全局配置处理.

有了这些知识储备, 再遇到跨域问题, 可以参考文章1和参考文章3, 讲的很详细, 对各种问题的分析也到全面.

这就是`CROS`协议的做法了.