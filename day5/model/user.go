package model

import (
	"fmt"
	"go_02/api/webapp/day5/utils"
)

//User 结构体
type User struct {
	ID int
	Username string
	Password int
	Email string
}

//AddUser　添加Ｕser的方法
func (user *User)AddUser()error{
	//1.写sql语句
	sqlstr := "insert into user(name,age) values(?,?)"
	//2.预编译
	inStmt,err := utils.Db.Prepare(sqlstr)
	if err != nil{
		fmt.Println("预编译出现异常:",err)
		return err
	}
	//3.执行
	_,err2 := inStmt.Exec("admin","123456")
	if err2 != nil{
		fmt.Println("执行出现异常",err2)
		return err
	}
	return nil
}


//AddUser　添加Ｕser的方法　不使用预编译方式
func (user *User)AddUser2()error{
	//1.写sql语句
	sqlstr := "insert into user(name,age) values(?,?)"
	//2.执行
	_,err := utils.Db.Exec(sqlstr,"admin3",44566)
	if err != nil{
		fmt.Println("执行出现异常:",err)
		return err
	}
	return nil
}

//查询一条记录GetUserById　根据用户的id从数据库中查询一条记录
func (user *User)GetUserByID()(*User,error){
	//写sql语句
	sqlStr := "select id,name,age from user where id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr,user.ID)
	//声明
	var id int
	var name string
	var age int
	err := row.Scan(&id,&name,&age)
	if err != nil{
		return nil,err
	}
	u := &User{
		ID: id,
		Username: name,
		Password: age,
	}
	return u ,nil
}

//GetUsers 获取数据库中所有的记录
func (user *User)GetUsers()([]*User,error){
	//写sql语句
	//写sql语句
	sqlStr := "select id,name,age from user"
	//执行
	rows, err := utils.Db.Query(sqlStr)//查询多行记录
	if err != nil{
		return nil,err
	}
	//创建users切片
	var users []*User
	for rows.Next(){
		//声明
		var id int
		var name string
		var age int
		err := rows.Scan(&id,&name,&age)
		if err != nil{
			return nil,err
		}
		u := &User{
			ID: id,
			Username: name,
			Password: age,
		}
		users = append(users, u)
	}
	return users,nil
}