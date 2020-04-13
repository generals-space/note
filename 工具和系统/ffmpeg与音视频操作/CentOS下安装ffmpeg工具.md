参考文章

1. [CentOS中yum安装ffmpeg](https://www.cnblogs.com/wpjamer/p/ffmpeg.html)

2. [Linux有问必答:如何在CentOS或者RHEL上启用Nux Dextop仓库](https://linux.cn/article-3889-1.html)

3. [Nux Dextop官方文档](http://li.nux.ro/repos.html)

`ffmpeg`没有在CentOS的官方源中, 要下载它需要使用第三方源`Nux Dextop`.

参考文章2中对`Nux Dextop`这个源的介绍比较详细, 可以看一下.

首先安装`repo`的rpm文件.

CentOS6

```
$ yum -y install epel-release && rpm -Uvh http://li.nux.ro/download/nux/dextop/el6/x86_64/nux-dextop-release-0-2.el6.nux.noarch.rpm
```

CentOS7

```
$ yum -y install epel-release && rpm -Uvh http://li.nux.ro/download/nux/dextop/el7/x86_64/nux-dextop-release-0-5.el7.nux.noarch.rpm
```

安装好后`/etc/yum.repos.d`目录下会出现`nux-dextop.repo`文件. 其内容为

```ini
[nux-dextop]
name=Nux.Ro RPMs for general desktop use
baseurl=http://li.nux.ro/download/nux/dextop/el6/$basearch/ http://mirror.li.nux.ro/li.nux.ro/nux/dextop/el6/$basearch/
enabled=1
gpgcheck=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-nux.ro
protect=0

[nux-dextop-testing]
name=Nux.Ro RPMs for general desktop use - testing
baseurl=http://li.nux.ro/download/nux/dextop-testing/el6/$basearch/
enabled=0
gpgcheck=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-nux.ro
protect=0
```

然后直接yum install就好了.

```
$ yum install ffmpeg ffmpeg-devel -y
```