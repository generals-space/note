# JS浏览器API

参考文章

1. [Web API 接口 - MDN官方文档](https://developer.mozilla.org/zh-CN/docs/Web/API)

## 1. js刷新当前页面

```js
history.go(0);
location.reload();
location = location
location.assign(location);
location.replace(location);
```

## 2. 停止加载当前页面

```js
// 已证实无效
document.execCommand("stop");
// 使用这种方法时, 可访问到的资源请求已经发送的情况下依然可以接收到服务器传来的资源, 而未能建立连接的资源可以停止接收
window.stop();
```

## 3. 自动刷新页面

1.页面自动刷新：把如下代码加入<head>区域中

```html
<meta http-equiv="refresh" content="20">
```

其中`20`指每隔20秒刷新一次页面.

2.页面自动跳转：把如下代码加入<head>区域中

```html
<meta http-equiv="refresh" content="20;url=http://www.jb51.net">
```

其中20指隔20秒后跳转到http://www.jb51.net页面

3.页面自动刷新js版

```html
<script language="JavaScript">
function myrefresh()
{
       window.location.reload();
}
setTimeout('myrefresh()',1000); //指定1秒刷新一次
</script>
```

## 4. 文本选中行为

获取选中文本

```js
function getSelectionText() {
    //如果是Firefox、Chrome、Safari、Opera
    if (window.getSelection) { 
        return window.getSelection().toString();
    //如果是IE 
    } else if (document.selection && document.selection.createRange) { 
        return copytext_keleyi_com = document.selection.createRange().text;
    }
    return '';
}
```