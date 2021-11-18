package main

import (
	"fmt"
	"net/http"
)

//创建处理函数
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"hello world!",r.URL.Path)
}

func main(){
	http.HandleFunc("/",handler)
	//创建路由
	http.ListenAndServe(":8090",nil)
}
