	var a int
	var b int32
	var c int64

	fmt.Printf("%d\n", unsafe.Sizeof(a)) // 8
	fmt.Printf("%d\n", unsafe.Sizeof(b)) // 4
	fmt.Printf("%d\n", unsafe.Sizeof(c)) // 8

	fmt.Printf("%d\n", unsafe.Sizeof(&a)) // 8
