# Hexo基础应用

参考文章

1. [Hexo 入门指南（三） - 文章 & 草稿](http://blog.csdn.net/wizardforcel/article/details/40684575?_t_t_t=0.7924863273750762)

2. [Github 搭建 hexo （五）- 站点地图（sitemap.xml）](http://blog.csdn.net/u010053344/article/details/50706790)

## 1. 创建标签及分类.

当在一篇文章头部写下如下内容时, `hexo g`操作就会在public目录下创建`tags`及`categories`目录. 

```md
---
title: Linux-ACL应用
tags: [ACL]
categories: general
---
```

tags目录会包含ACL目录, 而categories则会包含general目录. 但没有tags与categories子目录下都没有index.html主页面.

```
.
├── categories
│   └── general
│       └── index.html
└── tags
    └── ACL
        └── index.html

```

这种情况下, 点击页面上的tags链接将会得到404错误. 因为tags下面没有索引, 除非指定访问某个tag如/tags/ACL.

![](https://gitimg.generals.space/d8217e078ceceb8ef1901d64ea40f0a3.png)

所以我们还需要创建标签页和分类页的索引目录.

创建标签目录

```
$ hexo new page tags
```

它会在source目录下创建tags子目录, 并且在tags子目录下创建`index.md`文件. `/tags/index.md`文件内容如下

```md
---
title: tags
date: 2017-06-25 11:06:14
---
```

你还需要在这个index.md文件中添加一行`type: 'tags'`.

等到`hexo g`生成页面时, 你会发现`/public/tags`目录下出现了一个index.html文件.

```
└── tags
    ├── ACL
    │   └── index.html
    └── index.html
```

再次访问`/tags`将会看到如下页面.

![](https://gitimg.generals.space/6f56313d4840c5b9303738c79165ce8e.png)

注意: 一定要添加`type: 'tags'`字段, 不然标签索引页看不到可用标签, 只有一个空页面. 如下

![](https://gitimg.generals.space/f73f0660766da55d7505af6b4f8ca59c.png)

同理, 分类页categories也是如此.

```
$ hexo new page categories
```

然后在`/source/categories/index.md`文件中添加`type: 'categories'`.

再次生成页面即可.

## 2. 站点地图

安装站点地图生成器

```
$ npm install hexo-generator-sitemap --save
$ npm install hexo-generator-baidu-sitemap --save
```

保证`$HEXO/_config.yml`文件中有如下配置(默认存在)

```
# sitemap
sitemap:
  path: sitemap.xml
baidusitemap:
  path: baidusitemap.xml
```

之后执行`hexo g`就可以在public子目录下生成`sitemap.xml`与`baidusitemap.xml`两个地图文件了. 也就是说, 每次更新源文件就会重新生成地图.