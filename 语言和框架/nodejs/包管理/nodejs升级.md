# nodejs升级

无需编译新版源码包, 也无需重新`apt`/`yum`覆盖安装. node本身有一个`n`模块(名字真够短的), 用于管理NodeJS版本.

安装方法

```shell
sudo npm install -g n
```

之后系统中会出现`n`命令, 不过没有man手册, 使用`n --help`可以查看使用帮助, 用法很明了.
