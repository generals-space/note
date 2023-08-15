参考文章

1. [Go全局唯一ID选型集合](https://blog.csdn.net/DouJiangMs/article/details/126066648)
    - 各种算法的特性与使用示例
    - shortuuid
2. [xid](https://github.com/rs/xid)
    - 各种编码算法生成的字符串长度表格对比
3. [为什么没有人用md5(UUID(), 16)来生成16位大小写不敏感的唯一取值？](https://segmentfault.com/q/1010000002949015)

最近有一个需求, 在向数据库中新增记录时, 需要用一个字段作为唯一编号, 要求为16位, 可以是字母数字的组合.

之前对于唯一编号这种需求, 一直用的都是 uuid/md5, 这是第1次遇到有长度限制的需求, 没想到这么麻烦.

| Name        | Binary Size | String Size    |
|-------------|-------------|----------------|
| [UUID]      | 16 bytes    | 36 chars       |
| [shortuuid] | 16 bytes    | 22 chars       |
| [Snowflake] | 8 bytes     | up to 20 chars |
| [MongoID]   | 12 bytes    | 24 chars       |
| [xid]       | 12 bytes    | 20 chars       |

[UUID]: https://en.wikipedia.org/wiki/Universally_unique_identifier
[shortuuid]: https://github.com/stochastic-technologies/shortuuid
[Snowflake]: https://blog.twitter.com/2010/announcing-snowflake
[MongoID]: https://docs.mongodb.org/manual/reference/object-id/
[xid]: https://github.com/rs/xid

## uuid

使用`uuidgen`生成一个uuid串, 如`cb981de1-c10b-4aaa-b842-6ff8b56bae17`.

uuid虽然可以保证唯一, 但是ta的长度是36个字符, 就算移除中间的`-`, 也有32个字符.

------

参考文章3中题主提到了一种"野路子", 先得到一个uuid, 再用md5对uuid字符串进行加密, 如下

```
md5(UUID(), 16) = 148de9519b8bd264;
md5(UUID(), 32) = 7ac66c0f148de9519b8bd264312c4d64
```

因为uuid自身是唯一的, 那么用md5加密后的字符串也应该是唯一的.

...遗憾的是, 字符串的长度是不会骗人的, 你无法把长度为32个字符串的集合放到16个字符的筐里, 还要求ta不重复.

在上面md5的16位和32位两个字符串中, 前者其实是后者的子串, 只不过前后分别截取了8个字符串而已.

另外, md5只是一种摘要算法, **并不能保证唯一**, 有可能两个不同的uuid会得到同一个结果.

## shortuuid

参考文章1和2中提到了`shortuuid`算法, 但这种算法得到的字符串长度也有22个字符.

## xid

按照参考文章3中提到的, xid得到的字符串长度是最短的, 需要20个字符.

这些算法类似于 jwt token 的计算方法, 内容是分段的, 可以从生成的字符串中, 反向解析出源数据.

```
    时间戳    mac地址   pid 有序随机数
|           |        |     |        |
 00 01 02 03 04 05 06 07 08 09 10 11
```

> 上面的 00, 01 等是字节, xid 需要12个字节存储.

## snowflake雪花算法

偶然间了解到的, 由纯数字组成, 与uuid不同的是, ta可以保证递增, 即下一刻生成的结果一定大于上一刻的结果.

但是每个结果需要18个字符.

------

本来想着, 没有长度为16的唯一id字符串, 可以找长度为10, 12, 14的方案, 只要在前缀处加上固定的标记就行, 如SNxxxxxxxxxxxxxx. 但目前来看, 没有找到.
