# powershell-定义变量

参考文章

1. [Powershell 定义变量](http://www.pstips.net/powershell-define-variable.html)

2. [Powershell变量的类型和强类型](http://www.pstips.net/powershell-variable-strongly-typing.html)

3. [A bunch of cmdlets that i tend to have use of… sometimes.](http://blog.worldofjani.com/?cat=10)

## 1. 定义变量

作为脚本语言, `powershell`不需要显示地去声明，可以自动创建变量，只须记住**变量的前缀必须为`$`**.

创建好了变量后，可以通过变量名输出变量，也可以把变量名存在字符串中. 但是有个例外单引号中的字符串不会识别和处理变量名, 这和linux shell行为相似, 很好理解. 

### 1.1 选择变量名

在powershell中变量名均是以美元符”$”开始，剩余字符可以是数字、字母、下划线的任意字符，并且powershell变量名大小写不敏感（$a和$A 是同一个变量). 

某些特殊的字符在powershell中有特殊的用途，一般不推荐使用这些字符作为变量名. 当然你硬要使用，请把整个变量名后缀用花括号括起来. 

```ps1
PS C:\test> ${"I"like $}="mossfly"
PS C:\test> ${"I"like $}
mossfly
```

...好神奇有木有(⊙ˍ⊙)

...虽然有点鸡肋

### 1.2 赋值和返回值

赋值操作符为“=”，几乎可以把任何数据赋值给一个变量，甚至一条cmdlet命令
，为什么，因为Powershell支持对象，对象可以包罗万象. 

```ps1
PS C:\test> $item=Get-ChildItem .
PS C:\test> $item

    Directory: C:\test

Mode                LastWriteTime     Length Name
----                -------------     ------ ----
d----        2011/11/23     17:25            ABC
-a---        2011/11/24     20:04      26384 a.txt
-a---        2011/11/23     17:25          0 b.txt
-a---        2011/11/23     17:25          0 c.txt

PS C:\test> $result=3000*(1/12+0.0075)
PS C:\test> $result
272.5
```

一些常用的方法.

多变量同时赋值, 就是连等号了

```ps1
PS C:\test> $a=$b=$c=123
PS C:\test> $a
123
PS C:\test> $b
123
PS C:\test> $c
123
```

交换变量的值, 这语法糖很方便, 不过感觉没什么实用性

```ps1
PS C:\test> $value1=10
PS C:\test> $value2=20
PS C:\test> $value1,$value2=$value2,$value1
PS C:\test> $value1
20
PS C:\test> $value2
10
```

## 2. 变量作用域与专用命令

Powershell将变量的相关信息的记录存放在名为`variable:`的驱动中. 如果要查看所有定义的变量，可以直接遍历`ls variable:`

初步猜测这应该是个bash shell中环境变量地位相似, 不过是只对当前终端/脚本所在作用域有效.

查找某个变量时, 可以使用`ls variable:前缀*`通配符方式.

验证变量是否存在, 使用`Test-Path`命令. 为什么? 因为powershell把variable当成驱动器了.

```ps1
PS C:\> test-path $1
False
PS C:\> test-path 1
False
PS C:\> test-path variable:1
True
PS C:\> echo $1
123
```

所以还是加上`variable`域才可以

关于删除变量, 因为变量会在powershell退出或关闭时，自动清除. 一般没必要删除，但是你非得删除，也可以象删除文件那样删除它. 

```ps1
PS C:\test> Test-Path variable:value1
True
PS C:\test> del variable:value1
PS C:\test> Test-Path variable:value1
False
```

为了管理`variable`域中的变量，powershell提供了五个专门管理变量的命令

- Clear-Variable

- Get-Variable

- New-Variable

- Remove-Variable

- Set-Variable

因为虚拟驱动器`variable:`的存在，`clear`，`remove`，`set`打头的命令可以被代替. 但是`Get-Variable`，`New-Variable`却非常有用. 

`New-variable`可以在定义变量时，指定变量的一些其它属性，比如访问权限. 同样`Get-Variable`也可以获取这些附加信息.

变量写保护

可以使用`New-Variable`的option选项 在创建变量时，给变量加上只读属性，这样就不能给变量重新赋值了. 

```ps1
PS C:\test> New-Variable num -Value 100 -Force -Option readonly
PS C:\test> $num=101
Cannot overwrite variable num because it is read-only or constant.
At line:1 char:5
+ $num <<<< =101     + CategoryInfo          : WriteError: (num:String) [], SessionStateUnauthorizedAccessException     + FullyQualifiedErrorId : VariableNotWritable PS C:\test> del Variable:num
Remove-Item : Cannot remove variable num because it is constant or read-only. If the variable is read-only,
ration again specifying the Force option.
At line:1 char:4
+ del <<<<  Variable:num
    + CategoryInfo          : WriteError: (num:String) [Remove-Item], SessionStateUnauthorizedAccessExcepti
    + FullyQualifiedErrorId : VariableNotRemovable,Microsoft.PowerShell.Commands.RemoveItemCommand
```

但是可以通过删除变量，再重新创建变量更新变量内容. 

```ps1
PS C:\test> del Variable:num -Force
PS C:\test> $num=101
PS C:\test> $num
101
```

有没有权限更高的变量，有，那就是：选项Constant，常量一旦声明，不可修改

```ps1
PS C:\test> new-variable num -Value "strong" -Option constant

PS C:\test> $num="why? can not delete it."
Cannot overwrite variable num because it is read-only or constant.
At line:1 char:5
+ $num <<<< ="why? can not delete it."     + CategoryInfo          : WriteError: (num:String) [], SessionStateUnauthorizedAccessException     + FullyQualifiedErrorId : VariableNotWritable PS C:\test> del Variable:num -Force
Remove-Item : Cannot remove variable num because it is constant or read-only. If the variable is read-only,
ration again specifying the Force option.
At line:1 char:4
+ del <<<<  Variable:num -Force
    + CategoryInfo          : WriteError: (num:String) [Remove-Item], SessionStateUnauthorizedAccessExcepti
    + FullyQualifiedErrorId : VariableNotRemovable,Microsoft.PowerShell.Commands.RemoveItemCommand
```

变量描述

在New-Variable 可以通过-description 添加变量描述，但是变量描述默认不会显示，可以通过Format-List 查看. 

```ps1
PS C:\test> new-variable name -Value "me" -Description "This is my name"
PS C:\test> ls Variable:name | fl *

PSPath        : Microsoft.PowerShell.CoreVariable::name
PSDrive       : Variable
PSProvider    : Microsoft.PowerShell.CoreVariable
PSIsContainer : False
Name          : name
Description   : This is my name
Value         : me
Visibility    : Public
Module        :
ModuleName    :
Options       : None
Attributes    : {}
```

创建序列

```ps1
PS C:\Users\general> 1..10 | % {echo $_}
1
2
3
4
5
6
7
8
9
10
```

但仅限于数字类型, 你没有办法使用`[a..z]`或`[a-z]`创建字母形式的序列. 替代方法是使用类型转换, 将数字转换成ASCII对应的字符.

```ps1
## 这是小写字母
97..122 | %{[char]$_}
## 这是大写字母
65..90 | %{[char]$_}
```

创建大写字母为名称的目录.(貌似必须要使用`$a`赋值一下, 否则`[char]`相当于字符串拼接了)

```ps1
65..90 | %{ $a = [char]$_; mkdir $a }
```

## 2. 变量类型
