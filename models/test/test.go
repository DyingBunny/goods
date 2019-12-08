package test

import (
	"goods/models"
	"goods/models/table"
	"time"
)

func AddLogin(name string,pwd string,num string,gender string,email string,role string){
	models.Db.Create(&table.Login{
		Username:name,
		Password:pwd,
		PhoneNumber:num,
		Gender:gender,
		Email:email,
		Role:role})
}

func AddDriver(addid uint){
	models.Db.Create(&table.Driver{
		AddressId:addid})
}

func AddReceiveMapping(pid uint,cid uint){
	models.Db.Create(&table.ReceiveMapping{
		Pid:pid,
		Cid:cid})
}

func AddDeliverMapping(pid uint,cid uint){
	models.Db.Create(&table.DeliverMapping{
		Pid:pid,
		Cid:cid})
}

func AddAddress(address string,lng float64,lat float64){
	models.Db.Create(&table.Address{
		Address:address,
		Lng:lng,
		Lat:lat})
}

func AddOrder(text string,createTime time.Time,completeTime time.Time,price uint, distance float64){
	models.Db.Create(&table.Order{
		Text:text,
		CreateTime:createTime,
		CompleteTime:completeTime,
		Price:price,
		Distance:distance})
}

func TestAdd(){
	AddLogin("dyingbunny","Qq594112087","13772005470",
		"male","dyingbunny9@gmail.com","buyer")
	AddDriver(1)
	AddReceiveMapping(1,2)
	AddDeliverMapping(1,2)
	AddAddress("陕西省西安工业大学未央校区",109.11379651119838,34.53572956652158)
	AddOrder("meikuangyunshu",time.Now(),time.Now(),500,8.5)
}
