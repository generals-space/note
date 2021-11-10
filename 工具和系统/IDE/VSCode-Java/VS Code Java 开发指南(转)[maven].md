# VS Code Java 开发指南(转)

原文链接

[终极指南：如何使用 Visual Studio Code 进行 Java 开发？](https://www.infoq.cn/article/wsak-nm2zhk65ydrudgj)

[英文原文](https://dzone.com/articles/visual-studio-code-for-java-the-ultimate-guide-201)

VS Code 支持 了 Java 开发，许多 Java 拥护者和演讲者都使用它上台演讲做现场演示。

以下是 2019 年 Java 开发人员使用 VS Code 开发、运行、调试和部署其应用程序的终极指南。

如果您尚未下载 VS Code，那现在就下载并安装它吧。接下来本指南将假定您至少安装了 Java 8，尽管 VS Code 也支持 Java 11。您也可以完全跳过本指南，直接参考文档《 用 VS Code 开发 Java 应用》。但是，建议对照下面的指南走查一遍重点部分，这将有助于您更好地利用 VS Code 进行 Java 开发。

## 初始设置

为了使 VS Code 支持核心 Java，作为最低要求，您必须安装一组扩展。 让我们逐一深入研究这些扩展。

### 微软的 Java 扩展包（Java Extension Pack）

这个扩展包包含了下文将要详细介绍的五个扩展。借助它，你无需过多了解就可以开始使用 VS Code。

### 红帽的 Java 语言支持扩展（Language Support for Java）
要使 VS Code 支持 Java，这是唯一一个必须安装的扩展。其它扩展都是补充性的，您需要根据要开发的项目决定是否要安装它们。

安装完这个扩展后，您就可以阅读和编辑 Java 源代码了。首先创建一个 Helloworld.java 文件，然后用 VS Code 打开 (或者在 VS Code 中创建该文件，然后把它保存在某个文件夹中)。

![](https://gitee.com/generals-space/gitimg/raw/master/08db32e1dd6cf650b325a03e5a87bc57.png)

![](https://gitee.com/generals-space/gitimg/raw/master/9ce5e08c3cae57a89336c80faa729707.png)

![](https://gitee.com/generals-space/gitimg/raw/master/41a3cf0b60b34bbc67b7db9b161cd156.png)

![](https://gitee.com/generals-space/gitimg/raw/master/85925d8354836779388f7f0e38de4742.png)

完成后，您可以打开终端（在 Mac OS 中，使用组合键 ⌘+`），然后输入命令 javac HelloWorld.java 进行编译。

![VS Code 里内嵌的终端](https://gitee.com/generals-space/gitimg/raw/master/673f619fdec47c39b1f03e2552c5fb9f.png)

编译后会产生一个 Helloworld.class 文件。最后，用 Java 命令运行这个类：java helloworld。

Java 语言支持扩展通过使用 Eclipse 语言服务器协议（Eclipse Language Server Protocol）支持 Java。了解有关 Eclipse LSP 的更多信息。

### 其它特性

Java 语言支持扩展添加了许多其它功能，可以帮助您快速浏览、编写、重构和阅读 Java 源代码，您不妨使用 VS Code 这个轻量级文本编辑器来替代其它 IDE。

![面包屑导航 (顶部) — 大纲视图 (左下角)](https://gitee.com/generals-space/gitimg/raw/master/bbbbd54c4b883d9dedca7fabbf30466b.png)

查看概述页面可以获取这个扩展的特性以及重构快捷键的完整列表。

### 微软的 Java 调试器（Debugger for Java）

掌握了在 VS Code 中编写和阅读 Java 代码的基础知识后，下一步自然就是运行和调试代码了。 这正是该扩展提供的功能。 这个扩展使用您计算机上的默认 JAVA_HOME，当然您也可以自定义它。

![运行 | 调试 Java 应用](https://gitee.com/generals-space/gitimg/raw/master/42332c55fa900f9b8f59faf19b1cf127.png)

它具备常见的 Java IDE 调试特性的所有能力，并且支持更多的自定义，可以让您控制代码如何被执行以及调试器如何连接到 JVM。 它还支持远程 JVM。

安装这个扩展后，您会在主方法上方看到两个超链接，如上图所示。若您单击运行，代码将被编译并执行。您还可以设置断点并点击调试。

![](https://gitee.com/generals-space/gitimg/raw/master/9d0dc49115095ecf46e8a1d51734db13.png)

对于远程调试，您需要添加新的配置。 切换到调试视图（在 Mac 上按 Shift +⌘+ D）并点击配置按钮⚙。 这样会打开 launch.json 文件。 单击屏幕上的添加配置蓝色按钮。 这会打开一个如上图所示的弹出菜单。

现在，您可以自定义一个能插入远程 JVM 的 启动项了。您只需提供主机名、端口号等详细信息。

和其它调试 IDE 一样，您可以在运行期间查看变量、堆栈追踪，甚至对变量内容进行更改。

![VS Code 里在调试过程中更改 Java 变量](https://gitee.com/generals-space/gitimg/raw/master/a1e5896ab4726a25e9015468fa0d0088.png)

至此，我们完成了用于阅读、编写、运行和调试 Java 代码的 VS Code 基本设置。

## 中级设置

掌握了 Java 编码的基础知识后，您很快就需要使用库、依赖项、类路径等。在 VS Code 上进一步改进 Java 支持的最佳方法是添加以下三个扩展：

1. Java 依赖查看器（Java Dependency Viewer）
2. 针对 Java 的 Maven 扩展 （Maven for Java）
3. 微软的 Java 测试运行器（Java Test Runner）

让我们分别看一下这几个扩展。

### 微软的 Java 依赖查看器

这个扩展为您提供两个核心功能。 其中主要的一个功能是提供了“项目”的概念，您可以手动向项目中添加库（JAR）。 第二个功能使项目当前设置的 classpath 可视化，即使是 Maven 项目（参见下文针对 Java 的 Maven 扩展）。

打开命令托盘（Shift +⌘+ P）并输入 create java：

![创建一个 Java 工程](https://gitee.com/generals-space/gitimg/raw/master/9d00c584d65474f0b544269a3f9357ee.png)

您需要选择创建项目的位置。项目由与项目名称同名的文件夹组成（您选择了创建项目的位置后，接下来命令托盘会继续询问项目名称，例如你可以输入 myworkspace）。

创建项目后，VS Code 将在新窗口中打开这个新文件夹。

![新的 Java 工程](https://gitee.com/generals-space/gitimg/raw/master/b3feb5c5f587199992361db19caa0e7f.png)

就像您看到的那样，这个项目具有一个基本结构，其中包含了 bin 和 src 文件夹。 在 src 中，开始会有一个基本的 Java 类。 如果您是一位经验丰富的 Java 开发人员，一眼就会发现这个扩展使用的是 Eclipse 项目的格式，这是因为它与 Eclipse 语言服务器协议和其它扩展能很好的协同工作。

### 添加类库和 JAR 包

您可以编辑 .classpath 文件，指定全部自定义 JAR 包所在的目录，这个目录可以放在任何位置，例如某个 lib 文件夹。这些扩展会自动加载 classpath 中包含的类库，使你能够运行自己的代码。

![编辑.classpath 文件增加类库](https://gitee.com/generals-space/gitimg/raw/master/0f7655529df8cd9c4aff6d4b0dca73f8.png)

### 微软的针对 Java 的 Maven 扩展

Maven 是 Java 生态系统中使用最广泛的项目构建和依赖关系管理工具。因此，通过该扩展，您几乎可以用 VS Code 处理任何类型的 Java 项目。

您将能够通过 Maven 原型（archetype）生成和引导 Maven 项目、管理依赖关系并触发 Maven 目标（goal），并借助一些智能代码补全功能编辑 pom. xml 文件。

![Maven 命令](https://gitee.com/generals-space/gitimg/raw/master/a16a73e0dd0061e9100e1ea81d537e2f.png)

让我们来看一下：

- 再次打开命令托盘，然后输入 Maven。
- 选择 Generate from Maven Archetype。
- 选择 maven-archetype-quickstart。

这个扩展会要求您选择目标文件夹，以便在其下面生成项目文件夹。 输入焦点会跳转到终端，您必须在那里输入 Maven 命令行的参数，不过不用担心，它会一步一步地引导您。

创建项目后，直接从终端调用 code 即可打开它。

![在 VS Code 中打开新创建的目录](https://gitee.com/generals-space/gitimg/raw/master/795a04eace571bddc439c7c64d86753a.png)

好了，您现在应该已经在 VS Code 中打开您的 Maven 项目了。您可以做的最基本的事情就是运行您的代码。 您有两个选择：

1. 如前所述，使用 App 类中 main 方法旁边的 Run 超链接运行您的代码。
2. 使用 Maven。

如果您使用 Java 调试扩展（运行 | 调试）触发器，扩展将使用 Maven 生成的 classpath，以确保所有依赖项都正确地添加到类路径中。

如果使用 Maven 运行 Java 代码，您可以像往常一样使用终端，或者打开命令托盘并输入 Maven Execute Commands。

![显示 Maven 动作的命令托盘](https://gitee.com/generals-space/gitimg/raw/master/aaf5e0118c0319a29befe8d534a6af51.png)

它会要求您选择一个项目。 由于您只有一个项目，直接在其上按回车即可。 接下来，您将看到一个列表，它包含了所有默认的核心 Maven 目标。 选择 package 生成 JAR 文件。

![显示 Maven 动作的命令托盘](https://gitee.com/generals-space/gitimg/raw/master/20a1d16e8ffc07f425f0198b7be6dbe2.png)

如果要运行自定义目标，例如从 Maven 插件继承的目标，您可以使用 Maven 视图：

![执行来自 Maven 插件的目标](https://gitee.com/generals-space/gitimg/raw/master/30c895774ec01bf760c919451303517f.png)

编辑 pom.xml 文件并添加依赖项后，VS Code 将自动重新加载 classpath，然后您就可以从新的依赖项中导入类和包。 这个过程非常干净、顺畅。

### Microsoft 的 Java 测试执行器

最后一步是增强单元测试的运行、调试和测试结果的可视化。此扩展程序将超链接添加到可以单独执行的单元测试 (支持 JUnit 和 TestNG)，您可以立即在 VS Code 中看到报告，如下面的示例所示。

![在 VS Code 中运行单元测试](https://gitee.com/generals-space/gitimg/raw/master/5060aeb43edb0f7fb78c9cea3556a4ae.png)

此扩展还将启用测试资源管理器视图，因此您可以专注于代码的单元测试，并以更加符合测试驱动开发（TDD）的方式编写软件。

![测试浏览器](https://gitee.com/generals-space/gitimg/raw/master/2bdabbf39eafe1c1844a3b2d618eddcc.png)

此扩展目前仅适用于 Maven 项目，因此请确保您安装了针对 Java 的 Maven 扩展。

## 高级设置

如果您现在对使用 VS Code 进行 Java 开发感到满意，那么是时候进一步改善您的使用体验了。 以下是一些扩展列表，可以改善您的日常工作体验。

这只是接下来要添加哪些扩展的建议，而不是事实标准的列表，但它可以让您先行一步。

### GitLens

![](https://gitee.com/generals-space/gitimg/raw/master/5ae502971579b12c4c0e26b083fc85fc.png)

希望您已经在使用 Git 了，无论是通过 GitHub 还是其它任何服务或环境。此扩展为您提供有关源代码著作信息的洞察，例如 “谁添加了此方法以及何时添加”。

上图就是安装了 GitLens 后 Java 类的样子。看看那些没有数字的行，它们是对提交历史的注释。 您也可以简单地将鼠标悬停在特定的行上，它将显示这一行是何时、由谁、在哪个提交哈希上添加的。

### Rest 客户端

如果您是构建 REST API 的开发人员，那么这是您必须安装的 VS Code 的扩展。 有了它，您将能够编辑包含 HTTP 调用的 .http 文件。 编辑器将快速提供代码片段和模板，它会为您提供一个即点生效的神奇的超链接，它会触发 HTTP 调用并在旁边打开结果。下图是一个快速浏览。

![](https://gitee.com/generals-space/gitimg/raw/master/5ae502971579b12c4c0e26b083fc85fc.png)

就这样了！ 您现在拥有一份完整的 VS Code 设置，可以实际进行任何类型的 Java 开发了。

## 福利：Pivotal 的 Spring 设置和 Gradle

如果你是一个狂热的 Spring 开发人员，一定想知道 Pivotal 和微软提供的那些能增强 Spring Boot 应用开发体验的重要扩展。

最后，有一个可以帮助您编写 build.gradle 文件的 Gradle 扩展。

以下是一些额外的 Spring 工具供进一步学习：

Spring Boot Tools

Spring Initializr Java Support

Spring Boot Dashboard

Gradle Language Support
