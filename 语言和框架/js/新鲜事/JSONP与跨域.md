# JSONP与跨域

参考文章

1. [AJAX 跨域请求 - JSONP获取JSON数据](http://justcoding.iteye.com/blog/1366102)

2. [轻松搞定JSONP跨域请求 - CSDN](http://blog.csdn.net/u014607184/article/details/52027879)

3. [使用jsonp服务器端需要做什么？ - 知乎](https://www.zhihu.com/question/57058052/answer/151670760)

## 1. 认知准备

1. 跨域限制是由浏览器实现的机制, 目的是为了保护**客户端用户**的安全(而不是服务端)

2. 跨域只限制`xhr`操作, 即ajax请求, 而对`script`标签的`src`, `link`标签的`href`和`img`标签的`src`则没有限制, JSONP就是模拟了`script`标签的`src`代为发送get请求.

3. JSONP的实现需要服务端配合.

## 2. 实现原理

原生js

```html
<script type="text/javascript">  
    function resHandler(result) {  
        console.log(result);
    }  
    var JSONP=document.createElement("script");  
    JSONP.type="text/javascript";  
    JSONP.src="http://crossdomain.com/services.php?callback=resHandler";  
    document.getElementsByTagName("head")[0].appendChild(JSONP);  
</script> 
```

或

```html
<script type="text/javascript">  
    function resHandler(result) {  
        console.log(result);
    }
</script>  
<script type="text/javascript" src="http://crossdomain.com/services.php?callback=resHandler"></script>  
```

可以看到, 实际的请求操作是`script`的`src`发出的, 它请求了`http://crossdomain.com/services.php`这个接口.

```php
<?php  
$arr = array(
    'a'=>1,
    'b'=>2,
    'c'=>3,
    'd'=>4,
    'e'=>5
);

## result是一个常规的json字符串. 如果是同域请求的, 就可以直接返回了.
$result = json_encode($arr);

//动态执行回调函数
$callback=$_GET['callback'];
## 这里即是`script`脚本中将要展现的内容, 为`resHandler($result)`
echo $callback."($result)";
```

我们在前端先定义了一个(回调)函数`resHandler`, 再把这个函数名当作参数一同请求到目标域接口. 后端得到了我们想要的结果后不是直接返回, 而是拼接回调函数名与`$result`字符串, 得到`resHandler($result)`, 前端`script`标签得到这个这个字符串, 会开始执行, 相当于调用了`resHandler`, 参数是我们想要的结果$result. 其中`callback`这个参数是前后端协商确定的.

非常巧妙的设计.

如果单纯请求一个文件, 不经过后台处理, 则其格式可为

```js
resHandler(
    {
        id: 1,
        name: 'general'
    }
)
```

则前端可以得到`{id:1,name:'general'}`这个对象.

这样的优点是方便, 不经过后台, 但是前端就无法通过`jsonpCallback`这个参数自定义处理函数了, 可能会有些不够灵活.

## 3. jQuery的实现

jquery为我们封装了这个过程. 一共有3种实现方式.

1. `$.ajax`

2. `$.getJSON`

3. `$.get`

### 3.1 `$.ajax`

```js
$.ajax({  
    url: "http://crossdomain.com/services.php",  
    dataType: 'jsonp',
    data: '',
    jsonp: 'callback',               // get请求时的键名
    success: function(result) {
        console.log(result);
    },
});
```

默认回调函数即是`success`的回调. 也可以通过`jsonpCallback`选项额外指定回调函数名(必须是字符串形式), 貌似`success`回调也能直接接受`result`参数...不受`jsonpCallback`的影响.

如下

```js
function resHandler(result) {
    console.log(result);
}
$.ajax({  
    url: "http://crossdomain.com/services.php",  
    dataType: 'jsonp',
    data: '',  
    jsonp: 'callback',               // get请求时的键名
    jsonpCallback: resHandler
});
```

在`$.ajax`中加入`dataType: 'jsonp'`, 表明这是jsonp请求(虽然是封装在了`$.ajax`, 但jsonp类型的请求与普通xhr请求的处理代码是分开的).

### 3.2 `$.getJSON`

`$.getJSON(url,[data],[callback])`

> 在 jQuery 1.2 中，您可以通过使用JSONP形式的回调函数来加载其他网域的JSON数据，如 "myurl?callback=?"。jQuery 将自动替换 ? 为正确的函数名，以执行回调函数。 注意：此行以后的代码将在这个回调函数执行前执行。

```js
$.getJSON("http://crossdomain.com/services.php?callback=?", function(result){
    console.log(result);
});
```

### 3.3 `$.get`

```js
$.get('http://crossdomain.com/services.php?callback=?', 
    function (result){
        console.log(result);
    },
    'jsonp'
);  
```