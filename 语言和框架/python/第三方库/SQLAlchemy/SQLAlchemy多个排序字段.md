# SQLAlchemy多个排序字段

参考文章

1. [postgresql – 多个order_by sqlalchemy/Flask](https://codeday.me/bug/20180419/154608.html)

```py
User.query.order_by(User.popularity.desc(),User.date_created.desc()).limit(10).all()
```