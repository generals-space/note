# io.File列出目标目录下的文件内容

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
