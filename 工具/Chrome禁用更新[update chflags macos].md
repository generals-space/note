参考文章

1. [如何禁用 Google Chrome 自动更新 (macOS, Linux, Windows)](https://zhuanlan.zhihu.com/p/349454190?utm_id=0)
2. [彻底关闭Chrome浏览器自动更新](https://blog.csdn.net/weixin_37858453/article/details/126600461)
    - windows

~~Mac版直接删除更新程序, 有效~~

```
cd ~/Library/Google/GoogleSoftwareUpdate/
rm -rf GoogleSoftwareUpdate.bundle
```

新版的 Chrome 还需要禁用该目录的写入功能

```
sudo chflags schg ~/Library/Google/GoogleSoftwareUpdate
```

恢复命令: `sudo chflags noschg ~/Library/Google/GoogleSoftwareUpdate`

> `chflags`跟目录权限貌似没有关系, 变更后还是755.

```console
~/Library/Google$ ll
total 0
drwxr-xr-x   2 general  staff    64  5 15 14:10 GoogleSoftwareUpdate
~/Library/Google$ sudo chflags schg ~/Library/Google/GoogleSoftwareUpdate
Password:
~/Library/Google$ ll
total 0
drwxr-xr-x   2 general  staff    64  5 15 14:10 GoogleSoftwareUpdate
```

然后在"关于"页面, 就会显示

```
Google Chrome可能无法进行自动更新。
```
