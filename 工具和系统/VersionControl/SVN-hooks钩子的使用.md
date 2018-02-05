# SVN-hooks钩子的使用

参考文章

1. [利用SVN的post-commit钩子实现多项目自动同步](http://blog.51cto.com/qicheng0211/1563159)

2. [Windows Server配置Jenkins，实现监测SVN提交自动构建.net4.5的项目](https://www.cnblogs.com/walkerwang/p/3368986.html)

hooks, 类似于触发器, 或是事件监听回调的概念, 某个事件发生后, 触发一个动作.

下面是在码云的git仓库中添加的webhook返回的信息, 每次push操作, 它都会带着如下数据, 访问我的服务器接口.

```js
{ 
    before: '536910987d41e7d5f8e0378abfb2fd36eca528f0',
    after: 'bb345bfba94279ff19b5e153325c526db7c6ab03',
    ref: 'refs/heads/master',
    user_id: 871167,
    user_name: 'general',
    user: { 
        name: 'general',
        username: 'generals-space',
        url: 'https://git.oschina.net/generals-space' 
    },
    repository: { 
        name: 'gitimg',
        url: 'https://git.oschina.net/generals-space/gitimg',
        description: '',
        homepage: 'https://git.oschina.net/generals-space/gitimg' 
    },
    commits: [ 
        { 
            id: 'bb345bfba94279ff19b5e153325c526db7c6ab03',
            message: '移入所有gitbook所引用的图片\n',
            timestamp: '2017-09-06T14:34:45+08:00',
            url: 'https://git.oschina.net/generals-space/gitimg/commit/bb345bfba94279ff19b5e153325c526db7c6ab03',
            author: [Object] 
        } 
    ],
    total_commits_count: 1,
    commits_more_than_ten: false,
    project: { 
        name: 'gitimg',
        path: 'gitimg',
        url: 'https://git.oschina.net/generals-space/gitimg',
        git_ssh_url: 'git@git.oschina.net:generals-space/gitimg.git',
        git_http_url: 'https://git.oschina.net/generals-space/gitimg.git',
        git_svn_url: 'svn://git.oschina.net/generals-space/gitimg',
        namespace: 'generals-space',
        name_with_namespace: 'general/gitimg',
        path_with_namespace: 'generals-space/gitimg',
        default_branch: 'master' 
    },
    hook_name: 'push_hooks',
    password: 'ez1&kAm6weGw' 
}
```

git没有自带webhook这些功能, 应该是他们的开发人员写的. svn也没有自带, 但是它提供了这些事件的监听, 由我们自己来编写回调.

虽然在网页上为指定工程添加指定url的webhook是件很爽的事, 但是毕竟是用别的劳动成果. 公司自建的SVN服务器, 还是自己写代码解决.

## 如何编写

来看看如何编写svn的钩子脚本.

在svn目录的`hooks`子目录下, 默认有如下文件.

```
$ pwd
/root/svn/hooks
$ ls
post-commit.tmpl  post-lock.tmpl  post-revprop-change.tmpl  post-unlock.tmpl  pre-commit.tmpl  pre-lock.tmpl  pre-revprop-change.tmpl  pre-unlock.tmpl  start-commit.tmpl
```

在这些事件发生时, 都可以写自定义回调. 我们主要还是聚焦于`post commit`事件. 拷贝`post-commit.tmpl`模板为`post-commit`文件, 并赋予可执行权限, 就成了. 在`commit`事件发生后, svn服务会自动执行这个文件, 并传入两个参数:

- `$1`: repos-path, 就是当前提交的工程路径

- `$2`: rev, 当前提交的版本号

...只有这两个参数, git的那种当前提交的用户啊, 时间啊, 前后的commit id啊...呵呵.

不过这些信息是可以查询出来的, 参考文章1中给出了一段示例代码, 通过`svnlook`命令查询出一些信息.

```bash
#!/bin/sh
export LANG=en_US.UTF-8         # 提交日志可能为中文

REPOS="$1"                      # 仓库的路径(这个是在服务器上存储的绝对路径)
REV="$2"                        # 新提交的版本号
LOGFILE=/var/log/svnhook.log    # 钩子脚本的日志
# 脚本的标准输出和标准错误输出都打印到日志文件里
exec 1>>"$LOGFILE"
exec 2>&1
 
SVNLOOK='/usr/bin/env svnlook'
TIME=$(date "+%Y-%m-%d %H:%M:%S")
AUTHOR=$($SVNLOOK author -r $REV "$REPOS")  #本次提交作者
## CHANGEDDIRS=$($SVNLOOK dirs-changed $REPOS) #修改的目录集合
MESSAGE=$($SVNLOOK log -r $REV "$REPOS")    #提交时的备注信息

## 访问jenkins接口, 触发构建行为
## curl http://admin:a123456!@172.16.4.33:8000/job/skycmdb/build?token=python-skycmdb
```

> 其实在与jenkins时结合时也用不到这么多信息, 项目配置时只要在`构建触发器`这一节**触发远程构建**, 在身份验证时填写令牌信息就可以了.