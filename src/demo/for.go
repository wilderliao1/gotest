package demo

import "fmt"

func For() {
	//iMap := make(map[string]interface{})
	//iMap["a"] = 1
	//iMap["b"] = 2
	//iMap["c"] = 3
	//
	//fmt.Println("--------------------key:value")
	////key:value
	//for key, value := range iMap {
	//	fmt.Printf("key:%s,value:%d\n",key,value)
	//}
	//
	//fmt.Println("--------------------key")
	////只要key
	//for key := range iMap {
	//	fmt.Printf("key:%s\n",key)
	//}
	//
	//fmt.Println("--------------------value")
	////只要value
	//for _, value := range iMap {
	//	fmt.Printf("value:%d\n",value)
	//}
	//
	//str := "这是一个字符串，abc"
	//for pos, char := range str {
	//	fmt.Printf("pos：%d, char: %#U \n", pos, char)
	//}
	//fmt.Println(str[0])
	//fmt.Println(str[1])
	//fmt.Println(str[2])
	//fmt.Println(str[3])

	//deleteArr()
	//result := getA()
	//fmt.Println("result:",result)
	dd()
}


func deleteArr() {
	oriArr := []int{1,2,3,4,5,6,7,8,9,10}//原数组
	delA(oriArr)
	//fmt.Println("oriArr--->:",oriArr)
}

func delA(oriArr []int) {
	var result []int //满足条件的数组
	for index, value := range  oriArr {
		fmt.Println("value:",value)
		//if value == 3 {
		//	result = append(result, value)
		//	continue
		//}
		if value == 5 || value == 8 {//删除5和8
			//fmt.Println("index:",index)
			oriArr = append(oriArr[:index],oriArr[index+1:]...)
		} else {
			result = append(result, value)
		}
	}

	fmt.Println("oriArr:",oriArr)
	fmt.Println("result:",result)
}

func dd() {
	var result []int //满足条件的数组
	oriArr := []int{1,2,3,4,5,6,7,8,9,10}//原数组
	for i:=0; i<len(oriArr); i++{
		value := oriArr[i]
		if value == 5 || value == 8 {//删除5和8
			//fmt.Println("index:",index)
			oriArr = append(oriArr[:i],oriArr[i+1:]...)
			i--
		} else {
			result = append(result, value)
		}
	}
	fmt.Println("oriArr:",oriArr)
	fmt.Println("result:",result)
}

func getA() ([]int){
	oriArr := []int{1,2,3,4,5,6,7,8,9,10}//原数组
	var result []int //满足条件的数组
	for index, value := range  oriArr {
		if value == 3 {
			result = append(result, value)
			continue
		}
		if value == 5 || value == 8 {//删除5和8
			fmt.Println("index:",index)
			oriArr = append(oriArr[:index],oriArr[index+1:]...)
		}
	}

	fmt.Println("oriArr:",oriArr)
	fmt.Println("result:",result)
	return result
}