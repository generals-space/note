# git config免密码操作[gitconfig git-credentials]

参考文章

1. [linux 服务器git clone 无需输入名字和密码操作](https://blog.csdn.net/my__blog/article/details/123852454)

git 操作不使用密码, 是因为在本地有存储, 就在`~/.git-credentials`

```
https://账号:密码@gitee.com
https://账号:密码@github.com
```

同时在`~/.gitconfig`中有指定认证类型为本地存储

```ini
[credential]
	helper = store
```

二者缺一不可.

## 实现方法

两个文件都可以手动编辑.

其中, `.gitconfig`中的配置可以通过`git config`命令实现

```ini
$ cat .gitconfig
[user]
	name = general
	email = generals.space@gmail.com
$ git config --global credential.helper store
$ cat .gitconfig
[user]
	name = general
	email = generals.space@gmail.com
[credential]
	helper = store
```

然后`git clone`的时候还会提示输入账号密码, 不过只输入一次就可以了.

也可以自行创建`~/.git-credentials`文件.
