# Python内存泄漏

参考文章

1. [使用 Gc、Objgraph 干掉 Python 内存泄露与循环引用！](https://www.douban.com/note/645539928/)

2. [python 内存泄露的诊断](http://rstevens.iteye.com/blog/828565)

3. [Python内存泄漏问题查找](http://blog.csdn.net/i2cbus/article/details/20155273)



[root@192-168-169-51 docker]# docker login reg01.sky-mobi.com
Username: jiale.huang
Password: 
Login Succeeded



[root@192-168-169-51 docker]# docker push reg01.sky-mobi.com/jxtz/jxtz:1.5.0
The push refers to a repository [reg01.sky-mobi.com/jxtz/jxtz]
9192c1018048: Pushed 
b186c0cb116a: Mounted from base/jdk 
f4c30d29e3b0: Mounted from base/jdk 
b2a6f775b6e6: Mounted from base/jdk 
bc6573e57f55: Mounted from base/jdk 
cd51912c80c4: Mounted from base/jdk 
1.5.0: digest: sha256:ac4723521996cd872d12db49d8bdd8182fa61b200af604b6fad9143338496082 size: 1580




[root@192-168-169-51 docker]# docker push reg01.sky-mobi.com/jxtz/jxtz:1.5.1
The push refers to a repository [reg01.sky-mobi.com/jxtz/jxtz]
ecea0352eea1: Pushed 
b186c0cb116a: Layer already exists 
f4c30d29e3b0: Layer already exists 
b2a6f775b6e6: Layer already exists 
bc6573e57f55: Layer already exists 
cd51912c80c4: Layer already exists 
1.5.1: digest: sha256:d1e395f63aca95a713eaaed9164b1969b1ca4a89a9f941232cfa4ffa9024a87d size: 1580


```
[root@192-168-169-51 docker]# docker images
REPOSITORY                                  TAG                 IMAGE ID            CREATED             SIZE
reg01.sky-mobi.com/base/centos              6.0.0               2f7d7dd78e6f        4 months ago        462 MB
reg01.sky-mobi.com/base/jdk                 6.0.0               effa78a1faab        4 months ago        762.2 MB
reg01.sky-mobi.com/jxtz/jxtz                1.5.1               944fb3fbc2bc        10 weeks ago        791.1 MB
reg01.sky-mobi.com/jxtz/jxtz                1.5.0               a4995cacefb4        11 weeks ago        791.1 MB
reg01.sky-mobi.com/vpn-reset-service/pptp   1.0.0.0             85740a062835        4 months ago        720 MB
```


$ docker images
REPOSITORY                                           TAG                 IMAGE ID            CREATED             SIZE
daocloud.io/ubuntu                                   14.04               c69811d4e993        3 months ago        188 MB
$ docker tag c69811d4e993 reg01.sky-mobi.com/origin/ubuntu:14.04
$ docker images
REPOSITORY                                           TAG                 IMAGE ID            CREATED             SIZE
daocloud.io/ubuntu                                   14.04               c69811d4e993        3 months ago        188 MB
reg01.sky-mobi.com/origin/ubuntu                     14.04               c69811d4e993        3 months ago        188 MB



[root@220 ~]# docker push reg01.sky-mobi.com/origin/ubuntu:14.04
The push refers to a repository [reg01.sky-mobi.com/origin/ubuntu]
f668bc0d79c1: Pushed 
05e7ab001e8a: Pushed 
45dee3415f42: Pushed 
cf195c989ceb: Pushed 
826fc2344fbb: Pushing [==================>                                ] 68.48 MB/187.8 MB
826fc2344fbb: Preparing 
