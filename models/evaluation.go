package models

import "goods/models/table"

//增加商品评论
func AddEvaluation(score uint,goodsId uint,sellerId uint,buyerId uint,comment string){
	var auth table.Login
	Db.First(&auth,"user_id=?",buyerId)
	Db.Create(&table.Evaluation{
		Score:score,
		GoodsId:goodsId,
		SellerId:sellerId,
		BuyerId:buyerId,
		BuyerUsername:auth.Username,
		Comment:comment})
}
//对卖家进行评分
func BuyerEvaluateSeller(score uint,sellerId uint){
	var seller table.Seller
	Db.Where("seller_id=?",sellerId).First(&seller)
	newEvaluation:=score+seller.Evaluation
	newCount:=seller.Count+1
	newComprehensive:=newEvaluation/newCount
	Db.Where("seller_id=?",sellerId).First(&seller).Update("evaluation",newEvaluation)
	Db.Where("seller_id=?",sellerId).First(&seller).Update("count",newCount)
	Db.Where("seller_id=?",sellerId).First(&seller).Update("comprehensive",newComprehensive)
}
//对司机进行评分
func BuyerEvaluateDriver(score uint,driverId uint){
	var driver table.Driver
	Db.Where("driver_id=?",driverId).First(&driver)
	newEvaluation:=score+driver.Evaluation
	newCount:=driver.Count+1
	newComprehensive:=newEvaluation/newCount
	Db.Where("driver_id=?",driverId).First(&driver).Update("evaluation",newEvaluation)
	Db.Where("driver_id=?",driverId).First(&driver).Update("count",newCount)
	Db.Where("driver_id=?",driverId).First(&driver).Update("comprehensive",newComprehensive)
}
//查看商品评论
func FindGoodsEvaluation(goodsId uint,page int,pageSize int)[]table.Evaluation{
	var goodsEvaluation []table.Evaluation
	Db.Where("goods_id=?",goodsId).Limit(pageSize).Offset((page-1)*pageSize).Find(&goodsEvaluation)
	return goodsEvaluation
}
//查看商品评论的数量
func FindGoodsEvaNum(goodsId uint)int{
	var total=0
	Db.Model(&table.Evaluation{}).Where("goods_id=?",goodsId).Count(&total)
	return total
}