package controller

import (
	"fmt"
	"go_02/api/webapp/bookstore0612/dao"
	"html/template"
	"net/http"
)

//Login处理用户登录的函数
func Login(w http.ResponseWriter,r *http.Request){
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//调用userdao中验证用户名和密码的方法
	user,_ := dao.CheckUserNameAndPassword(username,password)
	fmt.Println("获取user是:",user)
	if user.ID != 0{
		//用户名和密码正确
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w,"")
	}else{
		//用户名或密码不正确
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w,"登录失败，请检查输入的用户名和密码。")
	}
}
//完成注册接口
func Regist(w http.ResponseWriter,r *http.Request){
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//调用userdao中验证用户名和密码的方法
	user,_ := dao.CheckUserName(username)
	fmt.Println("register user is:",user)
	if user.ID > 0{
		//用户名已经存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w,"用户名已存在！请重新输入。")
	}else{
		//正常注册逻辑
		err := dao.SaveUser(username, password, email)
		fmt.Println("注册报错是:",err)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w,"insert user err!")
	}
}
//用户名输入框验证
func FinduserByname(w http.ResponseWriter,r *http.Request){
	//获取用户名和密码
	username := r.PostFormValue("username")
	//调用userdao中验证用户名和密码的方法
	user,_ := dao.CheckUserName(username)
	fmt.Println("find user by name:",user)
	if user.ID > 0{
		//用户名已经存在
		w.Write([]byte("<font style='color:red'>用户名已经存在!</font>"))
	}else{
		//用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用!</font>"))
	}
}
