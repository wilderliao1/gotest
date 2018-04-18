package demo

import (
	"fmt"
)

func Slice() {
	iSlice := make([]int, 10)
	modifySlice(iSlice)//切片作为参数传递到函数中，在函数中改变之后会体现在调用处(函数内的改变会影响外部的值)
	for _,value := range iSlice {
		fmt.Println("value-->",value)
	}

	//获取0到数组结尾的所有数据
	fmt.Println(iSlice[0:])

	//获取len(iSlice)到数组开头的所有数据
	fmt.Println(iSlice[:len(iSlice)])

	//获取[2,5)范围内的数据
	fmt.Println(iSlice[2:5])

	fmt.Println("--------------------测试切片赋值")
	sliceFromSlice()

	fmt.Println("--------------------测试切片容量")
	sliceCap()

	fmt.Println("--------------------Append测试")
	sliceAppend()
}

func modifySlice(slice []int) {
	for index,_ := range slice {
		slice[index] = index
	}
}

func sliceFromSlice() {
	slice1 := []int{1,2,3}
	slice2 := slice1

	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)

	slice1[0] = 5

	//输出结果表明，修改slice1, 同时也修改了slice2
	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)
}

func sliceCap() {
	slice1 := []int{1,2,3}
	fmt.Println("cap(slice1)=",cap(slice1))
	//slice1[4] = 4 //直接赋值会crash，索引超出范围了
	//fmt.Println("cap(slice1)=",cap(slice1))
}

func sliceAppend() {
	slice1 := []int{1,2,3}
	fmt.Println("slice1:",slice1)
	slice1 = append(slice1,4)//append不会crash，内部会自动增加扩容
	fmt.Println("append slice1:",slice1)

	var slice2 []int
	fmt.Println("slice2:",slice2)
	slice2 = append(slice2,1)//可以直接往空数组里面添加元素
	fmt.Println("append slice2:",slice2)

	//一个数组追加到另一个数组
	slice3 := []int{1,2,3}
	slice4 := []int{4,5,6}
	slice3 = append(slice3, slice4...)
	fmt.Println("slice3+4:",slice3)
}