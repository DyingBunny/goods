package models

import "goods/models/table"

//增加商品评论
func AddEvaluation(score uint,goodsId uint,sellerId uint,buyerId uint,comment string){
	Db.Create(&table.Evaluation{
		Score:score,
		GoodsId:goodsId,
		SellerId:sellerId,
		BuyerId:buyerId,
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
//