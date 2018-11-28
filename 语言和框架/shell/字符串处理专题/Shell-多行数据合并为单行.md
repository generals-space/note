# Shell-多行数据合并为单行

主要是`tr`的使用, 它可以完成将一个数据流中的指定字符替换为另一个字符.

```bash
kill $(ps ux | grep crawl | awk '{print $2}' | tr '\n' ' ')
```