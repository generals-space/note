# Shell脚本-EOF标记

参考文章

[cat和EOF的使用](http://luxiaok.blog.51cto.com/2177896/711822)

EOF: End Of File, 表示文本结束符.

## 1. 初级使用方法

`EOF`最基本的使用方法就是输出多行字符串到文件. 如下

```
#!/bin/bash
cat > eof.txt << EOF
第1行
第2行
第3行
EOF
```

上面的命令会生成`eof.txt`文件, 其内容就是列出的3行文字. 这种使用方法在比如脚本中内嵌一段比较长的自定义的配置, 需要输出到指定配置文件时极其有用. 当然, `>`与`>>`重定向符都是可以使用的.

虽然可以使用`echo -e "第1行\n第2行\n第3行\n" > eof.txt`达到同样的效果, 但是可读性很差.

> PS: `echo`的`-e`选项将会解析其中的反斜线`\`转义, 不只是`\n`换行符, 还有`\t`制表符等. `-e`选项的出现不因后面的字符串是单引号还是双引号而有所区别, 双引号时`$`变量引用依然有效, 单引号就不行了.

```
## 单引号换行
$ echo -e '1\n2\n3'
1
2
3
## 双引号换行
$ echo -e "1\n2\n3"
1
2
3
```

------

~~在命令行中, 第2个`EOF`符并不是通过EOF字符串表示的, 而是`Ctrl+D`~~. 扯, 我自己试验的时候EOF字符串有效, 而`Ctrl+D`报错. 

```
$ cat > test << EOF
> 1
> 2
> 3
> EOF
$ cat test 
1
2
3
$ cat > test << EOF
> 1
> 2
> 3
> -bash: warning: here-document at line 37 delimited by end-of-file (wanted `EOF')
```

## 2. 进阶使用方法

### 2.1 内嵌变量

待输出的多行文本中可以包含变量, 在输出到文件之前就会被解析. 比如

```
#!/bin/bash

line_num=3
cat > test.txt << EOF
第1行
第2行
第$line_num行
EOF
```

生成的`test.txt`文件内容为

```
第1行
第2行
第3行
```

表面看来倒是很和谐美好, 但如果包含的变量并不是对当前shell脚本而言的呢? 比如

```
#!/bin/bash

cat >> /etc/profile << EOF
export JAVA_HOME=/usr/local
export PATH=$JAVA_HOME/bin:$PATH
EOF
```

上述脚本输出到`/etc/profile`的内容是

```
export JAVA_HOME=/usr/local
export PATH=/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/root/bin
```

看到了?

第2个`export`的语句将`$JAVA_HOME`与`$PATH`的值都预先解析了出来, 然而原脚本中并没有`JAVA_HOME`的定义, 所以它的值为空.

------

现在我们的目的是追加`$JAVA_HOME`与`$PATH`字符串到`/etc/profile`, 方法有两个:

1. 带`$`符的变量前加反斜线`\`转义

2. 第一个EOF带单引号, 即`'EOF'`

