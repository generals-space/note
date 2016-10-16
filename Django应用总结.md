# Django应用总结

## 1. 模板转义

在使用artTemplate.js, 以ajax方式将响应渲染到页面时, 需要在html文件中输出`{{`与`}}`, 然而django在生成html响应时会自动对其html模板中的`{%`与`%}` 进行变量替换. 

而类似下面这样的错误, 是因为`django`在遇到`{{`的第一个`{`就开始进行变量的赋值解析操作了, 但又找不到与之匹配的`%`, 所以会出错.

```
{{each list as value }}    #出错行

TemplateSyntaxError at /posts
Could not parse the remainder: ' list as value' from 'each list as value'
```

解决方法

使用`{% templatetag openvariable %}`与`{% templatetag closevariable %}`标签, 将分别输出`{{`与`}}`到html而不是把'templatetag'等当作变量来替换. 于是`{% templatetag openvariable %} each list as value {% templatetag closevariable %}`将会被django正常地输出为`{{each list as value }}`到静态html文件, 之后就可以被`artTemplate`编译渲染了.

另外, 其他的渲染模板是根据`{%`与`%}`作为变量标签的, 鉴于这些情况, django还提供了其他对应的输出方案:

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