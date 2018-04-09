package demo

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func Server() {
	//server1()
	server2()
}

//第一种方式，实现ServeHTTP接口，，NewServeMux做路由
type handle1 struct {}
type handle2 struct {}

func (h *handle1)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rspMap := make(map[string]interface{})
	rspMap["err_code"] = 0
	rspMap["err_msg"] = "handle1"
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