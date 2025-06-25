# git pull失败-Fatal：could not read Username for ''：No such device or address[jenkins]

1. [Fatal: could not read Username for ‘https://github.com’: No such device or address](https://community.jenkins.io/t/fatal-could-not-read-username-for-https-github-com-no-such-device-or-address/11254)

## 问题描述

jenkins CI流水线配置了git账户, 但流程里需要从两个仓库拉代码(比如github+gitlab), git账户只适配了其中一个, 拉取另一个时出了问题.

```log
fatal: could not read Username for 'https://gitlab.com': No such device or address
```

> 代码目录已经在本地了, CI脚本里只执行`git pull`而非`git clone`, 不过在这个目录下手动执行 git pull 是可以拉下来的.

## 原因分析

按照参考文章1中所说, 这个报错其实是git在让用户手动输入用户名密码, 但是CI流程又是非交互式的, 所以直接报错了.

## 解决方法(临时)

由于两个仓库都是内部库, 所以在`git pull`拉取第2个仓库的代码时直接在url指定用户名密码.

```
git pull https://test:1qaz%21QAZ@gitlab.com/xxx/test.git
```
