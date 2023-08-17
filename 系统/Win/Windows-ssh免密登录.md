# Win10-ssh免密登录

参考文章

1. [在windows中安装OpenSSH，无密码登录，永远不断线](https://www.cnblogs.com/chengchen/p/9610819.html)
    - openssh服务端配置文件`C:\ProgramData\ssh\sshd_config`, `ProgramData`为隐藏目录.
2. [linux通过openssh无密码访问window](https://www.nginx.cn/5170.html)

操作系统: win10

从 github 下载 opensshd-x64 压缩包后, 启动服务, 可以使用密码登录, 但是不支持密钥登录, 安装目录也找不到 sshd 服务端的配置文件.

后来找到参考文章1, ta提到了 openssh 的服务端配置文件在`C:\ProgramData\ssh\sshd_config`, openssh原本不需要安装, 所以这个目录应该是在初次注册为服务时创建的.

最初我把这个文件的`PubkeyAuthentication`设置为`yes`, 重启了服务, 但是不管用.

后来找到了参考文章2, 原来`sshd_config`最下面还有两行这样的说明

```
Match Group administrators
       AuthorizedKeysFile __PROGRAMDATA__/ssh/administrators_authorized_keys
```

看起来像是判断登录请求用户是否为管理员分组, 如果没有就不生效一样. 将其注释掉好了.

------

sshd服务的重启命令(可能需要管理员权限)

```
net stop sshd
net start sshd
```
