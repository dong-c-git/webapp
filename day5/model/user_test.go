package model

import (
	"fmt"
	"testing"
)

//TestMain函数可以在测试函数执行之前做一些其他相关方法
func TestMain(m *testing.M){
	fmt.Println("测试开始:test Main")
	//通过m.Run()来运行测试函数
	m.Run()
}

func TestUser(t *testing.T){
	fmt.Println("开始测试User中相关方法:test user")
	//通过t.Run运行子函数
	t.Run("测试添加用户:",testAddUser)
	t.Run("查询所有用户",testGetUsers)
}

//如果函数名不是用Test开头的，那么函数默认不执行，可以将它设置成一个子测试函数
func testAddUser(t *testing.T) {
	fmt.Println("测试添加用户: test addUser")
	fmt.Println("子测试函数。。。")
	//user := &User{}
	//user.AddUser()
}

func testGetUsers(t *testing.T){
	fmt.Println("测试查询所有记录")
	user := &User{}
	us,_ := user.GetUsers()
	//遍历切片
	for k,v := range us{
		fmt.Printf("第%v个用户是%v: \n",k+1,v)
	}
}