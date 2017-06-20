# CSS3滤镜-filter属性

<!tags!>: <!CSS!> <!CSS3!> <!滤镜!>

参考文章

[CSS3 filter(滤镜) 属性](http://www.runoob.com/cssref/css3-pr-filter.html)

[MDN-filter](https://developer.mozilla.org/zh-CN/docs/Web/CSS/filter)

```css
/*修改所有图片的颜色为黑白 (100% 灰度)*/
img{
    -webkit-filter: grayscale(100%);
    -moz-filter: grayscale(100%);
    -ms-filter: grayscale(100%);
    -o-filter: grayscale(100%);
    filter: grayscale(100%);
}
```

```
filter: none | blur() | brightness() | contrast() | drop-shadow() | grayscale() | hue-rotate() | invert() | opacity() | saturate() | sepia() | url();
```

> 提示: 可使用空格分隔多个滤镜