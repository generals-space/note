# Css多行字符截取方法详解(转)

参考文章

1. [Css多行字符截取方法详解](http://www.php.cn/css-tutorial-387729.html)
    - 很不错的实验验证型文章
2. [-webkit-line-clamp超过两行就出现省略号](https://www.cnblogs.com/ldlx-mars/p/6972734.html)
    - -webkit-line-clamp 是一个不规范的属性(unsupported WebKit property), 没有出现在 CSS 规范草案中

## 1. 

来自参考文章1

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <style>
        .wordcut-line-2 {
            /*height与line-height的位数决定可显示的行数*/
            height: 2em;
            line-height: 1em;
            overflow: hidden;
        }
        .wordcut-line-2::before {
            content: '';
            float: left;
            width: 1px;
            height: 2em;
        }
        .wordcut-line-2::after {
            float: right;
            content: "...";
            /* 为三个省略号的宽度 */
            width: 3em;
            height: 1em;
            line-height: 1em;
            /* 使盒子不占位置 */
            margin-left: -3em;
            /* 移动省略号位置 */
            position: relative;
            left: 100%;
            top: -1em;
            padding-right: 1px; /*与.text类的margin-left属性对应*/
            text-align: right;
            background: linear-gradient(90deg,  rgba(255, 255, 255, 0) 0%, rgba(255, 255, 255, 1) 60%);
            box-sizing: content-box !important;  /* bootstrap会全局设置为border-box */
        }

        .wordcut-line-2 ._word {
            float: right;
            margin-left: -1px;
            width: 100%;
            word-break: break-all;
        }
    </style>
</head>
<body>
    <div class="wordcut-line-2">
        <div class="_word">
            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dignissimos labore sit vel itaque delectus atque quos magnam assumenda quod architecto perspiciatis animi.
        </div>
    </div>
</body>
</html>
```

## 2.

示例灵感来自[CSDN](https://www.cnblogs.com/ldlx-mars/p/6972734.html), 经过了参考文章2的改进.

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <style>
        div.wrapper{
            width: 892px;
            height: 100px;
            margin: 0 auto;
        }
        p.txt{
            font-size: 14px;
            color: #858585;
            line-height: 24px;
            display: -webkit-box;
            -webkit-box-orient: vertical;
            -webkit-line-clamp: 2;
            overflow: hidden;
            word-wrap:break-word;
            word-break:break-all;
        }
    </style>
</head>
<body>
    <div class="wrapper">
        <p class="txt">
            xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
        </p>
    </div>
</body>
</html>

```
