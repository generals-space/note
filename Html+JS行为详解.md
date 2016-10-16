# Html+JS行为详解

1. 

```
<input type="button" onclick="clear();" value="clear"/>
```
点击时无法调用js中的clear()函数, 原因可能是clear是js的关键字(或保留字)冲突, 换一个名称即可

2. 

上传文件时, form标签的method与enctype属性都要写