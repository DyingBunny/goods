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

//查看买家待付款订单的数量
func FindBuyerOrderPay(buyerId uint)int{
	var total=0
	Db.Model(&table.Order{}).Where("buyer_id=? AND state=?",buyerId,1).Count(&total)
	return total
}
//查看买家配送中订单的数量
func FindBuyerOrderDeli(buyerId uint)int{
	var total=0
	Db.Model(&table.Order{}).Where("buyer_id=? AND state=?",buyerId,2).Count(&total)
	return total
}
//查看买家已完成订单的数量
func FindBuyerOrderComple(buyerId uint)int{
	var total=0
	Db.Model(&table.Order{}).Where("buyer_id=? AND state=?",buyerId,3).Count(&total)
	return total
}
//查看买家待付款的订单
func FindBuyerAllOrderPay(buyerId uint,page int,pageSize int)[]table.Order{
	var orders []table.Order
	Db.Where("buyer_id=? AND state=?",buyerId,1).Limit(pageSize).Offset((page-1)*pageSize).Find(&orders)
	return orders
}
//查看买家配送中的订单
func FindBuyerAllOrderDeliv(buyerId uint,page int,pageSize int)[]table.Order{
	var orders []table.Order
	Db.Where("buyer_id=? AND state=?",buyerId,2).Limit(pageSize).Offset((page-1)*pageSize).Find(&orders)
	return orders
}
//查看买家已完成的订单
func FindBuyerAllOrderComplete(buyerId uint,page int,pageSize int)[]table.Order{
	var orders []table.Order
	Db.Where("buyer_id=? AND state=?",buyerId,3).Limit(pageSize).Offset((page-1)*pageSize).Find(&orders)
	return orders
}