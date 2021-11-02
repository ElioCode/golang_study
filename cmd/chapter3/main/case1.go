package main

import (
	"golang_study/cmd/chapter4/a"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", a.SayhelloName)     //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
