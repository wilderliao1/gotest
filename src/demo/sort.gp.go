package demo

import "fmt"

func quickS(values []int, left int, right int) {
	if left >= right {
		return
	}
	temp := values[left]
	i, j := left, right

	for {
		//从右到左找到第一个小于基准值的数
		for values[j] >= temp && i < j  {
			j--
		}

		//从左到右找到第一个大于基准值的数
		for values[i] <= temp && i < j  {
			i++
		}

		if i >= j {
			break
		}

		values[i], values[j] = values[j], values[i]
	}

	values[left] = values[i]
	values[i] = temp

	quickS(values,left,i)
	quickS(values,i+1,right)
}

func QuickSort() {
	slice := []int{7,2,1,8,4}
	quickS(slice,0, len(slice)-1)
	fmt.Println("slice:",slice)
}