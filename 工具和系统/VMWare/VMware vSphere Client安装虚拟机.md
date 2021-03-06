# VMware vSphere Client安装虚拟机

参考文章

[如何通过Vmware vSphere Client安装虚拟机教程](http://jingyan.baidu.com/article/bea41d439726c1b4c51be629.html)

VMware vSphere Client客户端是用来连接与管理ESX或ESXi主机的, 在VMware vSphere Client可以方便的创建, 管理虚拟机, 并分配相应的资源. 使用vSphere Client创建虚拟机与使用VM Workstation类似, 都是需要配置虚拟机名称, CPU个数, 内存大小, 磁盘类型与容量等. 这里记录下与Workstation不同的地方.


## 1. 磁盘类型

在VM Workstation中分配硬盘大小时有"动态分配"与"固定大小"两种, 而vSphere Client创建时没这两个选项, 与之类似的选项是"thin  provision", "thick(也叫zeroedthick) "与"eagerzeroedthik "这三个.

其中默认是thick, 类似于固定大小, 创建时立即分配; thin则是动态分配, 不会立即分配设置的空间.

## 2. 系统引导

vSphere Client创建虚拟机时不会加载ISO镜像, 当创建完成后(相当于裸机), 第一次启动时加载. 可以选用光驱的方式安装, 即在虚拟机配置中将光驱的设置指向ISO镜像的地址.

在创建完成的虚拟上`右键 -> 编辑设置 -> 硬件 -> CD/DVD驱动器`.

右侧内容中, (1)客户端设备将可以使用本地的ISO镜像; (2)主机设备的作用不清楚; (3)数据存储ISO文件是需要将ISO镜像上传到ESXi服务器上才能加载的.

![](https://gitee.com/generals-space/gitimg/raw/master/518864c07beeee21d83ebac72c761664.png)

我曾尝试过使用第(3)项 - 数据存储ISO文件, 但是镜像文件太大, 使用filezilla上传的文件大小有限制, 上传总是失败. 只能使用第(1)项. 过程如下:

### 2.1 

首先启动虚拟机(这时无论是虚拟的硬盘还是光驱中都没有镜像存在, 它会自动寻找一段时间). 然后在工具栏上点击"连接/断开虚拟机的CD/DVD设备" -> 连接到本地磁盘上的ISO镜像, 选择本地磁盘上的ISO镜像. 转到虚拟机的控制台标签(或右键目标虚拟机->打开控制台).

如果显示没有找到操作系统则看下面的2.2.

![](https://gitee.com/generals-space/gitimg/raw/master/6187185fa302529fa276ed04f3062f4f.jpg)

### 2.2

启动虚拟机电源之后可能来不及找到本地的ISO镜像并让系统连接到设备(每次重启设备的设置都会丢失, 每次都要打开光驱选择ISO), 所以需要推迟引导(即打开电源等待一段时间后再去查询引导设备, 这样就足够我们设置好光驱了).

首先, 右键目标虚拟机 -> 编辑设置 -> 顶部的"选项"标签 -> 引导选项 -> 右侧内容中选中强制执行BISO设置. 然后开机/重启, 这样下次启动时可以设置引导顺序, 将光驱的顺序放在第一位.

![](https://gitee.com/generals-space/gitimg/raw/master/3dcfb8deb75ca4a0241eba3caa7a1a6a.jpg)

BIOS中的启动顺序就不用说了吧?

** 设置完成后不要急着再次重启, 不然下次启动还是进入BIOS. **

大致还是上次的位置, 右键目标虚拟机 -> 编辑设置 -> 顶部的"选项"标签 -> 引导选项 -> 右侧内容中取消强制执行BISO设置, 然后在它上面, "打开电源引导延迟", 设置它的值为10000ms(注意单位), 单击确定, 保存.

![](https://gitee.com/generals-space/gitimg/raw/master/d2f1d8577f3efa3e6dd2d0043050a048.jpg)

然后保存BIOS的引导设置, 重启, 你将看到BIOS在倒计时, 并没有查找引导设备, 这时候去设置光驱要加载的ISO. 10s过后将进入正确的系统安装界面.