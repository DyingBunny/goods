package models

import (
	"github.com/astaxie/beego/validation"
	"goods/models/table"
	logging "goods/pkg/logging/serve"
	"strconv"
)
//获取司机table
func FindDriver(UserId string) (table.Driver,bool) {
	var auth table.Login
	Db.First(&auth,"user_id=?",UserId)
	driverId:=auth.RoleId
	var driver table.Driver
	Db.First(&driver,"driver_id=?",driverId)
	valid:=validation.Validation{}
	valid.Required(driver.DriverId, string(driverId))
	if valid.HasErrors(){
		for _,err:=range valid.Errors{
			logging.Info(err.Key,err.Message)
		}
		return table.Driver{},false
	}
	return driver,true
}
//对司机进行评价
func EvaDriver(UserId string,Mark string) bool {
	driver, flag :=FindDriver(UserId)
	if flag==false{
		return false
	}else{
		mark, _ :=strconv.Atoi(Mark)
		driver.Evaluation+=uint(mark)
		driver.Count+=1
		return true
	}
}