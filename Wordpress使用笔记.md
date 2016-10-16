# Wordpress使用笔记

## 1. 设置文件上传目录

版本: 4.3.6

访问`http://generals.space/wp-admin/options.php`

找到`upload_path`字段(默认为空), 修改成想要的路径即可. 注意, 需要是绝对路径

> PS: 关于`upload_url_path`

> 这是一个url, 作为上传路径下的文件的访问url.

> 在http://generals.space/wp-admin/options.php的设置中默认为空, 但实际上wordpress将它的值使用为http://generals.space/wp-content/uploads/

> 若将它设置为http://uploads.jiangming7.com/, 而上传路径下(一般是uploads目录下)有文件a.jpg. 则上传文件的访问路径将由http://generals.space/wp-content/uploads/a.jpg改为http://uploads.jiangming7.com/a.jpg

## 2. wordpress迁移后图片无法显示

解决办法:

首先登录mysql数据库, 图片地址一般存于`wp_posts`表的`post_content`字段中, 即文章表的**内容字段**(为了安全起见, 可先用查询语句查看其中图片链接与实际图片地址的对应关系);

然后使用MySQL的Update Replace语句, 将post_content字段中出现图片链接的部分替换成新的地址

```sql
update wp_posts set post_content = replace(post_content, '原内容', '新内容');

//例如, 将图片的本地链接全部修改为远程的绝对地址
update wp_posts set post_content = replace(post_content, 'http://localhost/wp-content/uploads', 'http://blog.jiangming7.cn/wp-content/uploads') where post_content like '%http://localhost/wp-content/uploads%';
```

## 3. `站点地址url`与`wordpress地址url`写错导致无法进入后台

这两个地址一般存在于`wp_options`表中, 前者在`wp_options`表中的`option_name`字段中的值为`siteurl`, 后者的`option_name`字段的值为`home`, 可用`select`语句查询

```sql
select * from wp_options where option_name = 'siteurl';
select * from wp_options where option_name = 'home';
```

站点地址是要出现在浏览器地址栏中的访问地址, 后者则是网页正文中出现的链接中的地址, 可用update语句直接修改.

至于wordpress实际的安装地址(如使用blog文件夹存放的子站点), 只需要在http服务器中配置. 如

nginx需要在`server`指令中指明`root`的值为`/var/www/html/blog`;

apache则需要在`VirtualHost`标签中指明`DocumentRoot`标签的值为`/var/www/html/blog`

## 4. 文章列表页面, 图片的缩略图无法显示

参考文章

[TimThumb Troubleshooting Secrets](https://www.binarymoon.co.uk/2010/11/timthumb-hints-tips/)

[Timthumb can't show image after it's uploaded](http://stackoverflow.com/questions/14396991/timthumb-cant-show-image-after-its-uploaded)

F12得到缩略图的url, 在新标签页中打开, 页面输出:

```
A TimThumb error has occured

The following error(s) occured:
Could not create the index.html file - to fix this create an empty file named index.html file in the cache directory.
Could not create cache clean timestamp file.


Query String : src=/wp-content/uploads/2015/06/234241430-1.png&h=140&w=205&zc=1
TimThumb version : 2.8.13
可在wordpress目录下locate/find timthumb文件, 可能存在于主题目录下.
```

当时选用的解决办法:

将`$主题目录/cache/index`的权限改为666;

将**wordpress**, 或者**当前主题目录**所属用户/所属组改为apache:apache

再次访问缩略图url将可以正常访问