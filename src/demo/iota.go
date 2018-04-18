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

// iota第一个值赋值给空白标识符，这样可以忽略第一个值
type ByteSize float64
const (
	// 通过赋予空白标识符来忽略第一个值
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
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

	fmt.Println("KB ---->  ", KB)
	fmt.Println("MB ---->  ", MB)
	fmt.Println("GB ---->  ", GB)
	fmt.Println("TB ---->  ", TB)
	fmt.Println("PB ---->  ", PB)
	fmt.Println("EB ---->  ", EB)
}