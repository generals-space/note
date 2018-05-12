# Django模板转义与渲染技巧

参考文章

1. []()

2. [django模板的{%%}标记与js模板冲突](https://segmentfault.com/q/1010000002466016)

3. [如何关闭Django模板的自动转义](http://www.cnblogs.com/cyiner/archive/2011/11/10/2244838.html)

## 1. 直接输出特殊字符

在使用前端模板时, 需要在html文件中输出`{{`与`}}`, 然而django在使用`render`函数生成html响应时会自动对其html模板中的`{%`与`%}` 进行渲染. 

常见的使用场景是

```py
def index(req):
    return render(req, 'index.html')
```

在使用`artTemplate.js`时(关键字为`{{`与`}}`), 出现类似下面这样的错误, 是因为`django`在遇到`{{`的第一个`{`就开始进行变量的赋值解析操作了, 但又找不到与之匹配的`%`, 所以会出错.

```
{{each list as value }}    #出错行

TemplateSyntaxError at /posts
Could not parse the remainder: ' list as value' from 'each list as value'
```

解决方法

使用`{% templatetag openvariable %}`与`{% templatetag closevariable %}`标签, 将分别输出`{{`与`}}`到html. 于是

```
{% templatetag openvariable %} each list as value {% templatetag closevariable %}
```

将会被django正常地输出为

```
{{each list as value }}
```

到静态html文件, 之后就可以被`artTemplate`编译渲染了.

另外, 有些其他的渲染模板是根据`{%`与`%}`作为变量标签的, 鉴于这些情况, django还提供了其他对应的输出方案:

```
openvariable {{
closevariable }}
openblock {%
closeblock %}
openbrace {
closebrace }
opencomment {#
closecomment #}
```

类比上面的示例, 不难得到`templatetag`的正确使用方法.

## 2. 关闭渲染

上面的上面, 归根结底是通过django的特殊变量将我们想要的字符直接输出到html, 如`{{`, `{%`, 好在它们在django都有别名.

如果对一个单纯的前端模板, 希望render在执行时不进行渲染呢? 可以使用如下标签

```
{% verbatim %}
{% endverbatim %}
```

处于两者之间的内容将直接输出, 不进行变量的替换. 如

```html
{% verbatim %}
<form class="form-horizontal" role="form">
    <div class="form-group">
        <input type="text" value="{{ name }}">
    </div>
</form>
{% endverbatim %}
```

## 3. 关闭自动转义

上面两种情况说的还是模板引擎关键字冲突的问题, 现在我们还要关注html元素转义的问题.

Django的模板中会对HTML标签和JS等语法标签进行自动转义，原因显而易见，这样是为了安全。但是有的时候我们可能不希望这些HTML元素被转义，比如我们做一个内容管理系统，后台添加的文章中是经过修饰的，这些修饰可能是通过一个类似于FCKeditor编辑加注了HTML修饰符的文本，如果自动转义的话显示的就是保护HTML标签的源文件。为了在Django中关闭HTML的自动转义有两种方式，如果是一个单独的变量我们可以通过过滤器“|safe”的方式告诉Django这段代码是安全的不必转义。比如：

```html
<p>这行代表会被自动转义</p>: {{ data }}
<p>这行代表不会被自动转义</p>: {{ data|safe }}
```

其中第二行我们关闭了Django的自动转义。
我们还可以通过`{%autoescape off%}`的方式关闭整段代码的自动转义，比如下面这样：

```html
{% autoescape off %}
    Hello {{ name }}
{% endautoescape %}
```