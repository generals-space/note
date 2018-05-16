# npm与package.json创建工程

参考文章

1. [npm 与 package.json 快速入门教程](https://blog.csdn.net/u011240877/article/details/76582670)


`npm install`可以自动安装当前目录下`package.json`文件中定义的包, 类似pip的`requirements.txt`文件.

## 1. 基本信息

package.json 文件至少要有两部分内容：

1. “name” 

全部小写，没有空格，可以使用下划线或者横线

2. “version” 

`x.x.x` 的格式, 符合“语义化版本规则”

如

```json
{
    "name": "react",
    "version": "16.3.2"
}
```

其他元信息：

`description`: 描述信息，有助于搜索
`main:` 入口文件, 默认为`index.js`
`scripts`: 支持的脚本，默认是一个空的 test
`keywords`: 关键字，有助于在人们使用 npm search 搜索时发现你的项目
`author`: 作者信息
`license`: 默认是 MIT
`bugs`: 当前项目的一些错误信息，如果有的话...

## 2. 依赖信息

我们需要在`package.json`文件中指定项目依赖的包，这样别人在拿到这个项目时才可以使用`npm install`下载。

包有两种依赖方式：

- `dependencies`: 在生产环境中需要用到的依赖

- `devDependencies`: 在开发、测试环境中用到的依赖

```json
{
    "name": "react",
    "version": "16.3.2",
    "devDependencies": {
        "webpack": "^1.13.2",
    },
}
```

`dependencies`字段的内容，以 "webpack": "^1.13.2" 为例，我们知道 key 是依赖的包名称，value 是这个包的版本。那版本前面的 ^ 或者版本直接是一个 * 是什么意思呢？

**先了解 npm 包提供者应该注意的版本号规范。**

如果一个项目打算与别人分享，应该从 1.0.0 版本开始。以后要升级版本应该遵循以下标准：

- 补丁版本：解决了 Bug 或者一些较小的更改，增加最后一位数字，比如 1.0.1

- 小版本：增加了新特性，同时不会影响之前的版本，增加中间一位数字，比如 1.1.0

- 大版本：大改版，无法兼容之前的，增加第一位数字，比如 2.0.0

了解了提供者的版本规范后， npm 包使用者就可以针对自己的需要填写依赖包的版本规则。

作为使用者，我们可以在 package.json 文件中写明我们可以接受这个包的更新程度（假设当前依赖的是 1.0.4 版本）：

如果只打算接受补丁版本的更新（也就是最后一位的改变），就可以这么写： 

- `1.0`

- `1.0.x`

- `~1.0.4`

如果接受小版本的更新（第二位的改变），就可以这么写： 

- `1`

- `1.x`

- `^1.0.4`

如果可以接受大版本的更新（自然接受小版本和补丁版本的改变），就可以这么写： 

- `*`

- `x`

小结一下：**总共三种版本变化类型，接受依赖包哪种类型的更新，就把版本号准确写到前一位。**

## 安装依赖

使用npm安装package有两种方式：**本地（当前项目路径）安装**或者**全局安装**。

你选择哪种安装方式取决于你将如何使用这个包：

如果你只是想在当前项目里用 require() 加载使用，那你可以安装到本地, npm install 默认就是安装到当前工程的.

如果你想要在命令行里直接使用，比如 grunt CLI，就需要安装到全局了, `install`要加上`-g`参数.

npm install 默认会安装 package.json 中 dependencies 和 devDependencies 里的所有模块。

如果想只安装`dependencies`中的内容，可以使用`--production`字段：

```
npm install --production
npm install --dev //貌似已经被废弃了, 改用下面的
npm install --only=dev
```

## save

添加依赖时我们可以手动修改`package.json`文件，添加或者修改 dependencies devDependencies 中的内容即可。

另一种更酷的方式是用命令行，在使用`npm install`时添加`--save`或者`--save-dev`选项：

`npm install <package_name> --save`: 表示将这个包名及对应的版本添加到package.json的 `dependencies`字段

`npm install <package_name> --save-dev`: 表示将这个包名及对应的版本添加到 package.json的`devDependencies`字段.