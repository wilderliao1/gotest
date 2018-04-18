package demo

import (
	"sync"
	"fmt"
	"time"
)

func Once() {
	o := &sync.Once{}
	go do(o)
	go do(o)

	time.Sleep(2*time.Second)
}

func do(o *sync.Once) {
	fmt.Println("start do")

	o.Do(func() {
		//这句只会执行一次，这就是once达到的效果
		fmt.Println("only once")
	})

	fmt.Println("start end")
}