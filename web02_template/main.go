package main

import (
	"html/template"
	"net/http"
)

//创建处理器函数
func testTemplate(w http.ResponseWriter,r *http.Request){
	//解析模板
	//t,_ := template.ParseFiles("index.html")
	//通过Must函数让go帮助我们自动处理异常
	t:= template.Must(template.ParseFiles("/home/SENSETIME/dongchao_vendor/go/src/go_02/api/webapp/web02_template/index.html","/home/SENSETIME/dongchao_vendor/go/src/go_02/api/webapp/web02_template/index2.html"))
	//执行
	//t.Execute(w,"Hello Template")
	//将相应数据在index2.html文件中显示
	t.ExecuteTemplate(w,"index2.html","我要去index2.html中")
}

func main(){
	http.HandleFunc("/testTemplate",testTemplate)
	http.ListenAndServe(":8090",nil)
}
