# JS操作系统剪切板

<!tags!>: <!js!> <!剪切板!>

参考文章

1. [ZeroClipboard的时代或许已经过去了](http://div.io/topic/1384)

2. [Vimium P操作源码](https://github.com/philc/vimium/blob/master/lib/clipboard.coffee)

3. [W3C 标准 Clipboard API and events](https://www.w3.org/TR/clipboard-apis/)

在chrome中, 可以使用`document.execCommand()`方法, 直接复制指定`input`, `textarea`中的内容, 也可以将系统剪切板的内容粘贴到这些指定元素中. 

> 注意: 必须是这些元素中选中的内容, 可以通过`select()`方法将目标元素内容全选.

另外, `document.execCommand()`方法可以接受3种参数: `cut`, `copy`, `paste`.

```js
/*
    获取剪切板内容

    在chrome扩展中使用这getFromClipboard()方法时, 需要在manifest.json文件中的permissions字段声明'clipboardRead'的权限, 但是addToClipboard()不需要声明额外权限.
*/

/*
    @function: 将指定文本拷贝到剪切板
    @textStr: 目标字符串
*/
function addToClipboard(textStr){
    var buf = document.querySelector('#G_CopyBuf');
    if (!buf) {
        buf = document.createElement('textarea');
        buf.id = 'G_CopyBuf';
        buf.style.position = 'absolute';
        buf.style.left = '-9999px';
        buf.style.top = '-9999px';
        document.body.appendChild(buf);
    }
    buf.value = textStr;
    buf.focus();
    buf.select();
    document.execCommand('Copy', false, null);
    document.body.removeChild(buf);
}

/*
    @function: 读取并返回系统剪切板的内容(只有字符串)
    @return: 返回从系统剪切板读取的字符串.
*/
function getFromClipboard(){
    var buf = document.querySelector('#G_PasteBuf');
    if (!buf) {
        buf = document.createElement('textarea');
        buf.id = 'G_PasteBuf';
        buf.style.position = 'absolute';
        buf.style.left = '-9999px';
        buf.style.top = '-9999px';
        document.body.appendChild(buf);
    }
    buf.focus();
    document.execCommand('Paste');
    var textStr = buf.value;
    document.body.removeChild(buf);
    return textStr;
}
```