# for..range中的值传递

```go
package main

import "log"

type student struct {
	Name string
	Age  int
}

func main() {
	result := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		stu.Age += 10
		result[stu.Name] = &stu
	}

	for _, stu := range stus {
		log.Printf("stu: %+v\n", stu)
	}
	// 对Age的修改没有影响到stus对象中的成员.
	// 2019/05/04 21:12:33 stu: {Name:zhou Age:24}
	// 2019/05/04 21:12:33 stu: {Name:li Age:23}
	// 2019/05/04 21:12:33 stu: {Name:wang Age:22}

	for name, stu := range result {
		log.Printf("name: %s, addr: %d, info: %+v\n", name, &stu, stu)
	}
	// 2019/05/05 00:48:28 name: zhou, addr: 824633745456, info: &{Name:wang Age:32}
	// 2019/05/05 00:48:28 name: li, addr: 824633745456, info: &{Name:wang Age:32}
	// 2019/05/05 00:48:28 name: wang, addr: 824633745456, info: &{Name:wang Age:32}
}

```

注意:

1. `for..range..`是值拷贝, `stu`变量是一个成员副本, 所以通过`stu`修改的属性无法影响到`stus`中的成员.
2. 循环中只创建一个新对象, 每次为这个对象赋予不同的值, 所以`&stu`的地址是相同的.

> 可以使用索引完成对 slice/map 原数据的修改.
