package models

import (
	"github.com/astaxie/beego/validation"
	"goods/models/table"
	"goods/pkg/logging/serve"
)
//以用户id进行查找
func Find(UserId string) (table.Login,bool){
	var auth table.Login
	Db.First(&auth,"user_id=?",UserId)
	valid:=validation.Validation{}
	valid.Required(auth.UserId,UserId)
	if valid.HasErrors(){
		for _,err:=range valid.Errors{
			logging.Info(err.Key,err.Message)
		}
		return table.Login{},false
	}
	return auth,true
}
//检查登录信息
func Check(username string,password string) bool{
	var auth table.Login
	Db.First(&auth,"Username=? AND Password=?",username,password)
	valid:=validation.Validation{}
	valid.Required(auth.Username,username)
	if valid.HasErrors(){
		for _,err:=range valid.Errors{
			logging.Info(err.Key,err.Message)
		}
		return false
	}
	return true
}
//登录状态更新
func UpdateSta(userid interface{})bool{
	var auth table.Login
	Turnstatus(userid,auth,"True")
	return true
}
//用户名和密码查找id
func FindUserId(username string,password string)uint{
	var people table.Login
	Db.First(&people,"Username=? AND Password=?",username,password)
	return people.UserId
}
//用户名和密码查找
func FindRole(username string,password string)string{
	var people table.Login
	Db.First(&people,"Username=? AND Password=?",username,password)
	return people.Role
}

//增加用户
func AddUser(name string,pwd string,gender string,email string,role string){
	Db.Create(&table.Login{
		Username:name,
		Password:pwd,
		Gender:gender,
		Email:email,
		Role:role})
}

//改变登录状态
func Turnstatus(userid interface{},people table.Login,change interface{}){
	Db.Model(&people).Where("user_id=?",userid).Update("Login",change)
}
