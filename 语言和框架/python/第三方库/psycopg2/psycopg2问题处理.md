# psycopg2问题处理

## 1. FATAL: remaining connection slots are reserved for non-replication superuser connections

参考文章

1. [Heroku “psql: FATAL: remaining connection slots are reserved for non-replication superuser connections”](https://stackoverflow.com/questions/11847144/heroku-psql-fatal-remaining-connection-slots-are-reserved-for-non-replication)

**场景描述**

celery worker启动了100个进程, 在启动过程中, 报了上述错误. 

按照参考文章1中所说, 将数据库的最大连接数调高就可以了.
