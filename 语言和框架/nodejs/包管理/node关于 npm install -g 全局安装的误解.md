# node关于 npm install -g 全局安装的误解

参考文章

1. [node关于 npm -g 全局安装的误解](https://www.jianshu.com/p/3e34f29dffba)
2. [Where does npm install packages?](https://stackoverflow.com/questions/5926672/where-does-npm-install-packages)

`npm install -g`安装的包, 虽然叫全局包, 但是并不是那种, 所有工程都能引入的全局包. 而是会将二进制放到`/usr/bin/`下, 用来直接执行.
