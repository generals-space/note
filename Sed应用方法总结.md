# Sed应用方法总结

## 1. 追加append与插入insert

参考文章

[linux下在某行的前一行或后一行添加内容](http://www.361way.com/sed-process-lines/2263.html)

### 1.1 在某行的前一行或后一行添加内容

```
#匹配行前加
sed -i '/目标行匹配内容/i待添加新行内容' 目标文件
#匹配行前后
sed -i '/目标行匹配内容/a待添加新行内容' 目标文件
```

在书写的时候为便与区分，往往会在`i`和`a`后(新行内容前)加一个反斜扛.

```
#匹配行前加
sed -i '/目标行匹配内容/i\待添加新行内容' 目标文件
#匹配行前后
sed -i '/目标行匹配内容/a\待添加新行内容' 目标文件
```

### 1.2 在某行(已知具体行号)前或后加一行内容

```
## 第4行后追加
sed -i 'N;4a待添加新行内容' 目标文件
## 第4行前插入
sed -i 'N;4i待添加新行内容' 目标文件 
```

当前, 与上面一样, 也可以加上反斜线`\`

```
## 第4行后追加
sed -i 'N;4a\待添加新行内容' 目标文件
## 第4行前插入
sed -i 'N;4i\待添加新行内容' 目标文件 
```

**注意: 行号需要大于1, 不能等于.**

以如下文件为例, `index.html`

```html
<html>
<head>
</head>

<body>
hello world
</body>
</html>
```

以下两行都无效

```
$ sed -i 'N;1a\<!DOCTYPE>' ./index.html
$ sed -i 'N;1i\<!DOCTYPE>' ./index.html
```

下面的才有效, 会在最开始添加`<!DOCTYPE>`这个标记

```
$ sed -i 'N;2i\<!DOCTYPE>' ./index.html
```

然后`index.html`会被修改成

```
<!DOCTYPE>
<html>
<head>
</head>

<body>
hello world
</body>
</html>
```

也就是说, **`N;行号`中指定的行号实际上等于目标行号+1**. 简直...不可理喻.

## 2. sed递归操作

参考文章

[sed的递归问题](http://www.blogbus.com/kebe-jea-logs/59348026.html)

首先要明确, sed不能递归目录进行查找匹配(`p`选项)与文件操作(`s`与`a`等选项). 如果待查找的目录下存在子目录, 会报错如下.

```
sed: read error on ./expect交互脚本: Is a directory
```

貌似`sed`并没有提供递归操作的支持, 所以需要寻找替换方案. `p`选项这样的查找匹配工作, 完全可以交给`grep`命令来做. `s`与`a`选项则必须先查找再进行对原文件的修改. sed需要拿到所有匹配条件的目标文件路径, 所以需要其他命令提供给它这个路径.

### 2.1. 配合find

对文件名, 修改时间有要求的情况, 可以配合使用`find`. 比如在如下目录结构中, 要在所有`.html`文件的第一行添加`<!DOCTYPE>`.

```
$ tree
.
├── common
│   └── header.html
├── css
├── index.html
└── js

$ sed -i 'N;1i\<!DOCTYPE>' ./*
sed: couldn't edit ./common: not a regular file
```

首先`find`找出所有`.html`文件, 可以看到结果信息中包含着目标文件的路径信息, 这正是我们想要的.

```
$ find ./ -name '*.html'
./common/header.html
./index.html
```

然后用`xargs`将查找到的文件作为参数传给`sed`.

```
## 注意sed指定行号时, 实际上指定行号=目标行号+1
$ find ./ -name '*.html' | xargs sed -i 'N;2i\<!DOCTYPE>'
```

也可以使用`find`自带的对查找出的文件进行指定操作的`-exec`选项, 其中`{}`表示查找到的文件.

```
$ find ./ -name '*.html' -exec sed -i 'N;2i\<!DOCTYPE>' {} \;
```

### 2.2 配合grep

如果想要递归的替换指定字符串, 需要先使用`grep`查找目录中符合条件的文件路径信息, 然后传递给`sed`. 要实现这个功能, 会用到`grep`的`-l`选项.

测试用的目录树结构如下

```
$ tree
.
├── aa.txt
└── bb
    └── bb.txt
```

原始的`grep`命令查找的结果如下.

```
$ grep -ri 'cc' ./*
./aa.txt:ccc
./bb/bb.txt:cc
```

使用`-l`参数可以单独输出文件路径而不再输出匹配行的内容.

```
$ grep -ril 'cc' ./*
./aa.txt
./bb/bb.txt
```

将`grep`查询到的结果作为`sed`的输入参数.

```
## 注意sed命令中的被替换字符串要与grep查找的字符串相同, 都是'cc'
$ sed -i 's/cc/hh/g' $(grep -rl 'cc')
```

这样, `sed`可以递归地将用`grep`查到的包含'cc'行中的'cc'都替换成'hh'

> ps: 配合使用grep的`--include`选项会更加灵活