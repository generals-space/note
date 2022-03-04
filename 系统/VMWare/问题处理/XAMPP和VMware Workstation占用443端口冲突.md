# XAMPP和VMware Workstation占用443端口冲突

参考文章

1. [XAMPP和VMware Workstation占用443端口冲突的解决办法](http://www.weste.net/2014/10-28/99655.html)

今天安装了一个`VMware Workstation`，发现`XAMPP`的`Apache`就启动不了. 看了一下错误日志，似乎是VMware Workstation占用了443端口导致冲突引起的. 查看了一下，原来`VMware Workstation`有个共享虚拟机的服务，占用了443端口. 

对于单机安装虚拟机来说，这个功能没有用处，禁用掉就可以了. 操作步骤如下：

1. 打开VMware Workstation，点击菜单中的"编辑->首选项";
2. 找到左侧功能列表中的"共享虚拟机"，选择后，在右侧界面中点击"更改设置";
3. 这个时候，本来是disabled的"禁用共享"按钮就被激活了，点击"禁用共享"按钮，就可以将这个功能禁用了;
4. 如果还想使用此功能，可以将443端口修改成446或者其他端口都可以. 而且不需要关闭正在运行的虚拟机;
