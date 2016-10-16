# Nmap脚本编写

原文链接

[漏洞扫描 －－ 编写Nmap脚本](http://www.myhack58.com/Article/html/3/8/2014/54252.htm)

> 2006年12月份，`Nmap4.21 ALPHA1`版加入脚本引擎，并将其作为主线代码的一部分。 `NSE`脚本库如今已经有400多个脚本，覆盖了各种不同的网络机制(从SMB漏洞检测到Stuxnet探测，及中间的一些内容)。NSE 的强大，依赖它强大的功能库，这些库可以非常容易的与主流的网络服务和协议，进行交互。

## 挑战

我们经常会扫描网络环境中的主机是否存在某种新漏洞，而扫描器引擎中没有新漏洞的检测方法，这时候我们可能需要自己开发扫描工具。

你可能已经熟悉了某种脚本(例如: `Python`，`Perl`，etc.) ，并可以快速写出检测漏洞的程序。但是，如果面临许多主机时， 针对两三个主机的检测方法，可能并不奏效。

Nmap解救你! 使用内嵌的`Lua`语言和强大的集合库，你可以结合`nmap`高效的主机和端口扫描引擎，开发出针对多数主机的检测方法。 

## 实现

Nmap引擎脚本，由**`Lua`编程语言**、**`NmapAPI`** 、**系列强大的`NSE`库**实现。

为了达到本文的目的，现假设某个应用中存在一个叫`ArcticFission`漏洞。与许多其他的 web应用程序类似，可以通过探测特定的文件，假设这个文件就是`/arcticfission.html`，用正则表达式提取文件内容中的版本号，与有漏洞的值进行对比. 听起来好像很简单，让我们开始吧! 

## 框架代码 

基于传统的语言标准，我们写一个脚本，作用: 遇到开放的 HTTP 端口，就返回`Hello World`。

```lua
-- The Head Section --

-- The Rule Section --
portrule = function(host, port)
-- port.state为open的条件是必需的, 因为如果目标端口是一个filtered的状态, 没法执行这段代码
-- 所以做实验的时候要确认目标主机的防火墙已关闭
return port.protocol == "tcp" and port.number == 80 and port.state == "open"
end

-- The Action Section --
action = function(host, port)
return "Hello world !"
end
```

> 注意: 以`--`起始的行表示注释。

NSE 脚本主要由三部分组成:

### 1. The Head Section

该部分包含一些元数据，主要描述脚本的功能，作者，影响力，类别及其他。

### 2. The Rule Section

该部分定义脚本执行的必要条件。至少包含下面列表中的一个函数:

- portrule

- hostrule

- prerule

- postrule

此案例中，重点介绍`portrule`。`portrule`能够在执行操作前，检查`host`和`port`属性。`portrule`会利用`nmap`的API检查TCP 80端口。

### 3. The Action Section

该部分定义脚本逻辑。此处案例中，检测到开放 80 端口，则打印`Hello World`。脚本的输出内容，会在`nmap`执行期间显示出来。

```
## 我们上面编写的脚本保存为http-vuln-check.nse文件, 最后一个参数为扫描对象, 可自定义
root@security:/home/offensive/nmap_nse# nmap -sS -p 22,80,443 --script /home/offensive/nmap_nse/http-vuln-check.nse www.exploit-db.com
Starting Nmap 6.47 ( http://nmap.org ) at 2014-09-29 10:39 EDT
Nmap scan report for www.exploit-db.com (192.99.12.218)
Host is up (0.47s latency).
Other addresses for www.exploit-db.com (not scanned): 198.58.102.135
rDNS record for 192.99.12.218: cloudproxy71.sucuri.net
PORT    STATE    SERVICE
22/tcp  filtered ssh
80/tcp  open     http
|_http-vuln-check: Hello world !
443/tcp open     https
```

> 注: 上面`80/tcp`的输出中, `http-vuln-check`字符串是所用脚本的名称.

## 调用脚本库

优秀的库集合，促使其变的强大。例如，可调用现有库中的函数，针对http端口创建`portrule`。此处用到了 `shortport`.

```lua
-- The Head Section --
local shortport = require "shortport"
-- The Rule Section --
portrule = shortport.http
-- The Action Section --
action = function(host, port)
return "Hello world!"
end
```

同样的扫描，产生了不同的结果

```
root@security:/home/offensive/nmap_nse# nmap -sS -p 22,80,443 --script /home/offensive/nmap_nse/http-vuln-check_shortport.nse www.exploit-db.com
Starting Nmap 6.47 ( http://nmap.org ) at 2014-09-29 10:36 EDT
Nmap scan report for www.exploit-db.com (192.99.12.218)
Host is up (0.46s latency).
Other addresses for www.exploit-db.com (not scanned): 198.58.102.135
rDNS record for 192.99.12.218: cloudproxy71.sucuri.net
PORT    STATE    SERVICE
22/tcp  filtered ssh
80/tcp  open     http
|_http-vuln-check_shortport: Hello world!
443/tcp open     https
|_http-vuln-check_shortport: Hello world!
Nmap done: 1 IP address (1 host up) scanned in 6.32 seconds
```

该脚本对443执行了类似80端口的操作。主要是因为`shortport.http`表示类似HTTP的端口(80,443,631,7080,8080,8088,5800,3872,8180,8000)，也就是说，`nmap`会探测服务`http`、`https`、`ipp`、`http-alt`、`vnc-http`、`oem-agent`、`soap`、`http-proxy`非标准端口，如果想要获取更多的信息，请查阅 shortport 的文档.

## 服务探测

让我们把注意力放到 action 部分的逻辑上。上述漏洞的检测，首先需要探测页面`/arcticfission.html`

```lua
-- The Head Section --
local shortport = require "shortport"
local http = require "http"
-- The Rule Section --
portrule = shortport.http
-- The Action Section --
action = function(host, port)
local uri = "/arcticfission.html"
local response = http.get(host, port, uri)
return response.status
end
```

上述代码用到了`http`库处理web页面

```
root@security:/home/offensive/nmap_nse# nmap -sS -p 22,80,443 --script /home/offensive/nmap_nse/http-vuln-check_shortport2.nse www.exploit-db.com
Starting Nmap 6.47 ( http://nmap.org ) at 2014-09-29 11:16 EDT
Nmap scan report for www.exploit-db.com (192.99.12.218)
Host is up (0.48s latency).
Other addresses for www.exploit-db.com (not scanned): 198.58.102.135
rDNS record for 192.99.12.218: cloudproxy71.sucuri.net
PORT    STATE    SERVICE
22/tcp  filtered ssh
80/tcp  open     http
|_http-vuln-check_shortport2: 403
443/tcp open     https
|_http-vuln-check_shortport2: 400
```

上述输出表明，两个服务器端口不存在对应页面`arcticfission.html`，注意`http`库会自动在http与https端口切换，因此你不需要考虑去实现TLS/SSL。

如果只想输出存在该页面的web应用，可以如下操作:

```lua
-- The Head Section --
local shortport = require "shortport"
local http = require "http"
-- The Rule Section --
portrule = shortport.http
-- The Action Section --
action = function(host, port)
local uri = "/arcticfission.html"
local response = http.get(host, port, uri)
if (response.status == 200) then
return response.body
end
end
```

上述代码，返回状态码为200的页面的内容。

> 注意: 如果没有数据返回或返回的页面为空，将导致无输出显示.

## 漏洞探测

许多时候，可以通过一个简单的服务版本号，探测漏洞。这种情况，假象的服务器会返回一个包含版本号的标识。

```lua
local shortport = require "shortport"
local http = require "http"
local string = require "string"
-- The Rule Section --
portrule = shortport.http
-- The Action Section --
action = function(host, port)
local uri = "/arcticfission.html"
local response = http.get(host, port, uri)
if ( response.status == 200 ) then
local title = string.match(response.body, "<[Tt][Ii][Tt][Ll][Ee][^>]*>ArcticFission([^<]*)</[Tt][Ii][Tt][Ll][Ee]>")
return title
end
end
```

上述代码，用到了string库，以便获取, 匹配页面头。

```
offensive@security:~/nmap_nse$ nmap -p 80,443 --script /home/offensive/nmap_nse/http-vuln-check_shortport4.nse 192.168.1.105
Starting Nmap 6.47 ( http://nmap.org ) at 2014-09-30 03:49 EDT
Nmap scan report for localhost (192.168.1.105)
Host is up (0.00053s latency).
PORT    STATE  SERVICE
80/tcp  open   http
|_http-vuln-check_shortport4: 1.0
443/tcp closed https
Nmap done: 1 IP address (1 host up) scanned in 0.07 seconds
```

正如上面描述的那样，现在需要将获取的值与漏洞值比较， 确认是否存在漏洞。

```lua
-- The Rule Section --
local shortport = require "shortport"
local http = require "http"
local string = require "string"
-- The Rule Section --
portrule = shortport.http
-- The Action Section --
action = function(host, port)
local uri = "/arcticfission.html"
local response = http.get(host, port, uri)
if ( response.status == 200 ) then
local title = string.match(response.body, "<[Tt][Ii][Tt][Ll][Ee][^>]*>ArcticFission ([^<]*)</[Tt][Ii][Tt][Ll][Ee]>")
if ( title == "1.0" ) then
return "Vnlnerable"
else
return "Not Vulnerable"
end
end
end
```

测试结果如下(这里与前面的被测主机不一样, 并且443端口是close状态, 所以不会有`Not Vulnerable`输出):

```
offensive@security:~/nmap_nse$ nmap -p 80,443 --script /home/offensive/nmap_nse/http-vuln-check_shortport5.nse 192.168.1.105
Starting Nmap 6.47 ( http://nmap.org ) at 2014-09-30 04:05 EDT
Nmap scan report for localhost (192.168.1.105)
Host is up (0.00045s latency).
PORT    STATE  SERVICE
80/tcp  open   http
|_http-vuln-check_shortport5: Vnlnerable
443/tcp closed https
```

版本检测的另一种方法，生成Hash与有漏洞的页面对比。为了实现此效果，此处调用了`openssl`库。

```lua
-- The Head Section --
local shortport = require "shortport"
local http = require "http"
local stdnse = require "stdnse"
local openssl = require "openssl"
-- The Rule Section --
portrule = shortport.http
-- The Action Section --
action = function(host, port)
local uri = "/arcticfission.html"
local response = http.get(host, port, uri)
if (response.status == 200) then
local vulnsha1 = "398ffad678f17a4f16ccd00b1914ca986d0b9258"
-- 比较页面内容的哈希值???
local sha1 = string.lower(stdnse.tohex(openssl.sha1(response.body)))
if ( sha1 == vulnsha1 ) then
return "Vulnerable"
else
return "Not Vulnerable"
end
end
end
```

