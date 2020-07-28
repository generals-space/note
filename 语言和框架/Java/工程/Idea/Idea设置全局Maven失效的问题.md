# Idea设置全局Maven失效的问题

参考文章

1. [解决idea设置全局maven失效的问题](https://blog.csdn.net/qq_40644583/article/details/104483891)

问题描述

MacOS

Idea: 2020.1

Maven: 3.6.3

每次设置全局 Maven 的 setting 和本地仓库后, 重启 Idea 就会恢复成默认值.

按照参考文章1中的方法可以解决, 不过在 Mac 下, 相应的配置文件路径为`~/Library/Application Support/JetBrains/IntelliJIdea2020.1/options/project.default.xml`

