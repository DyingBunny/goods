package models

import (
	"github.com/astaxie/beego/validation"
	"goods/models/table"
	"goods/pkg/logging/serve"
	"time"
)
//以用户id进行查找
func Find(UserId uint) (table.Login,bool){
	var auth table.Login
	Db.First(&auth,"user_id=?",UserId)
	valid:=validation.Validation{}
	valid.Required(auth.UserId, string(UserId))
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
func AddUser(name string,pwd string,num string,gender string,role string){
	Db.Create(&table.Login{
		Username:name,
		Password:pwd,
		PhoneNumber:num,
		Gender:gender,
		Role:role})
}
//增加买家
func AddBuyer(buyerId uint){
	Db.Create(&table.Buyer{
		BuyerId:buyerId})
}
//填写收货地址
func ModifyAddress(people table.Login,address string,name string,receiveNumber string){
	if people.Role=="buyer"{
		var buyer table.Buyer
		Db.Model(&buyer).Where("buyer_id=?",people.UserId).Update("receive_address",address)
		Db.Model(&buyer).Where("buyer_id=?",people.UserId).Update("name",name)
		Db.Model(&buyer).Where("buyer_id=?",people.UserId).Update("receive_number",receiveNumber)
	}else{
		var seller table.Seller
		Db.Model(&seller).Where("seller_id=?",people.UserId).Update("deliver_address",address)
		Db.Model(&seller).Where("seller_id=?",people.UserId).Update("name",name)
		Db.Model(&seller).Where("seller_id=?",people.UserId).Update("deliver_number",receiveNumber)
	}
}
//增加卖家
func AddSeller(sellerId uint){
	Db.Create(&table.Seller{
		SellerId:sellerId})
}
//增加车主
func AddDriver(sellerId uint){
	Db.Create(&table.Seller{
		SellerId:sellerId})
}
//登录状态更新
func UpdateSta(userid interface{})bool{
	var auth table.Login
	Turnstatus(userid,auth,"True")
	return true
}
//改变登录状态
func Turnstatus(userid interface{},people table.Login,change interface{}){
	Db.Model(&people).Where("user_id=?",userid).Update("login",change)
	Db.Model(&people).Where("user_id=?",userid).Update("last_time",time.Now())
}


