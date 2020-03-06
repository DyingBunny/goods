package models

import (
	"goods/models/table"
	"time"
)

//增加新订单
func AddOrder(text string,createTime time.Time,completeTime time.Time,price uint, distance float64){
	Db.Create(&table.Order{
		Text:text,
		CreateTime:createTime,
		CompleteTime:completeTime,
		Price:price,
		Distance:distance})
}
