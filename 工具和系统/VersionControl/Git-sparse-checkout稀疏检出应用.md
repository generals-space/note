# Git-sparse-checkout稀疏检出应用

参考文章

1. [Git拉出无效的Windows文件名](https://xbuba.com/questions/33702140)

2. [使用Sparse Checkout，排除跟踪Git仓库中指定的目录](https://www.jianshu.com/p/e82c89e187c5)

3. [Git只获取部分目录的内容（稀疏检出）](https://www.jianshu.com/p/b6c61907049f)

稀疏检出的目的是只checkout一个仓库中的指定目录(可以同时指定白名单和黑名单). 一般在如下场景中非常有用.

1. 仓库非常大, 但是自己只需要更新其中的某个子目录;

2. Mac/Linux下提交的文件名包含windows中不支持的字符(或是路径过长什么的), 导致在windows下进行`clone`或`pull`失败, 可以屏蔽该非法路径;

开启`sparse checkout`需要修改`.git`下的文件(你一定不想改全局配置吧), 所以本地要先init一个空仓库, 把remote指向目标仓库, 再pull才行.

另外, 参考文章3有提到关闭`sparse checkout`的操作貌似不简单, 不过目前我也不关心.
