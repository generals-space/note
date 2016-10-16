# Json添加注释(转)

原文链接

[【部分解决】Json中添加注释](http://www.crifan.com/add_comments_for_json/)

## 问题描述

通过json文件给python脚本传递参数, 但是希望每个参数都有对应的注释, 以方便使用者知道该参数的确切含义.

问题转化为给json中添加注释.

## 解决过程

1.网上找了json的官网[JSON](http://www.json.org/),没看到关于添加注释的说明.

就像某人说的, json的出现, 本身就是为了压缩, 减少数据量, 所以理论上就不支持注释, 也是可以理解的.

2.参考：[Json文件如何加注释](http://www.oschina.net/question/163912_26244)

去尝试：

```json
{

…

}

// end of json comment test –> no work
```

```json
// begin of json comment test –> no work

{

…

}
```

的结果, 也还是无效.

其他的, 也试过了：

/* comments here */

都不行.

3.参考：[javascript – Can I comment a JSON file? – Stack Overflow](http://stackoverflow.com/questions/244777/can-comments-be-used-in-json)

得到一个妥协的办法, 那就是, 把需要添加的注释, 当成json中的某个key和value, 就像普通的数据一样, 比如那位给出的例子：

```json
{ 
   "_comment" : "这里是要添加的数据...", 
   "glossary": { 
      "title": "example glossary", 
      "GlossDiv": { 
         "title": "S", 
         "GlossList": { 
            "GlossEntry": { 
               "ID": "SGML", 
               "SortAs": "SGML", 
               "GlossTerm": "Standard Generalized Markup Language", 
               "Acronym": "SGML", 
               "Abbrev": "ISO 8879:1986", 
               "GlossDef": { 
                  "para": "A meta-markup language, used to create markup languages such as DocBook.", 
                  "GlossSeeAlso": ["GML", "XML"] 
               }, 
               "GlossSee": "markup" 
            } 
         } 
      } 
   } 
}
```

目前看来, 除此之外, 也没啥其他的好办法了.

4.另外, 也看到[这里](http://blog.getify.com/json-comments/)在讨论, 给json组织建议, 添加对应的spec规范说明, 希望支持：

json的decoder编码出来的数据, 不包含对应的comment, 但是encoder应该支持comments

然后希望对应的comments的格式是

```json
// single line comment

/* multi line comments */
```

之类的.

然后在传输数据过程中, 则不需要传输这些comments.

等等讨论和建议.个人觉得还是蛮合理的, 只是此刻, 我用的python 2.7中的json, 还是不支持`decode`带`comments`的json啊.

## 总结

目前我这里的Python 2.7中的json，不支持类似于`//xxx`和`/* xxx*/`的注释，暂时的妥协办法只能是，把需要添加的注释，当做数据，写入到json里面。虽然效率很低，但是也只能这样了。

希望以后json的`encoder`和`decoder`支持对应的带`comments`的编解码。
