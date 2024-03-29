# git 大文件清理

参考文章

1. [Git从库中移除已删除大文件](https://www.cnblogs.com/liangqihui/p/9880916.html)
2. [删除Git仓库中的大文件](https://cloud.tencent.com/developer/article/1559335)
3. [git项目大小优化笔记,删除历史提交中的大文件](https://www.cnblogs.com/fuhua/p/15527023.html)
    - 前两个参考文章中提到的方法只有在`push --force`然后再下载后才能看到效果, 只有这篇文章是可以直接把`.git`目录缩小的.

注意执行如下步骤时最好已经把大文件从仓库中删除了, 只要再从历史记录中移除就可以了.

## 1. 显示最大的文件的hash id

```console
$ git verify-pack -v .git/objects/pack/pack-*.idx | sort -k 3 -g | tail -n 1
a6616e84781c884c4524006be3079b5ff983ba38 blob   64047312 63520027 49277651
```

> `.git/objects/pack/`目录下只有`.idx`和`.pack`两种文件

`verify-pack`的结果中会有`commit`, `tree`, `blob`3种类型, 我们要找的大文件一般都是`blob`类型.

`sort -k 3 -n`: 以第3列信息进行排序. 由于`sort`默认的排序方法是按字母排序(如[0, 10, 101, 1015, 106, 11]), `-n`即是让`sort`按数值排序.

## 2. rev-list 查看文件路径

使用`rev-list`命令, 传入`--objects`选项, 它会列出所有 commit SHA 值, blob SHA 值及相应的文件路径. 

我们使用这个命令查看`blob`的文件名和路径(一般心中有数).

```console
$ git rev-list --objects --all | grep a6616e84781c884c4524006be3079b5ff983ba38
a6616e84781c884c4524006be3079b5ff983ba38 files/VSCodeUserSetup-x64-1.49.0.exe
```

这里`grep`的目标为上一步找到的大文件的 hash id.

## 3.

确认路径没错, 从commit历史中找到所有修改该文件的commit, 然后修改这些commit

```console
$ git log --pretty=oneline --branches -- files/VSCodeUserSetup-x64-1.49.0.exe
c31a96fbb31ae8617e21384e0fb5bcbe86afd73b (HEAD -> bigfile, origin/bigfile) bigfile
```

注意: 

1. 后面的文件路径要写全, 只写文件名是不行的.
2. 这个文件是从哪个分支提交的, 就要`checkout`到哪个分支, 否则是找不到的.

## 4. 重写所有修改这个文件的提交

找到所有修改这个对象的commit后, 我们找到最早的修改, 然后使用`filter-branch`命令来操作, 具体如下: 

```
$ git filter-branch --index-filter 'git rm --cached --ignore-unmatch files/VSCodeUserSetup-x64-1.49.0.exe' -- c31a96fbb31ae8617e21384e0fb5bcbe86afd73b
```

或者不用执行`git log`那一步, 直接从所有提交中删除这个对象: 

```console
$ git filter-branch --index-filter 'git rm --cached --ignore-unmatch files/VSCodeUserSetup-x64-1.49.0.exe' -- --all

Rewrite c31a96fbb31ae8617e21384e0fb5bcbe86afd73b (41/46) (2 seconds passed, remaining 0 predicted)    rm 'files/VSCodeUserSetup-x64-1.49.0.exe'

Ref 'refs/heads/bigfile' was rewritten
WARNING: Ref 'refs/heads/master' is unchanged
WARNING: Ref 'refs/remotes/origin/master' is unchanged
Ref 'refs/remotes/origin/bigfile' was rewritten
WARNING: Ref 'refs/remotes/origin/js' is unchanged
WARNING: Ref 'refs/remotes/origin/master' is unchanged
```

可以看到, 只有`bigfile`这个分支发生了变动, 其他分支都是`unchanged`, 这是因为目前只有在`bigfile`中有这个文件.

必要的时候, 需要用`-f`选项来强制地进行删除: 

```
git filter-branch -f --index-filter 'git rm --cached --ignore-unmatch files/VSCodeUserSetup-x64-1.49.0.exe' -- --all
```

## 5. 

这里需要删除`.git/refs`目录下的一些引用文件并重新打包, 具体命令如下, 比较固定:

```
rm -Rf .git/refs/original
rm -Rf .git/logs
git gc
```

## 6. 

删除多个大文件时, 每次执行到第5步, 再回去删除另一个. 全部删除完成后可以`git push`进行推送, 但是可能会报错.

```console
$ git push
To https://gitee.com/generals-space/snippet.git
 ! [rejected]        bigfile -> bigfile (non-fast-forward)
error: failed to push some refs to 'https://gitee.com/generals-space/snippet.git'
```

可以使用`-f`强推.

------

注意: 可能经过上面的步骤后, 本地的仓库依旧很大, 这是因为`.git/object`目录下还存在大文件, 但是推送到远程后重新clone, 就可以得到较小的工程了.
