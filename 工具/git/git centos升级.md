参考文章

1. [CentOS 升级git](https://blog.csdn.net/qq_21127151/article/details/125454394)


wget https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.24.0.tar.gz

yum install -y gcc curl-devel expat-devel gettext-devel openssl-devel zlib-devel

yum remove git
git --version

tar -zxf git-2.24.0.tar.gz
cd git-2.24.0
make prefix=/usr/local/git all
make prefix=/usr/local/git install

vim /etc/profile
export PATH=$PATH:/usr/local/git/bin
source /etc/profile

