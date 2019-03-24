# Postgres存储过程(二)-多返回值

参考文章

1. [SQL优化（四） PostgreSQL存储过程](http://www.jasongj.com/2015/12/27/SQL4_%E5%AD%98%E5%82%A8%E8%BF%87%E7%A8%8B_Store%20Procedure/)

返回多行或多列

## 使用自定义复合类型返回一行多列

PostgreSQL除了支持自带的类型外，还支持用户创建自定义类型。定义一个复合类型，并在函数中返回一个该复合类型的值，从而实现返回一行多列。

```sql
CREATE TYPE compfoo AS (col1 INTEGER, col2 TEXT);
CREATE OR REPLACE FUNCTION getCompFoo(in_col1 INTEGER, in_col2 TEXT) RETURNS compfoo AS 
$$
DECLARE result compfoo;
BEGIN
	result.col1 := in_col1 * 2;
	result.col2 := in_col2 || '_result';
	RETURN result;
END;
$$ LANGUAGE PLPGSQL;
```

> `||`操作符用于拼接字符串.

执行

```sql
SELECT * FROM getCompFoo(1, '1');
 col1 |   col2
------+----------
    2 | 1_result
(1 row)
```

返回多列还有一种比较简单的方法, 就是使用`table`标记: 

```sql
CREATE OR REPLACE FUNCTION getCompFoo(in_col1 INTEGER, in_col2 TEXT) RETURNS table(col1 INTEGER, col2 TEXT) AS
$$
BEGIN
	col1 := in_col1 * 2;
	col2 := in_col2 || '_result';
    -- 注意这里return query, 这种写法是强制的, 必须要有query标记
    -- ...另外我又尝试了一下, 由于在returns处已经声明了返回值名称, 所以下面的return语句是可以省略的...!!!
	RETURN QUERY select col1, col2;
END;
$$ LANGUAGE PLPGSQL;
```

## 使用输出参数名返回一行多列

在声明函数时，除指定输入参数名及类型外，还可同时声明输出参数类型及参数名。此时函数可以输出一行多列。

```sql
CREATE OR REPLACE FUNCTION get2Col(IN in_col1 INTEGER,IN in_col2 TEXT, OUT out_col1 INTEGER, OUT out_col2 TEXT) AS 
$$
BEGIN
	out_col1 := in_col1 * 2;
	out_col2 := in_col2 || '_result';
    -- 这里也没有return语句哦, 如果要return, 请用`return query`
END;
$$ LANGUAGE PLPGSQL;
```

```sql
SELECT * FROM get2Col(1, '1');
 out_col1 | out_col2 
----------+----------
        2 | 1_result
(1 row)
```

## 使用SETOF返回多行记录

实际项目中，存储过程经常需要返回多行记录，可以通过`SETOF`实现。

```sql
CREATE TYPE compfoo AS (col1 INTEGER, col2 TEXT);
CREATE OR REPLACE FUNCTION getSet(rows INTEGER) RETURNS SETOF compfoo AS 
$$
BEGIN
	RETURN QUERY SELECT i * 2, i || '_text'  FROM generate_series(1, rows, 1) as t(i);
END;
$$ LANGUAGE PLPGSQL;
```

> `generate_series`是一个生成器函数, 类似于python中的`range`, 可以生成以指定起始值, 指定步长和最大值的列表.

执行ta.

```sql
SELECT getSet(2);
   getset
------------
 (2,1_text)
 (4,2_text)
(2 rows)

SELECT * FROM getSet(2);
 col1 |  col2
------+--------
    2 | 1_text
    4 | 2_text
(2 rows)
```

本例返回的每一行记录是复合类型，该方法也可返回基本类型的结果集，即多行一列。

## 使用RETURN TABLE返回多行多列

```sql
CREATE OR REPLACE FUNCTION getTable(rows INTEGER) RETURNS TABLE(col1 INTEGER, col2 TEXT) AS 
$$
BEGIN
	RETURN QUERY SELECT i * 2, i || '_text' FROM generate_series(1, rows, 1) as t(i);
END;
$$ LANGUAGE PLPGSQL;
```

```sql
select getTable(2);
  gettable
------------
 (2,1_text)
 (4,2_text)
(2 rows)

select * FROM getTable(2);
 col1 |  col2
------+--------
    2 | 1_text
    4 | 2_text
(2 rows)
```

此时从函数中读取字段就和从表或视图中取字段一样，可以看此种类型的函数看成是带参数的表或者视图。