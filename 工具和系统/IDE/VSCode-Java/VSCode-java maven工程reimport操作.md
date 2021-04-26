# VSCode-java maven工程reimport操作

参考文章

1. [vscode官方文档 Java project management in VS Code](https://code.visualstudio.com/docs/java/java-project)

有时候从 git 上 clone 一个 java 工程, 总是会先报很多包找不到的错误. 但是在命令行构建执行又是可以运行起来了, 这样在开发时就非常不爽了.

尤其是, 跟踪那些找不到的包时, 找到目标包后, 下面的报错就消失了, 就跟缓存一样.

但是这种问题一报就是几十上百个, 总不能一个一个去找吧.

在idea中, maven 工程有一个reimport 按钮, 我找了找, 在vscode左侧`Maven Projects`, 还有什么`Java Dependencies`这些栏中都没有找到`reimport`选项.

后来找到了参考文章1, 使用`Ctrl+Shift+P`, 输入`Java: Import Java projects in workspace`即可实现`reimport`的效果, 实践有效.

