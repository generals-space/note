
[环境安装 - 官方下载](https://www.haskell.org/platform)


[Warp : Haskell 的高性能 Web 服务器(译文)](http://veryr.com/posts/warp/)

[Warp: 一个Haskell web服务器](http://yi-programmer.com/2011-05-05_warp-a-haskell-web-server.html)

[Haskell进入生产(Hasura.io)](http://www.jdon.com/47117)

[yesod官方手册(书籍)](https://www.yesodweb.com/book/)

[stack开发技术栈](https://haskell-lang.org/get-started)

wai之于warp, 等价于wsgi之于bottle.py, 一个是规范, 一个是最小实现

```
$ yum install haskell-platform
```

haskell包管理器`cabal`命令(`cabal-install`包中, 已于`haskell-platform`一同安装)

初次执行提示需要下载官方源缓存

```
$ cabal install warp
Config file path source is default config file.
Config file /root/.cabal/config not found.
Writing default configuration to /root/.cabal/config
Warning: The package list for 'hackage.haskell.org' does not exist. Run 'cabal
update' to download it.
cabal: There is no package named 'warp'.
You may need to run 'cabal update' to get the latest list of available
packages.
```

执行`cabal update`

```
$ cabal update
Downloading the latest package list from hackage.haskell.org
```