# Postgres存储过程(三)-动态sql

参考文章

1. [SQL优化（四） PostgreSQL存储过程](http://www.jasongj.com/2015/12/27/SQL4_%E5%AD%98%E5%82%A8%E8%BF%87%E7%A8%8B_Store%20Procedure/)

2. [存储过程里面动态查询sql语句，如何转义单引号？](https://bbs.csdn.net/topics/340016154)

有时在`PL/pgSQL`函数中需要生成动态命令，这个命令将包括他们每次执行时使用不同的表或者字符。`EXECUTE`语句用法如下

```
EXECUTE command-string [ INTO [STRICT] target] [USING expression [, ...]];
```

此时`PL/plSQL`将不再缓存该命令的执行计划。相反，在该语句每次被执行的时候，命令都会编译一次。这也让该语句获得了对各种不同的字段甚至表进行操作的能力。

`command-string`包含了要执行的命令，它可以使用参数值，在命令中通过引用如`$1`，`$2`等来引用参数值。这些符号的值是指`USING`字句的值。这种方法对于在命令字符串中使用参数是最好的：它能避免运行时数值从文本来回转换，并且不容易产生`SQL`注入，而且它不需要引用或者转义。

首先创建测试表

```sql
CREATE TABLE testExecute AS
SELECT
	i || '' AS a,
	i AS b
FROM generate_series(1, 10, 1) AS t(i);
```

> 这么清奇的建表方式我还是头一次见...

## 使用`using`传入变量

```sql
CREATE OR REPLACE FUNCTION execute(filter TEXT) RETURNS TABLE (a TEXT, b INTEGER) AS 
$$
BEGIN
	RETURN QUERY EXECUTE
		'SELECT * FROM testExecute where a = $1'
	USING filter;
END;
$$ LANGUAGE PLPGSQL;
```

执行查询

```sql
SELECT * FROM execute('3');
 a | b
---+---
 3 | 3
(1 row)
```

我们看看会不会被sql注入

```sql
SELECT * FROM execute('3'' or ''c''=''c');
 a | b
---+---
(0 rows)
```

完全没有风险.

当然，也可以使用**字符串拼接**的方式在command-string中使用参数，但会有SQL注入的风险。

## 使用字符串拼接传入变量

```sql
CREATE OR REPLACE FUNCTION execute(filter TEXT) RETURNS TABLE (a TEXT, b INTEGER) AS 
$$
BEGIN
	RETURN QUERY EXECUTE 'SELECT * FROM testExecute where a = ''' || filter || '''';
END;
$$ LANGUAGE PLPGSQL;
```

```sql
# SELECT * FROM execute('3');
 a | b
---+---
 3 | 3
(1 row)
```

功能相同, 但是容易被sql注入.

```sql
# SELECT * FROM execute('3'' or ''c''=''c');
 a  | b
----+----
 1  |  1
 2  |  2
 3  |  3
 4  |  4
 5  |  5
 6  |  6
 7  |  7
 8  |  8
 9  |  9
 10 | 10
(10 rows)

```

从该例中可以看出使用字符串拼接的方式在`command-string`中使用参数会引入SQL注入攻击的风险，而使用`USING`的方式则能有效避免这一风险。

## sql语句中单引号转义

这个问题其实在上面通过字符串拼接传入变量就已经解决了. 在`execute`语句中, 两个单引号`''`就可以表示原生单引号, 作用与其他语言中`\'`相同.

例如, 普通的sql语言写为`select * from my_table where name = 'general'`.

在存储过程中写的话要写成`execute 'select * from my_table where name = ''general'' '`.

所以, 再仔细看看上面的sql注入是如何发生的?

## 关于`execute`中的`into`

```sql
CREATE OR REPLACE FUNCTION execute(filter TEXT) RETURNS TABLE (a TEXT, b INTEGER) AS 
$$
BEGIN
	EXECUTE
		'SELECT * FROM testExecute where a = $1' into a, b
	USING filter;
    raise notice '%, %', a, b;
    return query select a, b;
END;
$$ LANGUAGE PLPGSQL;
```

尝试一下.

```sql
SELECT * FROM execute('3');
NOTICE:  3, 3
 a | b
---+---
 3 | 3
(1 row)
```
