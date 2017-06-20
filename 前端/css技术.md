
> CSS 是用来表现的，HTML 一切非正文的、装饰性的东西理论上都是要用 CSS 来实现的

选中元素样式
------

```css
::selection {background-color: #38485a;color: #fff;text-shadow: 0 1px 0 rgba(0, 0, 0, 0.3);}
::-moz-selection { background-color: #38485a; color: #fff; text-shadow: 0 1px 0 rgba(0, 0, 0, 0.3) }
img::selection { background: transparent }
img::-moz-selection { background: transparent }
```

清除浮动
------

```css
.clearfix:before,
.clearfix:after {
  content: ".";
  display: block;
  height: 0;
  visibility: hidden;
}
.clearfix:after {clear: both;}
.clearfix {zoom: 1;} /* IE < 8 */
```