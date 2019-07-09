# git config代理设置

<!tags!>: <!代理!>

参考文章

[git代理设置方法解决](http://www.cnblogs.com/jackyshan/p/5985590.html)

设置http及https代理

```
git config --global http.proxy http://127.0.0.1:1080
git config --global https.proxy https://127.0.0.1:1080
```

设置socks代理

```
git config --global http.proxy 'socks5://127.0.0.1:1080'
git config --global https.proxy 'socks5://127.0.0.1:1080'
```

设置了socks5的代理后, `git clone`操作无论是`https://`这种http协议, 还是`git@github`这种ssh协议, 都可以走代理. 不太清楚使用http层面的代理时, 对ssh协议是否有效.
