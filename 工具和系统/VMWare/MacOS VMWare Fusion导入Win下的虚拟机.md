# MacOS VMWare Fusion导入Win下的虚拟机

Mac下的虚拟机格式为`.vmwarevm`(在win下会显示为一个目录, 其内容与win的虚拟机文件夹貌似没区别), 需要先导入才能使用.

找到虚拟机所在目录, 选中磁盘文件(后缀名`.vmdk`) -> 右键 -> 打开方式 -> 找到`VMVware Fusion`

![](https://gitee.com/generals-space/gitimg/raw/master/16a9b44f45cc44d124409ce1e32dd304.png)

打开之后显示如下窗口.

![](https://gitee.com/generals-space/gitimg/raw/master/262be6836962ea0d6ce54ca90ae43fb6.png)

点击继续

![](https://gitee.com/generals-space/gitimg/raw/master/9d7d23e248ff718608c4db900ecf6a74.png)

保持默认, 继续, 出现如下, 点击"选择虚拟磁盘"

![](https://gitee.com/generals-space/gitimg/raw/master/20ae8a73c22bc104a751eb1d934b0536.png)

> 注意这里, 默认是"新建虚拟磁盘", 要选择"使用现有虚拟磁盘". 前者是会创建一个空硬盘, 后者则会从目标磁盘拷贝一个副本(选前者的话, 后面启动时会出现找不到硬盘的情况, 还得重新为这个硬盘装系统才行).

从 finder 窗口中找到本次导入的虚拟机的虚拟磁盘文件, 就是最开始时右键的那个, 创建ta的副本.

![](https://gitee.com/generals-space/gitimg/raw/master/788019a1e10d02e94096f2d3ae584189.png)

选择完成后, 出现如下结果, 点击继续.

![](https://gitee.com/generals-space/gitimg/raw/master/e3d8e9a34e1f920c537612ffd7fce773.png)

点击完成, 将为此虚拟机生成Mac下的虚拟机文件(后缀名`.vmwarevm`), 记得重命名一下. 点击"存储".

![](https://gitee.com/generals-space/gitimg/raw/master/7039562e9a4dbcc22146df0ccfeeddd8.png)

出现如下界面, 表示正在拷贝原虚拟机的磁盘内容.

![](https://gitee.com/generals-space/gitimg/raw/master/55c37b9557d6d5a518bb7ba1db78bdd7.png)

完成后虚拟机将自行启动.
