# Windows开启sshd服务

参考文章

1. [win10 开启ssh server服务 远程登录](https://blog.csdn.net/weixin_43064185/article/details/90080815)
2. [OpenSSH Key Management](https://docs.microsoft.com/en-us/windows-server/administration/openssh/openssh_keymanagement)
3. [Windows 10 v1803 OpenSSH Server Key File Authentication](https://www.reddit.com/r/Windows10/comments/8quejo/windows_10_v1803_openssh_server_key_file/)

想着能无密码远程登录win10的, 创建了`~/.ssh/authorized_keys`文件, 结果不行.

以为是文件权限问题, 于是找到了参考文章2, 又不行.

后来找到了参考文章3, 开启sshd服务的密钥认证机制, 重启服务, 结果还是不行.

fuck, 放弃了.

------

注意在使用scp命令时, windows端的路径分隔符为正斜线`/`.

```
scp gener@192.168.0.8:C:/Users/gener/Downloads/readme.md /root/
```
