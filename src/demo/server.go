package demo

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func Server() {
	server1()
	//server2()
	//server3()
}

//第一种方式，实现ServeHTTP接口，，NewServeMux做路由
type handle1 struct {}
type handle2 struct {}

func (h *handle1)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rspMap := make(map[string]interface{})
	rspMap["err_code"] = 0
	rspMap["err_msg"] = "handle1"
	result := make([]interface{},0)
	for i:=0; i < 500; i++{
		temp := make(map[string]interface{})
		temp["type"] = "addline"
		temp["content"] = "sdfsdfsdfsdfsdfsdfsdfsdfjflkasjdlfkjsdlfjsdklfjsldkfjlksdjfklasdjflksdjflk"
		result = append(result, temp)
	}
	rspMap["result"] = result
	rspByte, _ := json.Marshal(rspMap)
	w.Write(rspByte)
}

func (h *handle2)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rspMap := make(map[string]interface{})
	rspMap["err_code"] = 0
	rspMap["err_msg"] = "handle2"
	rspByte, _ := json.Marshal(rspMap)
	w.Write(rspByte)
}

func server1() {
	var h1 handle1
	var h2 handle2
	mux := http.NewServeMux()//路由
	mux.Handle("/path1", &h1)
	mux.Handle("/path2", &h2)
	err := http.ListenAndServe(":8080",mux)
	fmt.Println("ListenAndServe err-->",err)
}

//第二种方式, 使用默认的DefaultServeMux
func server2() {
	var h1 handle1
	var h2 handle2
	http.Handle("/path1", &h1)
	http.Handle("/path2", &h2)
	err := http.ListenAndServe(":8080",nil)
	fmt.Println("ListenAndServe err-->",err)
}

//第三种，文件服务器
func server3() {
	//参数详解
	//第一个/index/，是url中的路径，比如localhost:8080/index/
	//第二个/index/, 应该要和第一个相同，目的是过滤掉url中的/index/。那么为什么要过滤呢，目的是：当url中希望以另外的名字命名路径时，就需要在这里
	//过滤掉这个路径，才能找到对应的文件，举例：
	//要获取当前目录"./"下的文件gotest.go，
	//请求url:localhost:8080/index/gotest.go
	//过滤后：gotest.go, 查找文件时就在"./"目录下查找gotest.go,相当于./gotest.go,这样就能找到对应的文件，这里的index可以随意更换
	//网上详解帖子：https://my.oschina.net/shou1156226/blog/811096
	http.Handle("/index/",http.StripPrefix("/index/", http.FileServer(http.Dir("./"))))
	http.ListenAndServe(":8080",nil)
}