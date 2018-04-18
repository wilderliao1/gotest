package main

import (
	"demo"
	"runtime"
	"fmt"
)

var _ = demo.Iota // 用于调试，结束时删除。
var _ = runtime.NumCPU()
var _,_ = fmt.Println("")

func main() {
	//demo.Iota()
	//demo.For()
	//demo.Server()
	//demo.Once()
	//demo.Slice()
	//fmt.Println("runtime.NumCPU() = ",runtime.NumCPU())
	//demo.QuickSort()
	//demo.Chan_test()

	demo.PBTest()

}
