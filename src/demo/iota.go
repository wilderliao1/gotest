package demo

import "fmt"

// 常量本身不会自增长
const (
	const_1 = 0
	const_2
	const_3
)

// iota实现递增
const (
	iota_1 = iota//初始值是0
	iota_2
	iota_3
)

// iota实现指数递增
const (
	index_1 = 1 << iota
	index_2
	index_3
)

// iota初始值是1， 在const中是将该值带入代数式中计算，可以达到自定义递增的效果
func Iota() {
	fmt.Println("const_1  ---->  ", const_1)
	fmt.Println("const_2  ---->  ", const_2)
	fmt.Println("const_3  ---->  ", const_3)


	fmt.Println("iota_1  ---->  ", iota_1)
	fmt.Println("iota_2  ---->  ", iota_2)
	fmt.Println("iota_3  ---->  ", iota_3)

	fmt.Println("index_1 ---->  ", index_1)
	fmt.Println("index_2 ---->  ", index_2)
	fmt.Println("index_3 ---->  ", index_3)
}