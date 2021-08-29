# Python2-os.path路径操作

`__file__`可以得到当前python文件的文件名. 如果执行时指定的是相对路径则只能得到相对路径, 如果指定了绝对路径则可以得到绝对路径.

```py
#!/usr/bin/env python
## test.py
import os
print(__file__)
```

```console
$ ls
test.py
$ pwd
/tmp
$ python test.py 
test.py                 ## 执行时为相对路径, 输出时也是相对路径
$ python /tmp/test.py 
/tmp/test.py            ## 执行时为绝对路径, 输出时也是绝对路径
```

由于执行python脚本时不一定是在命令行下明确地指定其路径, 也有可能是在脚本中相互调用, 所以直接通过`__file__`可能无法获取到目标脚本的正确路径. 因此实际使用中一般与`os.path.dirname()`与`os.path.abspath()`等配合使用.

## 路径字符串解析

`os.path.basename(path)`: 可以得到`path`所表示的文件名, 其实就是对`path`中最后一个斜线`/`进行分割, 返回斜线后面的字符串, 如果以`/`结尾则会得到''.

```py
>>> import os
>>> os.path.basename('/etc/yum.conf')
'yum.conf'
>>> os.path.basename('/etc/')
''
>>> os.path.basename('/etc')
'etc'
```

`os.path.dirname(path)`: 与`os.path.basename()`相似, 可以得到`path`中的目录路径, 即对`path`中最后一个斜线进行分割, 取其前面的内容.

`os.path.split(path)`: 像是`basename()`与`dirname()`的合体, 返回`path`路径表示的目录与文件的二元组.

`os.path.abspath(path)`: 如果`path`是绝对路径, 则直接返回`path`; 如果`path`是相对路径, 则`abspath`会将当前路径与`path`进行拼接, 得到一个绝对路径. 

```py
>>> import os
>>> os.path.abspath('/etc')
'/etc'
>>> os.path.abspath('')
'/root'
>>> os.path.abspath('etc')
'/root/etc'
>>> 
```

当前路径应该是通过代码中的上下文计算的, 可以通过`os.path.abspath(path)`查看. 

一般来说, 在python交互终端得到的当前路径, 应该是当前用户的home目录, 而在代码中则可以得到当前python文件所在的目录. 

所以python中的常用用法`os.path.abspath(os.path.dirname(__file__))`, 得到的是当前文件 **所在目录**的绝对路径. 

## 路径字符串拼接

`os.path.join(path1[, path2[, ...]])`: 多个路径组合后返回，第一个绝对路径之前的参数将被忽略, 并且重复路径将不进行拼接. 

```py
>>> os.path.join('c:\\', 'csv', 'test.csv') 
'c:\\csv\\test.csv' 
>>> os.path.join('windows\temp', 'c:\\', 'csv', 'test.csv') 
'c:\\csv\\test.csv' 
>>> os.path.join('/home/aa','/home/aa/bb','/home/aa/bb/c') 
'/home/aa/bb/c' 
```
