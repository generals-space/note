# Process

参考文章

1. [Java魔法堂：找外援的利器——Runtime.exec详解](http://www.cnblogs.com/fsjohnhuang/p/4081445.html)
2. [Java Runtime.exec()的使用](http://www.cnblogs.com/mingforyou/p/3551199.html)
3. [execute file from defined directory with Runtime.getRuntime().exec](http://stackoverflow.com/questions/10689193/execute-file-from-defined-directory-with-runtime-getruntime-exec)

`Process`是java内置的用于执行 shell 命令的工具库.

`Runtime.getRuntime()`提供了以下几种`exec()`方法: 

| 函数原型                                                   | 使用说明                                                |
| :--------------------------------------------------------- | :------------------------------------------------------ |
| `Process exec(String command)`                             | 在单独的进程中执行指定的字符串命令.                     |
| `Process exec(String[] cmdarray)`                          | 在单独的进程中执行指定命令和变量.                       |
| `Process exec(String[] cmdarray, String[] envp)`           | 在指定环境的独立进程中执行指定命令和变量.               |
| `Process exec(String[] cmdarray, String[] envp, File dir)` | 在指定环境和工作目录的独立进程中执行指定的命令和变量.   |
| `Process exec(String command, String[] envp)`              | 在指定环境的单独进程中执行指定的字符串命令.             |
| `Process exec(String command, String[] envp, File dir)`    | 在有指定环境和工作目录的独立进程中执行指定的字符串命令. |

各个参数的含义: 

- `command`:    一条指定的系统命令. 
- `cmdarray`:   包含所调用命令及其参数的数组. 
- `envp`:       字符串数组, 其中每个元素的环境变量的设置格式为name=value; 如果子进程应该继承当前进程的环境, 则该参数为 null. 
- `dir`:        子进程的工作目录; 如果子进程应该继承当前进程的工作目录, 则该参数为 null. 

其中, 其实`cmdarray`和`command`差不多, 同时如果参数中没有`envp`参数或设为`null`, 表示调用命令将在当前程序执行的环境中执行; 

如果没有`dir`参数或设为`null`, 表示调用命令将在当前程序执行的目录中执行, 因此调用到其他目录中的文件和脚本最好使用绝对路径. 

为了执行调用操作, JVM会启一个`Process`, 可以通过调用`Process`类的以下方法, 得知调用操作是否正确执行: 

```java
abstract int waitFor()
```

导致当前线程等待, 如有必要, 一直要等到由该 Process 对象表示的进程已经终止. 
返回进程的出口值. 根据惯例, 0 表示正常终止; 否则, 就表示异常失败. 
另外, 调用某些Shell命令或脚本时, 会有返回值, 那么我们如果捕获这些返回值或输出呢？为了解决这个问题, Process类提供了: 

```java
abstract InputStream getInputStream() // 获取子进程的输入流. 最好对输入流进行缓冲. 
```

```java
package com.example.demo;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.lang.Process;
import java.lang.Runtime;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
public class DemoApplication {
	private static final Logger logger = LoggerFactory.getLogger(DemoApplication.class);

    @GetMapping("/")
	public String home(){
		String cmd = "ls /Users/general";
		int exitCode = 0;
		StringBuilder sb = new StringBuilder();
		// List<String> messages = new ArrayList<>();

		try {
			String line;
			Process subproc = Runtime.getRuntime().exec(cmd);
			exitCode = subproc.waitFor();
			if(exitCode == 0){
				InputStream inputStream = subproc.getInputStream();
				BufferedReader stdoutReader = new BufferedReader(new InputStreamReader(inputStream));

				while ((line = stdoutReader.readLine()) != null) {
					sb.append(line).append("\n");
					// messages.add(line);
				}
			} else {
				logger.error("执行出错");
				InputStream errorStream = subproc.getErrorStream();
				BufferedReader stderrReader = new BufferedReader(new InputStreamReader(errorStream));
				while ((line = stderrReader.readLine()) != null) {
					sb.append(line).append("\n");
					// messages.add(line);
				}
			}
		} catch(IOException | InterruptedException e) {
			e.printStackTrace();
		}
		logger.info("执行结果: " + sb.toString());

		return "hello world";
	}

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}

}

```
