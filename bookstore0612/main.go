package main

import (
	"go_02/api/webapp/bookstore0612/controller"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter,r *http.Request){
	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w,"")
}

func main(){
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("views/static"))))
	//直接去首页
	http.Handle("/pages/",http.StripPrefix("/pages/",http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/main",IndexHandler)
	//去登录
	http.HandleFunc("/login",controller.Login)
	//注册
	http.HandleFunc("/register",controller.Regist)
	//注册时候校验用户名
	http.HandleFunc("/FindUserByName",controller.FinduserByname)
	http.ListenAndServe(":8090",nil)
}
