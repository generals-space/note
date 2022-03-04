# VSCode-总是提示输入git用户名密码

参考文章

1. [vscode：解决操作git总让输入用户名及密码问题](https://www.cnblogs.com/finalanddistance/p/10476080.html)
2. [Git凭证存储（简单易懂，一学就会，认真看）](https://www.cnblogs.com/volnet/p/git-credentials.html)

参考文章1与我的问题一样, 但是无效.

参考文章2给出了`credential.helper`的涵义与作用, 可以了解一下git的认证机制.

我的场景是, 使用vscode远程连接到linux服务器上做开发, commit后在vscode中点push, 才会弹出用户名和密码窗口. 我在工作机和本地都执行了`git config --global credential.helper store`命令, 是不是因为工程和密码存储是分开的, 所以操作才无效?
