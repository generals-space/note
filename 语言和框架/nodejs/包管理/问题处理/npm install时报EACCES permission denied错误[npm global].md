# npm install时报EACCES: permission denied错误

参考文章

1. [npm ERR! Error: EACCES: permission denied, access](https://blog.csdn.net/testcs_dn/article/details/78869419)
2. [npm install报错 Error: EACCES: permission denied ](https://www.cnblogs.com/yechen2019/p/12331715.html)
    - 配置nodejs的npm-global全局path

`npm install`安装package.json文件中的依赖时报这个错误, 当时屏幕闪得太快没看到, 最下方提示查看`~/.npm`目录下的日志, 没得到什么有用的信息. 胡乱找了好久才想起来往上翻.

```
> node-sass@4.9.3 install /home/ksp-web/node_modules/node-sass
> node scripts/install.js

Unable to save binary /home/ksp-web/node_modules/node-sass/vendor/linux-x64-57 : { Error: EACCES: permission denied, mkdir '/home/ksp-web/node_modules/node-sass/vendor'
    at Object.fs.mkdirSync (fs.js:885:18)
    at sync (/home/ksp-web/node_modules/mkdirp/index.js:74:13)
    at Function.sync (/home/ksp-web/node_modules/mkdirp/index.js:80:24)
    at checkAndDownloadBinary (/home/ksp-web/node_modules/node-sass/scripts/install.js:114:11)
    at Object.<anonymous> (/home/ksp-web/node_modules/node-sass/scripts/install.js:157:1)
    at Module._compile (module.js:653:30)
    at Object.Module._extensions..js (module.js:664:10)
    at Module.load (module.js:566:32)
    at tryModuleLoad (module.js:506:12)
    at Function.Module._load (module.js:498:3)
  errno: -13,
  code: 'EACCES',
  syscall: 'mkdir',
  path: '/home/ksp-web/node_modules/node-sass/vendor' }

> node-sass@4.9.3 postinstall /home/ksp-web/node_modules/node-sass
> node scripts/build.js
```

解决方法如下.

在命令前加上 `sudo`

`sudo npm install`

不过这样子可能还是不行, 你需要这样

`sudo npm install --unsafe-perm=true`

或许你还是会遇到错误, 请尝试这样

`sudo npm install --unsafe-perm=true --allow-root`