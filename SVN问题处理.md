# SVN问题处理

```
$ svn commit -m '第一次发布'
svn: E000013: Commit failed (details follow):
svn: E000013: Can't open file '/home/svn/ibase4j/db/txn-current-lock': Permission denied
```

服务端权限配置完全正确, 也可以checkout, 但是提交时却提示权限拒绝.

可能原因: 客户端使用多个svn帐户进行验证, 由于客户端的验证缓存机制, 在checkout的时候没有重新输入验证, 所以当前版本库也许不是以你想要的用户checkout得到的.