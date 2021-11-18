package dao

import (
	"go_02/api/webapp/bookstore0612/model"
	"go_02/api/webapp/bookstore0612/utils"
)

//CheckUserNameAndPassword 根据用户名和密码find
func CheckUserNameAndPassword(username string,password string)(*model.User,error){
	//写sql语句
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr,username,password)
	user := &model.User{}
	row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	return user,nil
}

//CheckUserName 根据用户名和密码从数据库中查出一条记录
func CheckUserName(username string)(*model.User,error){
	//写sql语句
	sqlStr := "select id,username,password,email from users where username = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr,username)
	user := &model.User{}
	row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	return user,nil
}

//SaveUser数据库中插入用户信息
func SaveUser(username string,password string,email string)error{
	//写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//执行
	_,err := utils.Db.Exec(sqlStr,username,password,email)
	if err != nil{
		return err
	}
	return nil
}