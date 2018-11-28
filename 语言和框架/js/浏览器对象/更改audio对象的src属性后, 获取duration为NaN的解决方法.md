# 更改audio对象的src属性后, 获取duration为NaN的解决方法

参考文章

1. [HTML DOM Audio 对象](http://www.w3school.com.cn/jsref/dom_obj_audio.asp)

2. [更改audio src路径后，获取duration为NaN的解决方法](https://blog.csdn.net/chenjineng/article/details/77650870)

3. [HTML5 Audio and Video 的新属性简介](https://www.cnblogs.com/zxjwlh/p/4547662.html)

在使用vue开发一个类似音乐播放器的页面时(iot-ui), 采用了如下的解决方法

```js
audioBox = {
    audioEle: null,
    isPlaying: false,
    playingMediaId: null
};
this.audioBox.audioEle = new Audio()
this.audioBox.audioEle.onended = function(event) {
    this.stopIt()
}.bind(this);
```

...嗯, 不要在意细节.

就是全局使用`audioBox.audioEle`来表示实际的播放组件, 不用在`DOM`树中创建`audio`标签(反正也没想让用户看见, 原生audio元素太难看了).

各种操作都是通过js来完成的(播放停止切换歌曲等). 由于只有一个`audio`对象, 所以在切换音乐时就需要把ta的`src`属性改为目标歌曲的链接. 但我还想获得这个音乐的时长, `audio`对象有一个属性为`duration`来表示. 但是在我切换为新的歌曲链接后, 得到的`duration`为`NaN`, 即不合法的数字类型...

最初以为是因为音频数据流没有加载的原因, 尝试为`audio`对象添加`preload`属性, 也试过显式调用`load()`方法, 但都没有用.

后来找到了参考文章2, 意识到是异步操作的原因. 使用ta提到的`oncanplay`事件回调方法的确得到了正确的时长.

注意: 回调函数的注册方式. 如下是错误的

```js
audio.oncanplay(function(event) {
    console.log(audio.duration);
});
```

正确的是这样

```js
// 第一个参数为事件对象, 如果要绑定上下文需要使用bind()方法.
audio.oncanplay = function(event) {
    console.log(audio.duration);
}
```

------

在找这个问题的解决方法时, 我还发现了参考文章3.

这篇文章对h5中媒体对象的属性, 方法和事件罗列得非常详尽, 在其中我发现了很多可用的东西.

由于曾经手动尝试调用`audio`的`load()`方法, 所以实验了下`onload`事件, 但是根本就没有触发...

w3school网站上有提到, 支持onload事件的 HTML 标签有:

```
<body>, <frame>, <frameset>, <iframe>, <img>, <link>, <script>
```

...难道是真的?

然后在参考文章3中我又看到对`oncanplay`的解释: **浏览器由于后续资源不足自动暂停播放，经数据缓冲之后，可以恢复播放时触发的事件**.

在单纯想得到`audio`对象的时长时, 其实只需要ta的元属性即可, 正好有一个`onloadedmetadata`事件, 我想这更符合我们的目标, 经过实验也证明有效. good