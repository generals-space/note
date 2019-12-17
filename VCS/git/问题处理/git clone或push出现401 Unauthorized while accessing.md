# git clone或push出现401 Unauthorized while accessing

`git clone`或`git push`操作时报错如下

```
git push origin
error: The requested URL returned error: 401 Unauthorized while accessing https://git.oschina.net/generals-space/ansible.git/info/refs

fatal: HTTP request failed
```

问题分析

git版本问题, 当前版本为`1.7.1`

解决办法

重新安装高版本的git即可, 最好是1.9+