参考文章

1. [olang orm 框架之 gorm](https://segmentfault.com/a/1190000013216540)

2. [GORM文档学习总结](https://blog.csdn.net/wongcony/article/details/79063407)

## 表操作

`HasTable`判断表是否存在.

```go
	// 通过调用db.HasTable来判断是否存在表, 参数可以使用两种形式, 一种是表名的字符串, 一种是模型的地址类型.
	var isExist bool
	isExist = db.HasTable(&User{})
	log.Println(isExist)
	isExist = db.HasTable("users")
	log.Println(isExist)
	isExist = db.HasTable("profiles")
	log.Println(isExist)
```

## 创建

官方文档里新增记录有`Create()`方法, 但是还有一个`NewRecord()`方法, 看得我一脸萌b.

按照参考文章2中的解释, `NewRecord()`是用来判断一个实例对象是否已经插入数据库. 而判断的依据则是其中是否包含`id`属性.

```go
user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
db.NewRecord(user) // => 主键为空返回`true`
db.Create(&user)    // Create()方法会在user对象中加入Id属性
db.NewRecord(user) // => 创建`user`后返回`false`
```

> gorm没有批量创建的方法, 可以手动构建sql语句完成.

## 更新

更新前要先将目标记录查出来. 如已经用`First()`存到`user`对象中了.

`db.Model(&ser).Update[s](xxx)`

`Update`只能更新单个属性, 如`Update(字段名, 字段值)`, 字段名是实际在数据库中的字段, string类型, 而不是ORM中定义的成员属性名.

`Updates`可以更新多个属性, 它只接受两种类型的参数, `map`和`struct`, 并且注意不能是指针类型, 只能为对象的引用.

使用`struct`时, 注意gorm会自动忽略其中的空值, 比如, 如果其中某个字段为bool类型, 而它的是false, 那gorm就不会更新这个字段, 同理, 如果某个字符串字段的值为空, 也会被忽略. 使用map作为参数就不会有此问题.

## Scan查询

```go
type User struct {
	Name string
	Age uint8
}
```

一般与`Select()`配合使用, 当我们只想要查询某张表的指定字段时, 其他字段是没有用处的, 那么常规的`Find(&users)`其实有点得不偿失的感觉, 尤其是当这张表中的字段很多时, 结果列表可能占用较大内存. 

此时我们可以使用`Scan()`方法, 仅将我们需要的列放入结果中即可.

```go
type NameResult struct{
	Name string
}
nameResults := []*NameResult{}
db.Model(&User{}).Select("name").Scan(&nameResults)
```

> 注意: 如果使用`First()`或`Find()`去查询`&nameResults`的值什么也不会得到, 猜测是因为与`Model()`的参数类型不同. 就算没有Model, 也会因为数据库中根本不存在`name_results`表而查询失败.