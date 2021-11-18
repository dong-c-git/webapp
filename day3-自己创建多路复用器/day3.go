package main

import (
	"fmt"
	"net/http"
)
//创建处理器函数
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"!",r.URL.Path)
}


func main(){
	//创建多路复用器
	mux := http.NewServeMux()
	mux.HandleFunc("/",handler)
	//创建路由
	http.ListenAndServe(":8090",mux)
	//浏览器和服务器之间通信需要经历４个步骤
	//建立连接
	//发出请求信息
	//回送响应信息
	//关闭连接
	//1.0版本中，浏览器请求一个带有图片的网页，会由于下载图片与服务器之间开启一个新的连接
	//在http1.1版本中，允许浏览器在拿到当前请求对应的全部资源后断开连接提高了效率；

}

