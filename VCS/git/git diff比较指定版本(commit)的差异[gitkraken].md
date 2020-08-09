# git diff比较指定版本(commit)的差异[gitkraken]

参考文章

1. [git比较两个版本之间的区别](https://blog.csdn.net/zhezhebie/article/details/78496693)
    - `git diff 版本号码1 版本号码2` 查看任意两个版本之间的改动
    - `git diff 版本号码1 版本号码2 src` 比较两个版本号码的src 文件夹的差异
2. [gitkraken - How to compare 2 branches](https://stackoverflow.com/questions/53234113/gitkraken-how-to-compare-2-branches)
    - 使用`gitkraken`对比两个(可能是不同branch的)commit的差异.
3. [Diff, Blame and History](https://support.gitkraken.com/working-with-commits/diff/)
    - 官方文档, 比较两个 commit 的差异.

场景描述

我自己在写代码的时候, 经常会基于生产的某个基础分支(如公共开发分支)创建自己的开发分支进行编写, 但是在这个分支上很容易放飞自我, 如果在提测前将分支合并到公共开发分支的话, 感觉有点不太好, 所以希望能比较我自己的分支与公共开发分支的差异, 重新基于此时的公共开发分支(在我编写完成自己的模块后, 公共开发分支可能也有新的提交了), 把我做的修改搬过去.

这样就需要比较我目前最新的提交与公共开发分支最新的提交之间的变动, 使用`git diff`可以完成, 但是命令行看起来太不方便, 所以想找`gitkraken`有没有办法完成这样的操作.

答案当然是"有".

![](https://gitee.com/generals-space/gitimg/raw/master/7f76835193f22b595beb87ab94912aa9.png)

如果分支过多, 比较的两个 commit 之间夹杂了许多其他分支的提交的话. 可以选择将第一个分支进行 solo 展示(隐藏其他所有分支), 然后再按shift选取目标 commit, 这样干扰会少很多(没试过).
