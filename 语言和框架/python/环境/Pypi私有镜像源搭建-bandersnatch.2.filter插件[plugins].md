参考文章

1. [Mirror filtering](https://bandersnatch.readthedocs.io/en/latest/filtering_configuration.html#packages)
    - 官方文档
2. [bandersnatch报错No module Named的解决办法](https://blog.csdn.net/weixin_37813152/article/details/121634622)
    - 有效

- python: 3.6
- bandersnatch: 4.4.0

```ini
[plugins]
enabled =
    allowlist_project
    allowlist_release

[allowlist]
packages =
    bandersnatch
    ## bandersnatch==4.4.0
```

## FAQ

### No module named

见参考文章2.

```log
$ bandersnatch mirror
2022-12-16 09:56:03,563 INFO: Selected storage backend: filesystem
2022-12-16 09:56:03,731 ERROR: Unable to load entry point swift_plugin = bandersnatch_storage_plugins.swift:SwiftStorage: No module named 'keystoneauth1'
2022-12-16 09:56:03,867 INFO: Setting up mirror directory: /srv/pypi
2022-12-16 09:56:03,867 INFO: Setting up mirror directory: /srv/pypi/web/simple
2022-12-16 09:56:03,867 INFO: Setting up mirror directory: /srv/pypi/web/packages
2022-12-16 09:56:03,867 INFO: Setting up mirror directory: /srv/pypi/web/local-stats/days
```

```
pip3 install keystoneauth1
```

```log
$ bandersnatch mirror
2022-12-16 09:57:16,286 INFO: Selected storage backend: filesystem
2022-12-16 09:57:16,530 ERROR: Unable to load entry point swift_plugin = bandersnatch_storage_plugins.swift:SwiftStorage: No module named 'swiftclient'
2022-12-16 09:57:16,667 INFO: Status file /srv/pypi/status missing. Starting over.
2022-12-16 09:57:16,667 INFO: Syncing with https://pypi.org.
2022-12-16 09:57:16,667 INFO: Current mirror serial: 0
```

```
pip3 install python-swiftclient
```

### timeout

```log
$ bandersnatch mirror
2022-12-16 10:05:26,521 INFO: Selected storage backend: filesystem
2022-12-16 10:05:26,820 INFO: Initialized project plugin allowlist_project, filtering ['bandersnatch']
2022-12-16 10:05:26,837 INFO: Initialized release plugin allowlist_release, filtering [<Requirement('bandersnatch')>]
2022-12-16 10:05:26,867 INFO: Initialized project plugin allowlist_project, filtering ['bandersnatch']
...省略
2022-12-16 10:07:06,162 ERROR: Call to list_packages_with_serial @ https://pypi.org/pypi timed out: Timeout on reading data from socket
Traceback (most recent call last):
  File "/root/bandersnatch/bin/bandersnatch", line 8, in <module>
    sys.exit(main())
  File "/root/bandersnatch/lib64/python3.6/site-packages/bandersnatch/main.py", line 207, in main
    return loop.run_until_complete(async_main(args, config))
  File "/usr/lib64/python3.6/asyncio/base_events.py", line 484, in run_until_complete
    return future.result()
  File "/root/bandersnatch/lib64/python3.6/site-packages/bandersnatch/main.py", line 146, in async_main
    return await bandersnatch.mirror.mirror(config)
  File "/root/bandersnatch/lib64/python3.6/site-packages/bandersnatch/mirror.py", line 894, in mirror
    changed_packages = await mirror.synchronize(specific_packages)
  File "/root/bandersnatch/lib64/python3.6/site-packages/bandersnatch/mirror.py", line 67, in synchronize
    await self.determine_packages_to_sync()
  File "/root/bandersnatch/lib64/python3.6/site-packages/bandersnatch/mirror.py", line 282, in determine_packages_to_sync
    all_packages = await self.master.all_packages()
  File "/root/bandersnatch/lib64/python3.6/site-packages/bandersnatch/master.py", line 215, in all_packages
    raise XmlRpcError("Unable to get full list of packages")
bandersnatch.master.XmlRpcError: Unable to get full list of packages
```

直接超时, 一个包也没下载.

最初以为是因为Pypi官方把 xmlrpc 接口禁用了, pip search 无法使用, 使得 bandersnatch 直接无法使用, 那这事就大了...

后来才发现 bandersnatch 没有失败重试机制, 超时就直接退出了, 多试几次就可以了.
