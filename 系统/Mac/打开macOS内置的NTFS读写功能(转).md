# 打开macOS内置的NTFS读写功能(转)

原文链接

[打开macOS内置的NTFS读写功能](https://bbs.feng.com/read-htm-tid-11830448.html)


> <!general!>: 其实在采用文章中所说方法之前, 我已经使用了原文回复中提到的[Mounty for NTFS](https://mounty.app/)工具, 还蛮好用的, 挺方便, 而且免费. 转载这篇文章只是作为一个备份.

Mac的系统只支持NTFS读, 不支持NTFS写, 据说是版权问题, 这里也不深究了. 网上的即决方案主要是两种: 

1. 将移动硬盘格式化为exFAT格式: FAT32的硬盘格式大家肯定不会考虑, 因为该格式下单个存储文件不能超过4G. exFAT格式同时支持macOS和Windows, 但该格式很不稳定, 很多网友都遇到了数据丢失问题, 硬盘寿命堪忧. 
2. 有很多软件支持开启NTFS读写这个功能. 这里也不列举了, 但某些个流氓软件删了都还能继续运行. 这种无耻的收费软件楼主是说啥都不可能再次使用的. 

其实软件支持侧面反应了macOS内部是支持NFTS读写的. 下面提供开启masOS内置支持NTFS的方案.

## 

熟悉shell的朋友直接`sudo vim /etc/fstab`, 从第3步开始就可以了

1. 打开访达, `shift+command+G`打开"前往", 进入/etc目录

![1. 打开访达, shift+command+G打开"前往", 进入/etc目录](https://gitee.com/generals-space/gitimg/raw/master/81b135a4f7352bfcbbd21cb24103f5c1.jpg)

2. 创建fstab这个文件. 由于是系统目录, 需要一定的权限才能操作. 曲线救国方法: 

   1. 复制hosts文件到桌面, 重命名为fstab
   2. 右键—>打开方式—>文本编辑
   3. 编辑fstab文件, 见第3步
   4. 保存文件, 复制到访达中, 此步可能需要输入密码. 效果如图所示

![2.创建fstab这个文件. 由于是系统目录, 需要一定的权限才能操作. 曲线救国方法:  1）复制hosts文件到桌面,  ...](https://gitee.com/generals-space/gitimg/raw/master/8b7cc3ac59b6a2bf9a03eaf19e1646bb.jpg)

3. fstab文件中, 输入

```
LABEL=XXX none ntfs rw,auto,nobrowse
```

XXX是你的移动硬盘的名字. 例如我的移动硬盘名字是Elements, 效果如图

![3.输入 LABEL=XXX none ntfs rw,auto,nobrowse XXX是你的这个移动硬盘的名字.  例如我的移动硬盘名字是Elem ...](https://gitee.com/generals-space/gitimg/raw/master/76dd338f8a906133fd62f1c2bce439d7.jpg)

4. `shift+command+G`打开"前往", 进入`/Volume`目录

![4.shift+command+G打开"前往", 进入/Volume目录](https://gitee.com/generals-space/gitimg/raw/master/9e92eb1e203e4e750430e4f5581e798d.jpg)

5. 插入你的移动硬盘, 看到这个就是成功了. 

楼主把他固定到了左边边栏, 你可以注意到, 此时硬盘不在"设备"目录下了. 以后记得每次都要安全推出了硬盘, 不然下次左边边栏就看不到这个移动硬盘了, 当然按第4步还是可以再次打开. 

![5.插入你的移动硬盘, 看到这个就是成功了.  楼主把他固定到了左边边栏, 以后记得每次都要安全推出了硬盘,  ...](https://gitee.com/generals-space/gitimg/raw/master/56c1b6cb3c6f282bed1fdeafed0a978d.jpg)

这个方法个人使用完全足够了, 如果有多个硬盘, 建议命名成不同的名字, 继续往fstab文件中添加描述行就可以了. 
