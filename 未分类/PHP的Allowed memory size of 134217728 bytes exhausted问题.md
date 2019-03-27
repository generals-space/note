# PHP的Allowed memory size of 134217728 bytes exhausted问题

参考文章

1. [php的Allowed memory size of 134217728 bytes exhausted问题](https://blog.csdn.net/qdujunjie/article/details/43672579)

2. [zabbix api与php的配置](http://blog.51cto.com/caiguangguang/1407422)

3. [Zabbix服务网页报错汇总](https://www.cnblogs.com/bananaaa/archive/2017/11/21/7874978.html)

情景描述:

zabbix模板页有一些组件无法显示, 查看apache的日志发现频繁出现500错误. 重启apache, 重启系统不管用.

apache的`error`日志中显示如下报错.

```
PHP Fatal error: Allowed memory size of 134217728 bytes exhausted (tried to allocate 72 bytes) in /usr/local/webdata/andy/fanli/jd/job/parse.php on line 11
```

最有效的解决办法是修改`php.ini`的`memory_limit`字段, 在php5中这个值为`128M`. 它表示`Maximum amount of memory a script may consume (128MB) 最大单线程的独立内存使用量。也就是一个web请求，给予线程最大的内存使用量的定义`.

但是修改`php.ini`后(`memory_limit`修改为1024M)重启系统不生效, `error`日志中依然报这个错, 而且还是`134217728 bytes`.

后来找到参考文章3. 它对这个问题的处理方法是, 修改`/etc/httpd/conf.d/zabbix.conf`文件, 因为在部署zabbix的时候, 可能会在zabbix.conf里写一些配置, 我这边的配置如下.

```
#
# Zabbix monitoring system php web frontend
#

Alias /zabbix /usr/share/zabbix

<Directory "/usr/share/zabbix">
    Options FollowSymLinks
    AllowOverride None
    Require all granted

    <IfModule mod_php5.c>
        php_value max_execution_time 300
        php_value memory_limit 128M
        php_value post_max_size 16M
        php_value upload_max_filesize 2M
        php_value max_input_time 300
        php_value always_populate_raw_post_data -1
        # php_value date.timezone Europe/Riga
        php_value date.timezone Asia/Shanghai
    </IfModule>
</Directory>

<Directory "/usr/share/zabbix/conf">
    Require all denied
</Directory>

<Directory "/usr/share/zabbix/app">
    Require all denied
</Directory>

<Directory "/usr/share/zabbix/include">
    Require all denied
</Directory>

<Directory "/usr/share/zabbix/local">
    Require all denied
</Directory>
```

把其中的`php_value memory_limit`的128M改成1024M再重启apache就可以了.