package models

import (
	"goods/models/table"
	"time"
)

//增加订单
func AddOrder(goodsId uint,sellerId uint,buyerId uint,num uint){
	var good table.Goods
	var buyer table.Buyer
	Db.Where("goods_id=?",goodsId).First(&good)
	Db.Where("buyer_id=?",buyerId).First(&buyer)
	Db.Create(&table.Order{
		GoodsId:goodsId,
		SellerId:sellerId,
		BuyerId:buyerId,
		Text:good.Text,
		CreateTime:time.Now(),
		TotalNum:num,
		RemainNum:num,
		TotalPrice:num*good.Price,
		TotalTransPrice:num*good.TransPrice,
		DeliverAddress:good.DeliverAddress,
		ReceiveAddress:buyer.ReceiveNumber,
		State:1})
}

