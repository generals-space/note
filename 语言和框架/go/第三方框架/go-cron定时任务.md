# go-cron定时任务

参考文章

1. [Go cron定时任务的用法](https://www.cnblogs.com/zuxingyu/p/6023919.html)

...全网就只有一个流行的go-cron库`robfig/cron`, 不用费心选择了.

与linux的crontab比起来, 粒度精确到了秒级(共用6个字段: 秒, 分, 时, 日, 月, 周).