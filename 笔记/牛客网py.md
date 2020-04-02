## python

### 常用函数总结

#### split出空成员

```py
s = 'ABCDEF'
print(s.split('A')) ## ['', 'BCDEF']
```

`split()`会有空成员, 而不是只有`['BCDEF']`.

#### set()

```py
## 会初始化一个空set, 之后可以用myset.add()添加成员
myset = set()
```

### 遍历数组带索引

```py
for k, v in enumerate(array):
    pass
```

### `elif`不是`else if`

```py
if:
    pass
elif:
    ## 不是 else if
    pass
else:
    pass
```

### 标识符首字母大写

```py
True
False
None
```

### 交互式输入

`raw_input`只在`python2`, `input`才是`python3`.

```py
for line in sys.stdin:
    try:
        number = int(line.strip())
        ## 
    except:
        break
```

### input与sys.stdin

在做一道题的时候总是报"请检查是否存在语法错误或者数组越界非法访问等情况". [题目链接](https://www.nowcoder.com/profile/4027609/codeBookDetail?submissionId=12213544)

```py
while True:
    num_str = input()
    if num_str == '': break
```

我看使用python通过的示例代码都没有使用`input()`, 而是用了`sys.stdin.readline()`.

```py
import sys

while True:
    num_str = sys.stdin.readline()
    if num_str == '': break

```

这样就会只报未通过测试用例了.

查了查好像两者并没有明显的不同点, 不明白为什么会出错.

## 左闭右开区间

```py
a = [1, 2, 3, 4]
a[1:3] ## [2, 3] 只有两个元素
a[1:1] ## [] 空列表


for i in range(2) ## i的取值有0,1, 注意没有2, 所以for..range..也是左闭右开区间.
```

