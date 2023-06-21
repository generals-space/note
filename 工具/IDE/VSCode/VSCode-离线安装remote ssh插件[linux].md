参考文章

1. [Using "Remote SSH" in VSCode on a target machine that only allows inbound SSH connections](https://stackoverflow.com/questions/56718453/using-remote-ssh-in-vscode-on-a-target-machine-that-only-allows-inbound-ssh-co)
2. [Installing VSCode Server on Remote Machine in Private Network (Offline Installation)](https://medium.com/@debugger24/installing-vscode-server-on-remote-machine-in-private-network-offline-installation-16e51847e275)

第一次连接目标主机, 由于无法下载 vscode-server 工具, 一定会失败一次.

然后通过常规 ssh 方式, 登陆目标主机的终端, 进入`~/.vscode-server/bin`目录, 其中会有很多个hash目录, 如下

```console
$ ls ~/.vscode-server/bin
c9a2f78283b6e5ef708fb8869e2a5adaa476e42f
```

每一个目录表示一个vscode客户端, 不同版本的vscode(或是remote ssh插件??? 感觉是前者), 会建立不同hash值的目录, 因为内网服务器无法下载 vscode-server 包, 所以ta的hash目录会是空的.

```
https://update.code.visualstudio.com/commit:{COMMIT_ID}/server-linux-x64/stable
```

`{COMMIT_ID}`就是上面的hash值, 手动下载对应版本的 vscode-server 包, 放到`~/.vscode-server/bin/{COMMIT_ID}`目录下, 解压.

> 注意: `commit:{COMMIT_ID}`中, 是`:`冒号不是`/`斜杠

注意, vscode-server 的 tar 包解压后会得到一个目录, 我们需要将这个目录下的所有内容移动到 hash 目录本身下面.

```console
[root@hua-dlzx1-i1106-gyt 5763d909d5f12fe19f215cbfdd29a91c0fa9208a]# ls
bin  extensions  LICENSE  node  node_modules  out  package.json  product.json  server.sh  vscode-server-linux-x64  vscode-server-linux-x64.tar.gz
```

然后在vscode中重新连接即可.
