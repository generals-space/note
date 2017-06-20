## 列出目标目录下的文件内容

```java
import java.io.File;

public class App
{
    public static void main( String[] args )
    {
        String targetPath = "/home/pay/Downloads";
        File file = new File(targetPath);
        File[] tempList = file.listFiles();
        System.out.println("该目录下对象个数："+tempList.length);
        for(int i = 0; i < tempList.length; i ++){
                if(tempList[i].isFile()){
                        System.out.println("文件: " + tempList[i]);
                }
                if(tempList[i].isDirectory()){
                        System.out.println("目录: " + tempList[i]);
                }
        }
    }
}
```

## 调用命令行

参考文章

[Java魔法堂：找外援的利器——Runtime.exec详解](http://www.cnblogs.com/fsjohnhuang/p/4081445.html)

[Java Runtime.exec()的使用](http://www.cnblogs.com/mingforyou/p/3551199.html)

[execute file from defined directory with Runtime.getRuntime().exec](http://stackoverflow.com/questions/10689193/execute-file-from-defined-directory-with-runtime-getruntime-exec)

简单调用

`exec(String command)`: 调用外部程序，入参command为外部可执行程序的启动路径或命令。

```java
Process proc = Runtime.getRuntime().exec("/usr/bin/cp /etc/yum.conf /tmp/yum.conf");
```

图形界面下调用可以唤出`gedis`工具.

```java
Process proc = Runtime.getRuntime().exec("gedit");
```



对于某些需要chdir改变工作目录的命令, 需要在`exec`函数中指定运行目录.

`exec(String command, String[] envp, File dir)`: 除了设置系统环境变量外，还通过参数dir设置当前工作目录。

```java
Process proc = Runtime.getRuntime().exec("npm start", null, new File("/home/pay/Downloads/electron-api-demos"));
```

```
Process proc = r.exec("cmd /c dir > %dest%", new String[]{"dest=c:\\dir.txt", new File("d:\\test")});
```