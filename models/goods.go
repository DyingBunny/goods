package models

import (
	"goods/models/table"
	"strconv"
	"time"
)

//增加商品
func AddCommodities(text string,sellerId uint,sellerUsername string,
	totalNum uint,price uint,transPrice uint, deliverAddress string,detail string)string{
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
		CreateTime:time.Now(),
		State:1,
		Detail:detail})
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
	Db.Where("seller_id=? AND identification=?",sellerId,identification).First(&good)
	return good.GoodsId
}

//查看卖家发布的上架商品
func FindSellerAllComUp(sellerId uint,page int,pageSize int)[]table.Goods{
	var goods []table.Goods
	Db.Where("seller_id=? AND state=?",sellerId,1).Limit(pageSize).Offset((page-1)*pageSize).Find(&goods)
	return goods
}

//查看卖家发布的下架商品
func FindSellerAllComLow(sellerId uint,page int,pageSize int)[]table.Goods{
	var goods []table.Goods
	Db.Where("seller_id=? AND state=?",sellerId,0).Limit(pageSize).Offset((page-1)*pageSize).Find(&goods)
	return goods
}

//查看卖家发布的上架商品的数量
func FindSellerComUp(sellerId uint) int{
	var total =0
	Db.Model(&table.Goods{}).Where("seller_id=? AND state=?",sellerId,1).Count(&total)
	return total
}

//查看卖家发布的下架商品的数量
func FindSellerComLow(sellerId uint) int{
	var total =0
	Db.Model(&table.Goods{}).Where("seller_id=? AND state=?",sellerId,0).Count(&total)
	return total
}

//价格降序查看所有商品
func FindAllComByPriceDesc(page int,pageSize int)[]table.Goods{
	var goods []table.Goods
	Db.Where("state=?",1).Limit(pageSize).Offset((page-1)*pageSize).Order("price desc").Find(&goods)
	return goods
}
//价格升序查看所有商品
func FindAllComByPriceAsc(page int,pageSize int)[]table.Goods{
	var goods []table.Goods
	Db.Where("state=?",1).Limit(pageSize).Offset((page-1)*pageSize).Order("price asc").Find(&goods)
	return goods
}

//查看所有商品的数量
func FindAllComPageNum() int {
	var total=0
	Db.Model(&table.Goods{}).Count(&total)
	return total
}

//id查找商品结构
func FindCommodity(goodsId uint)table.Goods{
	var good table.Goods
	Db.Where("goods_id=?",goodsId).First(&good)
	return good
}

//商品减去数量
func AdjustRemainNum(goodsId uint,num uint){
	var good table.Goods
	var tmp table.Goods
	Db.Where("goods_id=?",goodsId).First(&tmp)
	result:=tmp.RemainNum-num
	Db.Model(&good).Where("goods_id=?",goodsId).Update("remain_num",result)
}

//商品下架
func LowCommodity(goodsId uint){
	var good table.Goods
	Db.Where("goods_id=?",goodsId).First(&good).Update("state",0)
}