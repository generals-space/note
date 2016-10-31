# Shell中的各种括号

参考文章

[shell中的括号（小括号，中括号，大括号）](http://blog.csdn.net/tttyd/article/details/11742241)

[shell中的各种括号](http://blog.csdn.net/weihongrao/article/details/17007575)

[linux中双括号和双中括号，括号和中括号](http://blog.csdn.net/weihongrao/article/details/17006931)

## 1. 小括号(圆括号)

### 1.1 单小括号

#### 1.1.1 命令组

括号中的命令将会新开一个子shell并顺序执行，所以**括号中可以引用括号外面的变量, 但括号内定义的变量不能够被括号外面的部分使用, 并且括号中对外面变量的修改也不会生效**。括号中多个命令之间用分号隔开，最后一个命令可以没有分号，括号与之后的命令之间也需要有分号, 各命令和括号之间不必有空格。

```
## 括号中可以引用括号外面的变量
$ x1='123'
$ (echo $x1)
123

## 括号内定义的变量不能够被括号外面的部分使用
$ (x2='321')
$ echo $x2

## 括号中对外面变量的修改也不会生效
$ x1='123'
$ (x1='abc'; echo $x1); echo $x1
abc
123
```

#### 1.1.2 命令替换

等同于反引号`cmd`，shell扫描一遍命令行，发现了`$(cmd)`结构，便将`$(cmd)`中的cmd执行一次，得到其标准输出，再将此输出放到原来命令行中。有些shell不支持，如tcsh。

```
$ xy=$(echo '123')
$ echo $xy
123
```

#### 1.1.3 用于初始化数组

```
$ array=(a b c def)
$ echo ${array[2]}
c
$ echo ${array[3]}
def
```

### 1.2 双小括号

#### 1.2.1 整数运算

1. 四则运算: 加减乘除, 求余. (不过没有办法设置精度, 除法只能保留整数部分), 可用括号

2. 位运算: 与(&), 或(|), 非(^); 

3. 数值比较: `>` `<`, `>=` `>=`, `==` `!=`, 与(&&), 或(||), 返回值为1/0(真或假).

4. 三目运算符: (( 1 ? 2: 3))

**注意**

1.  **内容只可以是数字, 且只可以是整数**

2.  如果要将双小括号的结果赋值给变量, 需要使用$运算符, 如`$((1 + 2))`

3. 双括号内外共享变量, 内部还可以使用自增(++), 自减(--)运行符; 另外**双括号内部引用变量时不可以加`$`符号**

```
$ echo $((1+2))
3
$ echo $(( 2 || 1))
1

## 自增运算
$ a=1
$ ((a ++))
echo $a
2

## 三目运算, 不过好像没多大用, 只能进行数据运算
$ echo $(( 1==2 ? 4 : 7))
7
#### 如果第一个参数是空字符串或0, 就返回第3个参数的值(即认为这个条件为假);反之, 就返回第2个参数串的值. 但前提是, 第2, 3个参数为数值类型, 如果它们是字符串, 那就会返回0
$ a=123
## 变量a不用加$符哦
$ echo $(( a ? 8 : 9 ))
8
$ a=
$ echo $(( a ? 8 : 9 ))
9

$ a=123
## 貌似原因是abc被当作了变量, 所以返回0
$ echo $(( a ? abc : def ))
0
## 但是不能为abc, def赋值为字符串, 双小括号会认为存在语法错误
$ abc='abc'
$ echo $(( a ? abc : def ))
-bash: abc: expression recursion level exceeded (error token is "abc")
```

#### 1.2.2 扩展C语言条件判断语法

这一节与单中括号与单大括号相关, 建议先了解这两者的用法.

默认情况下, shell中的条件判断如`if`, `while`, `for`的语法类似于

```
## if [ 1 = 1 ]; then shell命令; fi
if [ 1 == 1a ]; then echo yes; else echo no; fi
## for i in {1..10}; do shell命令; done
for i in {1..10}; do touch test$i.txt; done
## while [ a -lt 10 ]; do shell命令; done
while [ $a -lt 10 ]; do echo $a; ((a++)); done
```

有了(()), 可以使用类似于C语言的循环语句

**if**

```
## 默认[]中是不可以使用`&&`与`||`操作符, 并且数值比较不能使用> < =等符号的, 双小括号中可以, 这一点与双中括号类似
$ if (( 'abc' == 'abc' || 1 == 2 )); then echo yes; else echo no; fi
yes
$ if [[ 'abc' == 'abc' || 1 == 2 ]]; then echo yes; else echo no; fi
yes
```

**for**

```
$ for ((a=0; a<10; a++)); do echo $a; done
0
1
2
3
4
5
6
7
8
9
```

```
$ a=0
$ while ((a<10)); do echo $a; ((a++)); done
0
1
2
3
4
5
6
7
8
9
```

> 在作判断条件时, 双括号中的内容不再局限于数字, 也可以是字符串.

## 2. 中括号(方括号)

### 2.1 单中括号

等同于linux中的`test`命令, 都可用于`if`, `while`语句的条件判断. 

可判断的条件有:

1. 文件系统相关(文件是否存在, 路径是否为目录等)

2. 字符串判断, 只能使用`==`与`!=`(< , > 比较没有意义)

3. 数值比较, 需要使用相关的`-eq`, `-gt`等操作符.

4. 逻辑判断, `-a`, `-o`标识符

```
$ a=0
$ while [ $a -lt 10 ]; do echo $a; a=`expr $a + 1`; done
0
1
2
3
4
5
6
7
8
9
```

**注意**

1. [, ]左右都需要有空格

2. [ ] 中字符串或者${}变量尽量使用**双引号**扩住，以避免值未定义引用而出错

### 2.2 双中括号

**双中括号比单中括号更加通用的使用方式**.

数值判断, 可以不必使用单中括号的`-eq`, `-lt`等操作符, 直接使用`>`, `<`, `=`, `!=`等符号(貌似`>=`, `<=`不太好使)

逻辑判断, 可以直接使用`&&`, `||`

另外, 字符串比较时, 可以使用通配符和正则表达式两种. 右边的字符串可以是一个模式, 但要注意不能用引号包裹(单双引号都不行).

不过**单中括号的注意点双中括号也需要留心**

```
$ if [[ 1 == 1 ]]; then echo yes; else echo no; fi
yes
$ a='abc'
$ if [[ $a == 'abc' ]]; then echo yes; else echo no; fi
yes
## >=和<=不太好用
$ if [[ 1 >= 1 ]]; then echo yes; else echo no; fi
-bash: syntax error in conditional expression
-bash: syntax error near `1'

## 通配符模式, ?匹配任意单一字符, *匹配任意个任意字符, 右边为模式, 不可以用引号包裹

$ if [[ "hashes" == "hash??" ]]; then echo yes; else echo no; fi
no
$ if [[ "hashes" == hash?? ]]; then echo yes; else echo no; fi
yes
$ if [[ "hashes" == hash* ]]; then echo yes; else echo no; fi
yes

## 正则模式, 要使用`=~`符号

### 这第一行好像有点正则和通配符混用的感觉啊?
$ if [[ "hashes" =~ hash?? ]]; then echo yes; else echo no; fi
yes
$ if [[ "hashes" =~ hash[ed]s ]]; then echo yes; else echo no; fi
yes
$ if [[ "hashes" =~ hash(ed)s ]]; then echo yes; else echo no; fi
no
$ if [[ "hashes" =~ hash(e|d)s ]]; then echo yes; else echo no; fi
yes
$ if [[ "hashes" =~ ^hash(e|d)s ]]; then echo yes; else echo no; fi
yes
```

## 3. 大括号(花括号)

### 3.1 单大括号(没有双的)

#### 3.1.1 字符串生成

1. 两个.点号生成顺序字符串

2. ,逗号分隔, 不可以有空格

```
$ touch test{1..4}.txt
$ ls 
test1.txt  test2.txt  test3.txt  test4.txt

$ touch {test{1..4},testab}.txt
$ ls 
test1.txt  test2.txt  test3.txt  test4.txt  testab.txt
```

#### 3.1.2 代码块

代码块，又被称为内部组，这个结构事实上创建了一个匿名函数 。与小括号中的命令不同，大括号内的命令不会新开一个子shell运行，即脚本余下部分仍可使用括号内变量。括号内的命令间用分号隔开，**最后一个也必须有分号**。**{}中的第一个命令和左括号之间必须要有一个空格**。

```
$ { a=1; ((a++));}; echo $a;
```

#### 3.1.3 模式匹配/替换

几种特殊的替换结构：`${var:-string}`, `${var:+string}`, `${var:=string}`, `${var:?string}`

作用类似于类C语言中的三目运算符, js中类似于`&&`, `||`的条件判断.

A. `${var:-string}`和`${var:=string}`: 若变量`var`为空，则用在命令行中用string来替换`${var:-string}`，否则用变量var的值来替换`${var:-string}`；对于`${var:=string}`的替换规则和${var:-string}是一样的，所不同之处是${var:=string}若var为空时，用string替换${var:=string}的同时，把string赋给变量var： ${var:=string}很常用的一种用法是，判断某个变量是否赋值，没有的话则给它赋上一个默认值。

```
## abc为空
$ abc=''
$ echo ${abc:-123}
123
$ echo $abc

$ abc=''
$ echo ${abc:=123}
123
$ echo $abc
123

## abc不为空
$ abc='321'
$ echo ${abc:-123}
321
```

B. `${var:+string}`的替换规则和上面的相反，即只有当var不是空的时候才替换成string，若var为空时则不替换或者说是替换成变量 var的值，即空值。(因为变量var此时为空，所以这两种说法是等价的) 

```
$ abc=''
$ echo ${abc:+321}

$ abc='123'
$ echo ${abc:+321}
321
```

C. `${var:?string}`替换规则为：若变量var不为空，则用变量var的值来替换`${var:?string}`；若变量var为空，则把string输出到标准错误中，并从脚本中退出。我们可利用此特性来检查是否设置了变量的值。

```
$ echo $abc
123
$ echo ${abc:?hehe}
123
$ abc=''
$ echo ${abc:?hehe}
-bash: abc: hehe
```

> PS：在上面这五种替换结构中string不一定是常值的，可用另外一个变量的值或是一种命令的输出。