package main

import (
	"fmt"
	"net/http"
)

//创建处理器函数
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"测试http协议;;;")
}

func main(){
	//调用处理器处理请求
	http.HandleFunc("/http",handler)
	//路由
	http.ListenAndServe(":8090",nil)
	//请求报文响应信息：
	//Response Headers　相应头中的时间
	/*
	2xx成功表示请求已经被成功接收，理解，接受
	3xx重定向－－要完成请求必须进行更进一步的处理
	４xx客户端错误－－请求有语法错误或请求无法复现
	5xx服务器端错误--服务器未能实现合法的请求
	*/

}
