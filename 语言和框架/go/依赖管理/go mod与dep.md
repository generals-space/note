# go mod与dep

`etcd-operator`工程使用`go dep`管理依赖, 尝试使用`go mod`时发现了比较有意思的事情.

```console
/home/project/etcd-operator
# ls
CHANGELOG.md  codecov.yaml        CONTRIBUTING.md  doc      Gopkg.lock  hack     MAINTAINERS  pkg        ROADMAP.md      test
cmd           code-of-conduct.md  DCO              example  Gopkg.toml  LICENSE  NOTICE       README.md  source-read.md  version
```

在`go mod init xxx`时, 貌似会自动读取`Gopkg.lock`中的内容, 不需要重新解析代码中的`import`信息.

```
/home/project/etcd-operator
# go mod init github.com/coreos/etcd-operator
go: creating new go.mod: module github.com/coreos/etcd-operator
go: copying requirements from Gopkg.lock
...省略
```
