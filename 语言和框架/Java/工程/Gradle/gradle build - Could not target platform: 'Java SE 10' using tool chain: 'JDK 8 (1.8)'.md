# gradle build - Could not target platform: 'Java SE 10' using tool chain: 'JDK 8 (1.8)'

参考文章

1. [Gradle 错误:Eclipse环境下gradle报错Could not target platform: 'Java SE 10' using tool chain: 'JDK 8 (1.8)'.](https://blog.csdn.net/blueboz/article/details/82822113)

```console
$ gradle build
Starting a Gradle Daemon, 1 busy Daemon could not be reused, use --status for details
> Task :compileJava FAILED

FAILURE: Build failed with an exception.

* What went wrong:
Execution failed for task ':compileJava'.
> Could not target platform: 'Java SE 11' using tool chain: 'JDK 8 (1.8)'.

* Try:
Run with --stacktrace option to get the stack trace. Run with --info or --debug option to get more log output. Run with --scan to get full insights.

* Get more help at https://help.gradle.org

Deprecated Gradle features were used in this build, making it incompatible with Gradle 7.0.
Use '--warning-mode all' to show the individual deprecation warnings.
See https://docs.gradle.org/6.6/userguide/command_line_interface.html#sec:command_line_warnings

BUILD FAILED in 3m 50s
1 actionable task: 1 executed
You have new mail in /var/mail/general
```

这是因为我在创建`gradle`工程时, 选择了使用 Java 11 来编译(默认), `gradle.build`文件如下.

```
group = 'space.generals.java'
version = '0.0.1-SNAPSHOT'
sourceCompatibility = '11'
```

把`11`改成`1.8`重新构建即可.
