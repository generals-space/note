# GIT命令应用-覆盖分支

参考文章

1. [git@Osc当中怎么把一个分支的内容完全替换成另一个分支的内容呢？](http://www.oschina.net/question/1993919_224813)

我现在项目有2个分支，一个是master分支，一个是develop分支，

我们项目的成员一般推送都是往develop分支里面推送东西，master作为一个稳定版本只隔一段时间发布一个版本。

请问一下，我现在想让develop分支的所有文件覆盖master分支的文件，请问怎么做，就说简单一点，就是替换，但是不删除原有分支，完全替换。

```
git checkout master
git reset --hard develop  //先将本地的master分支重置成develop
git push origin master --force //再推送到远程仓库
```

------

使用图形界面时, 记得也是先切换回master, 然后到develop分支最新的commit处右键选择`reset maste to this commit [hard]`

这个方法适用于长时间推送到某个分支直接其稳定, 不需要与其他分支进行比对, 想直接使用此分支上的代码的情况.