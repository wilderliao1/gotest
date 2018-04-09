package demo

import "fmt"

func For() {
	iMap := make(map[string]interface{})
	iMap["1"] = 1
	iMap["2"] = 2
	iMap["3"] = 3

	isT := false
	for key, _ := range iMap {
		if !isT {
			if key == "2" {
				fmt.Println("--->2")
				continue
				fmt.Println("--->111")
			}
		}
		fmt.Println("--->",key)
	}
}