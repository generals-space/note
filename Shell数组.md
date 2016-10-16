# Shell数组

原文链接

[Shell数组：shell数组的定义、数组长度](http://c.biancheng.net/cpp/view/7002.html)

bash支持一维数组（不支持多维数组），并且没有限定数组的大小。类似与C语言，数组元素的下标由0开始编号。获取数组中的元素要利用下标，下标可以是整数或算术表达式，其值应大于或等于0。

## 1. 定义数组

在Shell中，用括号来表示数组，数组元素用“空格”符号分割开。定义数组的一般形式为：`array_name=(value1 ... valuen)`

例如：

```
array_name=(value0 value1 value2 value3)
```

或者

```
array_name=(
value0
value1
value2
value3
)
```

还可以单独定义数组的各个分量：

```
array_name[0]=value0
array_name[1]=value1
array_name[2]=value2
```

可以不使用连续的下标，而且下标的范围没有限制。

## 2. 读取数组

读取数组元素值的一般格式是：`${array_name[index]}`

例如：

```
valuen=${array_name[2]}
```

举个例子：

```
#!/bin/sh
NAME[0]="Zara"
NAME[1]="Qadir"
NAME[2]="Mahnaz"
NAME[3]="Ayan"
NAME[4]="Daisy"
echo "First Index: ${NAME[0]}"
echo "Second Index: ${NAME[1]}"
```

运行脚本，输出：

```
$./test.sh
First Index: Zara
Second Index: Qadir
```

使用`@`或`*`可以获取数组中的所有元素，例如：

```
${array_name[*]}
${array_name[@]}
```

举个例子：

```
#!/bin/sh
NAME[0]="Zara"
NAME[1]="Qadir"
NAME[2]="Mahnaz"
NAME[3]="Ayan"
NAME[4]="Daisy"
echo "First Method: ${NAME[*]}"
echo "Second Method: ${NAME[@]}"
```

运行脚本，输出：

```
$./test.sh
First Method: Zara Qadir Mahnaz Ayan Daisy
Second Method: Zara Qadir Mahnaz Ayan Daisy
```

**获取数组的长度**

获取数组长度的方法与获取字符串长度的方法相同，例如：

```
# 取得数组元素的个数
length=${#array_name[@]}
# 或者
length=${#array_name[*]}
# 取得数组单个元素的长度
lengthn=${#array_name[n]}
```

## 3. 扩展

使用索引直接定义数组时, 可以跳跃定义

```
## 不定义array[1]的值
array[0]=value0
array[2]=value2
```

这种情况下, 数组的长度是实际有数据的元素个数而不是索引范围. 这一点与高级编程语言不同.