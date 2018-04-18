package demo

import (
	"fmt"
	"unsafe"
)

func Int()  {
	fmt.Println("")
	i := int(1)
	fmt.Println("size(int):",unsafe.Sizeof(i))

	j := uint(1)
	fmt.Println("size(uint):",unsafe.Sizeof(j))
}