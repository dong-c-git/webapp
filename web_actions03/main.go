package main

import (
	"html/template"
	"net/http"
)

func testIf(w http.ResponseWriter,r *http.Request){
	//解析模板文件
	t := template.Must(template.ParseFiles("if.html"))
	age := 17
	//执行
	t.Execute(w,age > 18)

}

func testwith(w http.ResponseWriter,r *http.Request){
	//解析模板文件
	t := template.Must(template.ParseFiles("with.html"))
	//执行
	t.Execute(w,"毛毛")
}

func testTemplate(w http.ResponseWriter,r *http.Request){
	//解析模板文件
	t := template.Must(template.ParseFiles("template1.html","template2.html"))
	t.Execute(w,"测试数据")
}
func testDefine(w http.ResponseWriter,r *http.Request){
	//解析模板文件
	t := template.Must(template.ParseFiles("define.html"))
	//执行
	t.ExecuteTemplate(w,"model","")
}
func main(){
	http.HandleFunc("/testIf",testIf)
	http.HandleFunc("/testwith",testwith)//修改传过来的数据
	http.HandleFunc("/testTemplate",testTemplate)
	http.HandleFunc("/testdefine",testDefine)
	http.ListenAndServe(":8090",nil)
}
