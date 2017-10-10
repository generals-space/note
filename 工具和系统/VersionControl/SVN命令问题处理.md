# SVN命令问题处理

## 1.

```
$ svn commit -m '第一次发布'
svn: E000013: Commit failed (details follow):
svn: E000013: Can't open file '/home/svn/ibase4j/db/txn-current-lock': Permission denied
```

服务端权限配置完全正确, 也可以checkout, 但是提交时却提示权限拒绝.

可能原因: 客户端使用多个svn帐户进行验证, 由于客户端的验证缓存机制, 在checkout的时候没有重新输入验证, 所以当前版本库也许不是以你想要的用户checkout得到的.

## 2. 

```
$ svn commit -m 你的提交说明
svn: E155010: Commit failed (details follow):
svn: E155010: '某文件' is scheduled for addition, but is missing
```

这是因为这个文件最近新加入版本库, 但是在提交之前就从文件系统中移除了, 然后提交时svn还是会寻找这个文件准备将其提交但是已经找不到了.

...通常是不小心把一些不该加入版本库的文件执行了`add`, 又没有正常移除.

解决办法是执行`svn revert 目标文件`, 此文件就从版本为库中移除了, 然后就可以执行提交了.

这也是误添加文件到版本库的解救方式.

> 注意: 如果是文件夹, 需要先revert子文件.

## 3. svn update 树冲突 (local unversioned, incoming add upon update)

```
$ svn update
D     C Runtime
      >   local unversioned, incoming add upon update
```

原因分析

这是命令行下的 svn 树冲突, 文件本身没有改变，只是本地版本库里面出现冲突.

解决方法如下

```
$ svn resolve --accept working ./Runtime
Resolved conflicted state of 'Runtime'
$ svn revert ./Runtime
Reverted 'Runtime'
$ svn status
```

移除本地svn版本库里面的冲突信息, 其中`resolve`子命令会将冲突的文件/目录从本地删除, 而`revert`则会将其找回, 找回后执行`status`就没有冲突问题了.