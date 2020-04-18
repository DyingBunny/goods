package models

import (
	"goods/models/table"
	"strconv"
	"time"
)

//增加商品
func AddCommodities(text string,sellerId uint,sellerUsername string,
	totalNum uint,price uint,transPrice uint, deliverAddress string)string{
		a:=strconv.FormatInt(time.Now().Unix(),10)
	Db.Create(&table.Goods{
		Text:text,
		SellerId:sellerId,
		SellerUsername:sellerUsername,
		TotalNum:totalNum,
		RemainNum:totalNum,
		Price:price,
		TransPrice:transPrice,
		DeliverAddress:deliverAddress,
		Identification:a,
		CreateTime:time.Now()})
		return a
}

//插入卖家-商品
func AddSellerGoods(sellerId uint,goodsId uint){
	Db.Create(&table.SellerGoods{
		SellerId:sellerId,
		GoodsId:goodsId})
}

//查找商品id
func FindCommodities(sellerId uint,identification string)uint{
	var good table.Goods
	Db.Where("seller_id=? AND identification=?",sellerId,identification).Find(&good)
	return good.GoodsId
}