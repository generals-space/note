# iptables 应用情景

## 1. 规则管理

参考文章

[iptables删除指定某条规则](http://www.111cn.net/sys/linux/58445.htm)

[iptables详解](http://blog.chinaunix.net/uid-26495963-id-3279216.html)

### 1.1 删除指定规则

首先要查看规则, 获取其行号

```
$ iptables [-t nat] -vnL --line-number
```

- -t 选择table类型, 默认显示filter, 可以显示其他如nat类型的

- -L 不用说了

- -v 输出详细信息, 包括通过该行所示规则的数据包数量, 总字节数及相应的网络接口

- -n 不对IP地址进行反查, 加上这个参数显示速度会快很多, 否则iptables会把规则中它已知的IP都替换成域名的

- --line-number 显示规则的序列号，这个参数的输出在删除或修改规则时会用到

然后删除指定行号代表的规则.

```
$ iptables [-t nat] -D INPUT 行号
```

- -D 删除指定规则, 后面需要接chain名与行号参

------

与之操作相似的有

- -I 链名 n: 插入，把当前规则插入为第n条。

```
## 插入为INPUT链的第三条
$ iptables -I INPUT 3 目标规则
```

- -R 链名 n: Replace替换/修改第n条规则

```
## 将INPUT的第3条规则修改为如下
$ iptables -R INPUT 3 目标规则
```

- -D 链名 n: 删除，明确指定删除目标链上的第n条规则

## 2. iptables实现端口转发(端口映射)

环境描述

A: 作为转发服务器, 作为相当于路由, 开放端口8080

B: 后端实际服务器, 开放端口80

客户端通过访问A的8080端口, 获取B的80端口的响应.

以下操作都是在A中执行

首先要确认开启Linux的数据转发

```
$ sysctl -a | grep ip_forward
net.ipv4.ip_forward = 0
```

如果`net.ipv4.ip_forward`的值为1, 则不必修改. 如果为零, 则需要执行如下操作

```
$ echo 'net.ipv4.ip_forward=1' >> /etc/sysctl.conf
$ sysctl -p
```

然后设置iptables的端口映射

```
## 将来自客户端的, 目标是A服务器的8080端口的请求, 重写为访问到B服务器的80端口的请求
$ iptables -t nat -A PREROUTING -p tcp -m tcp --dst A的IP地址 --dport 8080 -j DNAT --to B的IP地址:80
## 将由A服务器转发出去的目标是B服务器80端口的请求, 添加MASQUERADE标记
$ iptables -t nat -A POSTROUTING -p tcp -m tcp --dst B的IP地址 --dport 80 -j MASQUERADE
## 然后要确认B服务器上的80端口上确实有ACCEPT规则
```

- -p protocol

- -m match, 匹配

- -d/--dst 请求的目标地址

- -s/--src 请求的来源地址(POSTROUTING链上的规则添加-s选项可以用来伪装<IP></IP>)

- --dport/--sport 请求的目标/来源端口

建议保存设置并重启服务, 保存设置执行`service iptables save`命令即可.

### 扩展

#### CentOS7的iptables安装问题

CentOS7下的防火墙默认是`firewalld`, 可能没有安装`iptables`, 使用`iptables`管理网络需要先停止`firewalld`服务, 启动`iptables`服务.

```
$ ps aux | grep firewalld
root      16286  9.2  2.3 323580 23308 ?        Ssl  22:59   0:01 /usr/bin/python -Es /usr/sbin/firewalld --nofork --nopid
root      16744  0.0  0.0 112644   952 pts/10   R+   22:59   0:00 grep --color=auto firewalld
$ systemctl stop firewalld

$ yum install iptables
$ systemctl start iptables
```

需要注意的一点是, iptables的`save`命令依然需要通过`service`管理, 使用如下命令都的错误的.

```
$ systemctl save iptables
$ systemctl iptables save
```

#### 为所有来源添加`MASQUERADE`标记

```
$ iptables -t nat -A POSTROUTING -s 0/0 -j MASQUERADE
```

## 3. 链管理

### 3.1 修改指定链默认规则

```
## iptables [-t table名] -P chain名 规则(一般为DROP, ACCEPT等)
$ iptables -P INPUT DROP
```

### 3.2 查看指定链规则

```
$ iptables [-t {nat|filter}] --list-rules 链名 
```

说明: 不必添加`-L`选项, 可以查看目标链中的规则与子链名称(但不可查看子链下的规则). 若不指定链名, 将打印出当前表中所有规则.

示例

```
$ iptables --list-rules INPUT

```

### 3.3 添加自定义链

练习之前首先清空已经存在的规则. **注意: 正式场景下不可进行如下操作**

```
## 清空filter表
$ iptables -F
$ iptables -X
$ iptables -Z
$ iptables -L
Chain INPUT (policy ACCEPT)
target     prot opt source               destination         

Chain FORWARD (policy ACCEPT)
target     prot opt source               destination         

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination  
## 清空nat表       
$ iptables -t nat -F
$ iptables -t nat -X
$ iptables -t nat -Z
$ iptables -t nat -L
Chain PREROUTING (policy ACCEPT)
target     prot opt source               destination         

Chain INPUT (policy ACCEPT)
target     prot opt source               destination         

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         

Chain POSTROUTING (policy ACCEPT)
target     prot opt source               destination       
```

然后进行实际操作, 尝试创建, 修改, 删除自定义链.

```
## 新建一条空链docker(默认在filter表上)
$ iptables -N docker
$ iptables -L
Chain INPUT (policy ACCEPT)
target     prot opt source               destination         

Chain FORWARD (policy ACCEPT)
target     prot opt source               destination         

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         

Chain docker (0 references)
target     prot opt source               destination

## 修改链名称, 将docker修改为vpn
$ iptables -E docker vpn
$ iptables -L
Chain INPUT (policy ACCEPT)
target     prot opt source               destination         

Chain FORWARD (policy ACCEPT)
target     prot opt source               destination         

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         

Chain vpn (0 references)
target     prot opt source               destination   

## 删除空链vpn, 如果自定义链不为空, 则只能先将其清空, 无法直接删除
$ iptables -X vpn
$ iptables -L
Chain INPUT (policy ACCEPT)
target     prot opt source               destination         

Chain FORWARD (policy ACCEPT)
target     prot opt source               destination         

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination    
```

单纯创建空链是没有意义的, 自定义链的意义在于, 更加结构化地描述, 管理访问规则. 比如安装docker时, docker服务将创建属于它自己的iptables链, 所有的规则写在自定义链里. 卸载docker时, 只要清空docker链, 然后删除空链即可.

为了能使自定义链生效, 我们还需要将自定义链挂到某一默认链上, 否则网络请求是不会流经我们的自定义链的.

```
## 创建docker链
$ iptables -N docker
$ iptables -L
Chain INPUT (policy ACCEPT)
target     prot opt source               destination         

Chain FORWARD (policy ACCEPT)
target     prot opt source               destination         

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         

Chain docker (0 references)
target     prot opt source               destination   
## 将docker链挂到INPUT链下, 这里可以使用`-t`选项指定其他表, 也可以使用与`-A`同级的`-I`选项, 指定其他链名. `-j`选项不再是原来的`ACCEPT`或是`DROP`了, 它的值就是我们的自定义链名
$ iptables -A INPUT -j docker
## INPUT链下有了target列为docker的行, 并且docker链后面显示了`1 references`(原来是0)      
$ iptables -L
Chain INPUT (policy ACCEPT)
target     prot opt source               destination         
docker     all  --  anywhere             anywhere            

Chain FORWARD (policy ACCEPT)
target     prot opt source               destination         

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         

Chain docker (1 references)
target     prot opt source               destination       
```

然后尝试在docker链中添加一条规则检查是否生效. 当前所有链都为空, 默认不会存在端口屏蔽的情况. 在iptables主机上执行如下命令, 监听5000端口.

```
$ nc -l 0.0.0.0 5000
```

然后在其他主机上执行如下命令连接iptables主机的5000端口, 原则上是能正常连接并且之后可以通信的.

```
$ nc iptables主机的IP 5000
```

现在我们在docker链上屏蔽5000端口, 观察是否还能从其他主机上连接进来

```
$ iptables -A docker -p tcp --dport 5000 -j DROP
$ iptables -L
Chain INPUT (policy ACCEPT)
target     prot opt source               destination         
docker     all  --  anywhere             anywhere            

Chain FORWARD (policy ACCEPT)
target     prot opt source               destination         

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         

Chain docker (1 references)
target     prot opt source               destination         
DROP       tcp  --  anywhere             anywhere             tcp dpt:commplex-main
```

再次执行上述nc命令, 会发现无法再从其他主机上连接到iptables主机的5000端口(连接时显示timeout, 因为iptables主机上的drop规则会悄悄丢弃请求包, 也不会明确拒绝)

然后清空docker链并删除它

```
## docker链不为空时删除会报错
$ iptables -X docker
iptables: Too many links.
## -F选项不加链名的话会清空当前表中所有链的规则
$ iptables -F docker
## 还要记得docker挂在了INPUT链上, 所以还要将其从INPUT移除
$ iptables -D INPUT 1
$ iptables -L
Chain INPUT (policy ACCEPT)
target     prot opt source               destination         
docker     all  --  anywhere             anywhere            

Chain FORWARD (policy ACCEPT)
target     prot opt source               destination         

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         

Chain docker (1 references)
target     prot opt source               destination         
## 删除链
$ iptable -X docker
Chain INPUT (policy ACCEPT)
target     prot opt source               destination         
DROP       tcp  --  anywhere             anywhere             tcp dpt:terabase

Chain FORWARD (policy ACCEPT)
target     prot opt source               destination         

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination  
```