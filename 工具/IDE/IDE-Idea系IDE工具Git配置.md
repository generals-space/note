# Idea系IDE工具Git配置

## 1. 准备

1. 下载并安装git；

2. 在studio中设置git插件：File->Setting->Version Control->Git, 设置git可执行文件的路径

3. 点击Test测试是否能够关联成功

## 2. 初始化git项目(git init)

操作如下：

VCS->Enable Control Integration->Select "Git".

### 3. 为git添加remote

这一步，studio没有提供可视化的GUI，需要使用Git命令。

(Windows下需要使用Git Bash)将目录切换到项目的目录，输入git添加remote的命令，例如：

```
git remote add origin "https://github.com/xxx/xxx.git"
```

## 4.将新增代码添加到VCS(git add)

当前要提交的文件->VCS->Git->Add

## 5. 提交变化(git commit)

VCS->Commit Changes

在提交的时候可以选择Commit and Push,就可以直接push到服务器。

## 6. 提交变化(git push)

VCS->Git->Push.

最后，如果要是clone project到studio

VSC -> Checkout from Version Control -> Git