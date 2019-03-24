# Postgres存储过程(一)入门

参考文章

1. [SQL优化（四） PostgreSQL存储过程](http://www.jasongj.com/2015/12/27/SQL4_%E5%AD%98%E5%82%A8%E8%BF%87%E7%A8%8B_Store%20Procedure/)

2. [PostgreSQL pl/pgsql 编写存储过程](https://abcdkyd.github.io/2018/06/15/postgres%20pl-pgsql%20%E7%BC%96%E5%86%99%E5%AD%98%E5%82%A8%E8%BF%87%E7%A8%8B/)

在网上搜索存储过程相关资料时, 有一个很热门的条目是**存储过程与函数的区别**. 虽然从题目上看, 两者肯定是有区别的, 但是既然有这个话题, 就说明两者有很多的相似之处, 以致于很多人无法辨别. 所以我们可以先把存储过程简单地理解为一种特殊的函数.

最初接触存储过程是要完成类似mysql中`on update CURRENT_TIMESTAMP`标记的, 更新某行数据时自动更新`updated_at`时间字段的操作, postgres里没有这种标记, 所以借助了存储过程+触发器来实现.

然后是在写后端业务逻辑时, 有一个停车场占用率历史的查询相当复杂, 需要查询某天的1, 3, 5, 7...23点的占用率, 返回一个数组给前端. 数据库表中只存储了不同时间停车场各车位的使用情况, 我写出的sql语句可以查询单个时间点时的占位, 比如1点时的占位个数, 但是要循环查询12次才能得到全部数据. 所以当时就找了下postgres中循环查询的方法, 一直追到存储过程才解决.

## PostgreSQL支持的过程语言

PostgreSQL官方支持`PL/pgSQL`，`PL/Tcl`，`PL/Perl`和`PL/Python`这几种过程语言。同时还支持一些第三方提供的过程语言，如`PL/Java`，`PL/PHP`，`PL/Py`，`PL/R`，`PL/Ruby`，`PL/Scheme`，`PL/sh`。

## 基于SQL的存储过程定义

```sql
CREATE OR REPLACE FUNCTION add(a INTEGER, b NUMERIC) RETURNS NUMERIC AS 
$$
	SELECT a+b;
$$ LANGUAGE SQL;
```

调用方法

```sql
SELECT add(1,2);
 add
-----
   3
(1 row)
SELECT * FROM add(1,2);
 add
-----
   3
(1 row)
```

上面这种方式参数列表只包含函数输入参数，不包含输出参数。下面这个例子将同时包含输入参数和输出参数

```sql
CREATE OR REPLACE FUNCTION plus_and_minus (IN a INTEGER, IN b NUMERIC, OUT c NUMERIC, OUT d NUMERIC) AS 
$$
	SELECT a+b, a-b;
$$ LANGUAGE SQL;
```

调用方式

```sql
SELECT plus_and_minus(3,2);
 plus_and_minus
----------------
 (5,1)
(1 row)
SELECT * FROM plus_and_minus(3,2);
 c | d
---+---
 5 | 1
(1 row)
```

该例中，`IN`代表输入参数，`OUT`代表输出参数。这个带输出参数的函数和之前的`add`函数并无本质区别。事实上，输出参数的最大价值在于**它为函数提供了返回多个字段的途径**。

在函数定义中，可以写多个`SQL`语句，不一定是`SELECT`语句，可以是其它任意合法的`SQL`, 比如`insert`语句。

## 基于PL/PgSQL的存储过程定义

PL/pgSQL是一个块结构语言。函数定义的所有文本都必须是一个块。一个块用下面的方法定义：

```
[ <<label>> ]
[DECLARE
	declarations]
BEGIN
	statements
END [ label ];
```

- 中括号部分为可选部分

- 块中的每一个`declaration`和每一条`statement`都由一个分号终止

- 块支持嵌套，嵌套时子块的`END`后面必须跟一个分号，最外层的块`END`后可不跟分号

- `BEGIN`后面不必也不能跟分号

- `END`后跟的label名必须和块开始时的标签名一致

- 所有关键字都不区分大小写。标识符被隐含地转换成小写字符，除非被双引号包围

- 声明的变量在当前块及其子块中有效，子块开始前可声明并覆盖（只在子块内覆盖）外部块的同名变量

- 变量被子块中声明的变量覆盖时，子块可以通过外部块的label访问外部块的变量

使用`PL/PgSQL`语言的函数定义如下：

```sql
CREATE FUNCTION somefunc() RETURNS integer AS $$
DECLARE
   -- 声明变量
	quantity integer := 30;
BEGIN
	-- 打印变量值
	RAISE NOTICE 'Quantity here is %', quantity;
	quantity := 50;
	RAISE NOTICE 'Quantity here is %', quantity;
   RETURN quantity;
END;
$$ LANGUAGE plpgsql;
```

执行ta

```sql
# select somefunc();
NOTICE:  Quantity here is 30
NOTICE:  Quantity here is 50
 somefunc
----------
       50
(1 row)
```

> `raise notice`在存储过程调试中很有帮助.

### 声明函数参数

如果只指定输入参数类型，不指定参数名，则函数体里一般用`$1`，`$n`这样的标识符来使用参数。

```sql
CREATE OR REPLACE FUNCTION discount(NUMERIC) RETURNS NUMERIC AS 
$$
BEGIN
	RETURN $1 * 0.8;
END;
$$ LANGUAGE PLPGSQL;
```

但该方法可读性不好，此时可以为`$n`参数声明别名，然后可以在函数体内通过别名指向该参数值。

```sql
CREATE OR REPLACE FUNCTION discount(NUMERIC) RETURNS NUMERIC AS 
$$
DECLARE
	total ALIAS FOR $1;
BEGIN
	RETURN total * 0.8;
END;
$$ LANGUAGE PLPGSQL;
```

当然上述方法仍然不够直观, 平常当然是用参数名+参数类型来表示的.

```sql
CREATE OR REPLACE FUNCTION discount(total NUMERIC) RETURNS NUMERIC AS 
$$
BEGIN
	RETURN total * 0.8;
END;
$$ LANGUAGE PLPGSQL;
```

执行ta

```sql
# select discount(12);
 discount
----------
      9.6
(1 row)
# select * from discount(12);
 discount
----------
      9.6
(1 row)
```

## 存储过程的查看与删除

`\df`: 可以查看当前数据库下的所有存储过程列表.

`\sf 存储过程名称`: 查看指定存储过程的详细语句.

`\ef 存储过程名称`: 编辑存储过程代码.

使用`drop function 存储过程名称`即可删除指定名称的存储过程.

...貌似在postgres 10.x的时候要使用`drop function 存储过程名称()`?

## 一点感想-关于要不要使用存储过程的取舍

存储过程要看dba给不给你创建的权限, 如果没有权限就需要dba创建后程序调用.

一般不建议搞存储过程, 后续维护是天坑. 比如存储过程调试很烦, 修改的时候容易因为各种各样的原因造成阻塞、对象无效等等.

而且数据库本身只应该做一个存储的功能, 其他不管是要扩展、维护还是解读, 都是脱离于数据库正常使用范畴的.

由于是dba应该也不想搞让开发人员搞存储过程给自己挖坑, 不好维护.