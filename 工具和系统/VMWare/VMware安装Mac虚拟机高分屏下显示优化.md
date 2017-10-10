# VMware安装Mac虚拟机高分屏下显示优化

参考文章

[轻松开启Hi-DPI, 尽享OS X清晰世界](http://bbs.pcbeta.com/viewthread-1679769-1-1.html)

VMware安装的Mac OS X在高分屏下显示的字体太小, 而Mac本身没有调节字体大小的功能, 只能通过分辨率调节. 在4k屏幕下同时有3840x2160与1920x1080两种分辨率, 3840x2160还是字体还是太小, 使用1920x1080倒是正合适而且字体也不失真模糊. 但是由于不是自带的Retina屏, 在笔记本本身的3k(3200x1800)显示屏下只有一个3840x2160的分辨率, 字体太小了, 而且没有其他可调的分辨率. 

根据参考文章中的做法, 执行如下命令并赋予`~/enable-HiDPI.sh`这个脚本的执行权限.

```shell
curl -o ~/enable-HiDPI.sh https://raw.githubusercontent.com/syscl/Enable-HiDPI-OSX/master/enable-HiDPI.sh
chmod +x ~/enable-HiDPI.sh
```

执行该脚本, 它会提示输入你想适配的屏幕分辨率, 之后它会计算得到一个比较低的分辨率让你选择, 这样字体会整体放大. (虽然分辨率低了, 但是字体依然很清晰)

```
$ ./enable-HiDPI.sh 
## 需要root权限
Password:
[  OK  ]  Remove /Users/general/Downloads/DisplayVendorID-.
## 这里输入了3200x1800
Enter the Resolution you want to enable HiDPI(e.g. 1600x900, 1440x910, ...), enter 0 to quit: 3200x1800
## 然后输入0结束
Enter the Resolution you want to enable HiDPI(e.g. 1600x900, 1440x910, ...), enter 0 to quit: 0
[ ---> ]  Backuping origin Display Information...
cp: /System/Library/Displays/Contents/Resources/Overrides/DisplayVendorID-: Operation not permitted
cp: /Users/general/Downloads/DisplayVendorID-: unable to copy extended attributes to /System/Library/Displays/Contents/Resources/Overrides/DisplayVendorID-: Operation not permitted
cp: /System/Library/Displays/Contents/Resources/Overrides/DisplayVendorID-/DisplayProductID-: No such file or directory
[  OK  ]  Done, Please Reboot to see the change! Pay attention to use Retina Display Menu(RDM) to select the HiDPI resolution!.
```

重启之后, 打开`系统偏好设置`->`显示器`->`缩放`, 在3k(3200x1800)屏幕上会看到两个分辨率3200x1800和1600x900, 选择3200x1800时字体一如既往的小, 而选择1600x900时就比较合适了, 而且字体照样清晰.

这里记录下该脚本的github地址: [Enable HiDPI](https://github.com/syscl/Enable-HiDPI-OSX)