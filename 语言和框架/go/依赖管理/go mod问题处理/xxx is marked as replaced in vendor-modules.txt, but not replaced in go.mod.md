# xxx is marked as replaced in vendor-modules.txt, but not replaced in go.mod

参考文章

1. [cmd/go: 1.18 fails when vendoring and main module replaces itself](https://github.com/golang/go/issues/51285)

在编译 containerd v1.5.8 版本时, 将 go.mod 文件中的 kubernetes 相关包 replace 到了本地包, 然后 make 就报错了.

```console
$ make
+ bin/ctr
go: inconsistent vendoring in /home/github.com/containerd/containerd:
	k8s.io/api: is replaced in go.mod, but not marked as replaced in vendor/modules.txt
	k8s.io/apimachinery: is replaced in go.mod, but not marked as replaced in vendor/modules.txt
	k8s.io/apiserver: is replaced in go.mod, but not marked as replaced in vendor/modules.txt
	k8s.io/client-go: is replaced in go.mod, but not marked as replaced in vendor/modules.txt
	k8s.io/cri-api: is replaced in go.mod, but not marked as replaced in vendor/modules.txt

	To ignore the vendor directory, use -mod=readonly or -mod=mod.
	To sync the vendor directory, run:
		go mod vendor
make: *** [bin/ctr] Error 1
```

vendor 目录下存在 gopath 模式下依赖的包, 同时包含一个 modules.txt 文件.

可以再次执行 go mod vendor, 将 replace 后的本地包, 放到 vendor 目录, 再次 make 就可以了.
