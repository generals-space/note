# Nginx访问规则详解

本文涉及到`location`的正则匹配, `if`条件判断及之后的url重写, `proxy_pass`前后发生的事情等的详细研究.

## 1. location与if的理解

`location`与`if`都可以看作是条件匹配然后执行某些操作的方式, 并且两者都可以通过正则匹配和字符串匹配. 只不过`location`匹配的只能是uri路径, 而if则可以匹配任何变量(不管是nginx内置变量还是自定义变量).

```
## if与括号()之间必须要有空格
if ($request_uri ^/main){
  set $abc 123;
}

location /main {
  set $abc 321;
}
```

------

我想两者的另一个最大的区别就是, `location`可以由nginx自身生成响应而`if`不行.

比如, `if`块中可以使用`rewrite`指令重定向让浏览器再发出一次http请求, 这一点与在`location`中相同; 但是, `if`语句块中没有办法通过`root`指令指定本地的响应目录以返回哪怕是一个静态资源, 而且它也不能通过`proxy_pass`指令为客户端代理一个请求.

所以, 如果`if`语句在`location`之前完成了匹配并通过`rewrite`指令的, 就不会再执行`location`中的操作了.

## 2. location匹配顺序

参考文章

[Nginx之location 匹配规则详解](http://www.cnblogs.com/lidabo/p/4169396.html)

> 当`server`块内同时存在正则形式与普通形式的`location`匹配规则时, 会优先匹配普通形式的location.

我们知道, `location`与`if`的匹配规则都有如下几种

- 没有任何符号 表示普通匹配, 可作为路径前缀匹配

- `=` 表示精确匹配, 其后面不能再接任何字符串

- `^~` 表示uri以某个常规字符串开头，不是正则匹配

- `~`|`!~` 表示区分大小写的正则匹配/不匹配

- `~*`|`!~*` 表示不区分大小写的正则匹配/不匹配

### 2.1 普通匹配

上面5种情况, 前3种都是普通匹配.

在第1种普通匹配中, 默认是匹配所有普通匹配项里最长最准确的那一个. 比如

```
location /main {

}
location /main/abc {

}
```

当用户访问的uri为`/main/abc/index.php`, 将匹配到第2条`location`. 这一点很容易理解.

另外`=`的级别高于`^=`, 也高于什么都不写的时候. 例如

```
location = /main/ {

}
location ^~ /main/ {

}

location = /abc_def/{

}
location ^~ /abc_{

}
```

当用户访问`/main/`时, 将会匹配到第1条; 当访问`/abc_def/`时将会匹配到第3条.

同理, 无符号最长路径uri也将高于`^~`指定的前缀匹配. 例如

```
location /main/abc{

}
location ^~ /main/ {

}
```

用户访问`/main/abc/index.html`时, 将匹配到第1条. 这与`location`放置的先后顺序无关.

```
location = / {

}
location / {

}
```

这样, 网站根路径可以是单独的处理规则, 其他路径则是另外一种规则.

### 2.2 正则匹配

正则匹配按照在配置文件的先后顺序进行, 一旦匹配成功, 就不再继续向下执行匹配了. 这一点与普通匹配不同, 所以, 有相同前缀的正则表达式, 还是把较长的项放在前面吧.

比如

```
location ~* ^/([^\/_]*)_([^\/_]*)/main/(.*)$ {
    proxy_pass http://backend/$3;
}   
location ~* ^/([^\/_]*)_([^\/_]*)/(.*)$ {
    proxy_pass http://frontend/$3;
}
```

以上正则匹配的是以类似于`/abc_def/`, `/lmn_xyz/`为前缀的uri, 如果是`/abc_def/main/index.html`将会把请求的`index.html`部分转发至名为`backend`的`upstream`池, 如果是`/abc_def/index.html`则会将`index.html`部分转发至名为`frontend`的`upstream`池.

注意: 如果将两个`location`块调换位置, 所有以类似于`/abc_def/`为前缀的uri都将被转发至`frontend`, 不管`/abc_def`后有没有`/main`路径. 这就是正则匹配的位置决定顺序的特性.

### 2.3 混合

如果配置文件中同时存在普通匹配与正则匹配, nginx会优先完成普通匹配, 然后进行正则匹配, 所以普通匹配的优先级是大于正则匹配的.

但是这样的话, 如果一个uri同时满足普通匹配与正则匹配, 将会执行后者的操作. 所以说, **优先级高不代表最终以其为准**.

那如果希望普通匹配与正则匹配同时满足的情况下执行前者指定的操作, 怎么办?

普通匹配中`=`与`^~`就可以达到这个目的, 当完成以这两个符号开始的普通匹配后, 将不再执行正则匹配. 但是它们并不阻止继续进行普通匹配, 也就是说仍然会按照最长最准确的`location`进行操作, 所以, 要达到终止正则匹配的目的, 需要在最长的`location`字符串前使用`^~`.

先看第1个例子, 依然与2.2节中的正则匹配有关.

```
location ~* ^/([^\/_]*)_([^\/_]*)/main/(.*)$ {
    proxy_pass http://backend/$3;
}   
location ~* ^/([^\/_]*)_([^\/_]*)/(.*)$ {
    proxy_pass http://frontend/$3;
}
location /abc_def/{

}
```

当用户访问`/abc_def/index.html`, nginx将先进行第3条`location`匹配, 之后执行第1, 2条, 于是最后匹配到第2条.

第2个例子.

```
location ~* ^/([^\/_]*)_([^\/_]*)/main/(.*)$ {
    proxy_pass http://backend/$3;
}   
location ~* ^/([^\/_]*)_([^\/_]*)/(.*)$ {
    proxy_pass http://frontend/$3;
}
location = /abc_def/ {

}
location ^= /abc_def/{

}
```

当用户访问`/abc_def/`, 将匹配到第3条, 由此也可以看出同样匹配字符串情况下, `=`的级别高于`^~`; 如果访问`/abc_def/index.html`, 将匹配到第4条.

无论哪一种, 都不会在进行正则方式的`location`匹配.

第3个例子.


```
location ~* ^/([^\/_]*)_([^\/_]*)/main/(.*)$ {
    proxy_pass http://backend/$3;
}   
location ~* ^/([^\/_]*)_([^\/_]*)/(.*)$ {
    proxy_pass http://frontend/$3;
}
location ^= /abc_{

}
location /abc_def/{

}
```

如果用户访问`/abc_def/index.html`, nginx将会先匹配到第3条, 然后继续进行普通匹配, 于是到了第4条, 但是第4条并没有以`=`或`^~`开头, 所以匹配到这里后, 还会继续进行正则匹配, 与是最终还是会到第2条.

## 3. rewrite 理解

首先`rewrite`有4种标志位: `last`|`break`|`redirect`|`permanent`.

用法: rewrite 正则表达式 替换字符串 标志位

其中`redirect`与`permanent`分别为301和302重定向, 而`last`|`break`更适合被称为 **url重写**. 区别在于

- 重定向会为当前请求返回`301`|`302`代码, 并在响应头`Location`字段中填写上`替换字符串`, 这将引起浏览器再次发起一次http请求(如果客户端是`curl`等命令就算了), 目标就是这个`Location`字段中重定向的url.

- url重写完成后会以重写后的url在nginx配置文件中进行匹配, 而不是像重定向那样干脆返回一个`Location`字段让浏览器再发起一次请求.

### 3.1 rewrite重写规则

不管是重定向还是重写, `rewrite`指令的`正则表达式`部分匹配的只会是`uri`路径部分, 不包括域名, 端口和请求参数. 所以匹配时是无法匹配到`http://`字段或是域名信息的, 这样会导致匹配失败. 如果必须要根据域名, 端口等进行匹配, 建议结合`if`指令使用.

而替换字符串目前发现的有3种情况:

- 以`http://域名[:端口]`开头的全路径.

- 以`/`开头的绝对路径, 将会以当前访问域名为前缀进行拼合.

- 不带`/`的路径, 这将会被视为以当前请求资源为准的相对路径, 会与当前资源的访问路径进行拼合.

原访问路径的请求参数不会被重写, 它默认追加到重写后的`uri`后面, 如果想舍弃原来的请求参数, 在`替换字符串`部分末尾加上'?'即可.

### 3.2 last与break的区别

参考文章

[nginx rewrite规则语法](http://blog.csdn.net/xiao_jun_0820/article/details/9397011)

我们知道, `rewrite`合法位置在`server`, `if`, `location`块内. 并且上面也说了, `last`与`break`的`rewrite`称为 **url重写** 更为贴切. 我们首先要理解url重写之后的处理流程.

假设nginx安装在`/usr/local/nginx`, 这样在nginx配置文件中使用`root html;`可以指定目标目录为`/usr/local/nginx/html`.

第1个例子. `rewrite`指令只出现在`if`块内.

在`html`目录下创建`download/music/index.html`, 并在该文件中写入`download/music/index.html`; 然后继续在`html`目录下创建`music/index.html`, 在其中写入`music/index.html`;

```
server {
    root html;
    if ($request_uri ~* ^/download){
        rewrite ^(/download)(.*)$ $2 last;
    }
    location /music/ {
        add_header request_uri $request_uri;
        add_header uri $uri;
    }   
}
```

我们访问`/download/music/index.html`, `if`语句会将其重写成`/music/index.html`, 继续匹配到`location`, 然后页面上显示的为`/music/index.html`, 这应该不难理解. nginx内置变量`$request_uri`表示请求的原始uri, 不可更改, 而`$uri`表示被`rewrite`/`try_files`重写后的值, 我们在浏览器中可以看到此页面的响应头中多了两个自定义的字段: `request_uri`与`uri`, 分别为`/download/music/index.html`与`/music/index.html`. 也印证了我们的猜测.

然后, 添加另外的条件, 将上面的配置修改为如下

```conf
server {
  root html;
  if ($request_uri ~* ^/download){
      rewrite ^(/download)(.*)$ $2 last;
  }
  if ($uri ~* ^/music){
      rewrite ^(/music)(.*)$ https://www.taobao.com redirect;
  }
  location / {
      add_header request_uri $request_uri;
      add_header uri $uri;
  }
}
```

好吧, 上面第2,3个`if`语句, 是为了验证`if`语句块中的`rewrite`是不是按照顺序来的. 第2,3条匹配的变量是`$uri`, 如果是按顺序来的, 访问`/download/music/index.html`, 页面应该会被重定向到淘宝网.

然而事实是, 页面显示的是`/music/index.html`, 即, **经`if`(也包多`server`块本身的)语句块中的`rewrite`指令, 一旦修改, 就会直接去匹配`location`, 之后任何其他的`if`匹配都不会执行...!!!**.

细思极恐的一点是...也许`if`块内`rewrite`重写之后, `if`语句与`location`语句之间的任何操作都不会执行...

```conf
server {
  root html;
  set $abc '123456';
  if ($request_uri ~* ^/download){
      rewrite ^(/download)(.*)$ $2 last;
  }
  ## set $abc '123456';
  location / {
      add_header request_uri $request_uri;
      add_header uri $uri;
      add_header abc $abc;
  }
}
```

访问`/download/music/index.html`, 返回的页面上还是`/download/music/index.html`, 而且响应头多了自定义字段`abc`, 值为`123456`; 尝试将`if`语句上面的`set`指令注释掉, 解开下面的`set`指令的注释, 刷新页面...响应头里没有`abc`字段了. 也就是说, **`if`与`server`块内的`rewrite`一旦完成url重写, 就立刻去匹配`location`, 之间的任何操作都不会执行**.

然后我们尝试将`if`块中`rewrite`的`last`标志位换成`break`, 再次重复上面的操作, 你会发现情况与使用`last`标记时完全相同, 因为`last`与`break`在`if`,`server`块内表现没有任何区别. 这两个地方出现的`rewrite`次数多了还会按顺序执行. **它们的区别在与`location`块内的使用.**

------

第2个例子, `rewrite`同时出现在`if`块与`location`块内. `html`目录结构与上面一个例子相同.

```conf
server {
  root html;
  if ($request_uri ~* ^/download){
      rewrite ^(/download)(.*)$ $2 last;
  }
  location /music/ {
      rewrite ^(.*)$ /download$1 last;
  }
}
```

我们还是访问`/download/music/index.html`, `if`语句将其修改为`/music/index.html`, 匹配到`location`, `location`块又将其重写成了`/download/music/index.html`, 但是之后并没有再次从`if`语句进行重写, 因为如果那样的话相当于陷入了死循环, 而实际上页面上显示了`/download/music/index.html`.

将`location`内中的`last`标记修改为`break`, 结果不变.

再进一步, 我们再添加一段`/download`的`location`匹配.

```conf
server {
  root html;
  if ($request_uri ~* ^/download){
      rewrite ^(/download)(.*)$ $2 last;
  }
  location /music/ {
      rewrite ^(.*)$ /download$1 last;
  }
  location /download/ {
      rewrite ^(/download)(.*)$ $2 last;
  }
}
```

再次访问`/download/music/index.html`, 不出意外的话, 将得到500错误;

将第1个`location`中的`last`标记修改成`break`, 第2个保持`last`不变, 刷新页面, 将得到`/download/music/index.html`;

将第2个`location`中的`last`标记修改成`break`, 第2个保持`last`不变, 刷新页面, 将得到`/music/index.html`.

这就是`last`与`break`的区别了, 这点区别只在`location`块内有所表现.

- `last`: url重写完成后立即结束当前`location`内的`rewrite`检测, 并且以重写后的uri继续对所有(包括本身)的`location`再次匹配;

- `break`: url重写完成后立即结束当前`location`内的`rewrite`检测, 并且不再匹配任何`location`与`if`语句, 直接以重写后的url为路径去访问目标文件.

上面的第1种情况, `if`语句将`/download/music/index.html`重写成`/music/index.html`, 匹配到第1个`location`, 又将其修改成`/download/music/index.html`, 虽然不再匹配上面的`if`语句, 但`last`标记表示结束当前`location`的`rewrite`匹配, 但又开始重新对所有`location`匹配, 于是这两个`location`之间陷入了死循环, nginx内部设置了`rewrite`的最大次数为10, 超过这个值就会返回500.

第2种情况, `if`语句将`/download/music/index.html`重写成`/music/index.html`, 匹配到第1个`location`, 又将其修改成`/download/music/index.html`, 此时`break`的作用就体现出来了, 它停止了当前`location`的所有`rewrite`, 直接去寻找`html`根目录下`/download/music/index.html`文件, 于是...

那第3种情况也很容易理解了...

这就是`last`与`rewrite`的区别了.

### 3.3 总结

1. `rewrite`正则部分匹配的是uri部分, 不包括`http(s)://`, 域名, 端口信息, 确切一点说, 这个部分最终匹配到的是当前nginx内部的`$uri`变量, 而`$uri`随着nginx的处理流程, 是可以被修改的.

2. location中正则匹配分组是可以被该location块内部的语句引用的.

2.
