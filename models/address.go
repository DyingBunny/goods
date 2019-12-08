package models

import (
	"github.com/astaxie/beego/validation"
	"goods/models/table"
	logging "goods/pkg/logging/serve"
)

//查找用户的发货位置
func FindDelAddress(UserId string)(table.Address,bool){
	var deliver table.DeliverMapping
	Db.First(&deliver,"pid=?",UserId)
	addressId:=deliver.Cid
	var address table.Address
	Db.First(&address,"address_id=?",addressId)
	valid:=validation.Validation{}
	valid.Required(address.AddressId, string(addressId))
	if valid.HasErrors(){
		for _,err:=range valid.Errors{
			logging.Info(err.Key,err.Message)
		}
		return table.Address{},false
	}
	return address,true
}
//查找用户的收货位置
func FindRecAddress(UserId string)(table.Address,bool){
	var receive table.ReceiveMapping
	Db.First(&receive,"pid=?",UserId)
	addressId:=receive.Cid
	var address table.Address
	Db.First(&address,"address_id=?",addressId)
	valid:=validation.Validation{}
	valid.Required(address.AddressId, string(addressId))
	if valid.HasErrors(){
		for _,err:=range valid.Errors{
			logging.Info(err.Key,err.Message)
		}
		return table.Address{},false
	}
	return address,true
}
