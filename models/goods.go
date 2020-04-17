package models

import (
	"goods/models/table"
	"time"
)

//增加商品
func AddCommodities(text string,sellerId uint,sellerUsername string,
	totalNum uint,price uint,transPrice uint, deliverAddress string,
	receiveAddress string){
	Db.Create(&table.Goods{
		Text:text,
		SellerId:sellerId,
		SellerUsername:sellerUsername,
		TotalNum:totalNum,
		RemainNum:totalNum,
		Price:price,
		TransPrice:transPrice,
		DeliverAddress:deliverAddress,
		ReceiveAddress:receiveAddress,
		CreateTime:time.Now()})
}
