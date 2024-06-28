参考文章

1. [SecureCRT修改全局外观及连接linux乱码的办法](https://www.cnblogs.com/lisn666/p/11882312.html)
    - 全局外观可以看一下
2. [SecureCRT中文乱码解决方案](https://blog.csdn.net/yongqi_wang/article/details/81392638)
3. [securecrt 8之后版本， new host key 取消显示](https://blog.csdn.net/warcraftzhaochen/article/details/73867385)
    - 自动接受服务器公钥

## 默认外观设置

外观设置比如字体, 颜色, 闪烁的分隔符等, 在`选项` -> `全局选项`中不存在, 只在`选项` -> `会话选项`中有, 如果要修改全局的默认外观, 需要按照如下步骤.

选项 -> 全局选项 -> 左侧常规 -> 预设的会话设置 -> 右侧 编辑预设的设置

## 选定内容的单词分隔符

` '",;@/\[]{}()`: 分别是 空格, 单引号, 双引号, 逗号, 分号, at符, 斜杠, 反斜杠, 大中小括号.

## 中文乱码

这个在`会话选项`中.

左侧 终端 -> 外观 -> 右侧 字符编码.

## 自动接受服务器公钥

初始连接某个服务器, 需要用户手动点击"Accept and Save"按钮, 我希望ta可以自动接受, 不用每次都点.

选项 -> 全局选项 -> 左侧常规 -> 配置文件路径, 打开对应的配置文件的目录, 编辑`SSH2.ini`文件, 修改`Automatically Accept Host Keys`的值, 由`00000000`改为`00000001`

```ini
D:"Automatically Accept Host Keys"=00000000
D:"Automatically Accept Host Keys"=00000001
```
