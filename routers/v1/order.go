package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goods/models"
	"goods/pkg/e"
	"net/http"
)

type PurchaseBuyer struct{
	GoodsId		uint	`json:"goods_id"`
	SellerId 	uint	`json:"seller_id"`
	BuyerId 	uint	`json:"buyer_id"`
	Num 		uint	`json:"num"`
}

//买家购买商品
func PurchaseCommodities(context *gin.Context){
	var purchase PurchaseBuyer
	_=context.ShouldBindBodyWith(&purchase,binding.JSON)
	good:=models.FindCommodity(purchase.GoodsId)
	if purchase.Num<good.RemainNum {
		models.AddOrder(purchase.GoodsId,purchase.SellerId,purchase.BuyerId,purchase.Num)
		models.AdjustRemainNum(good.GoodsId,purchase.Num)
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
		})
	}else if purchase.Num==good.RemainNum{
		models.AddOrder(purchase.GoodsId,purchase.SellerId,purchase.BuyerId,purchase.Num)
		models.AdjustRemainNum(good.GoodsId,purchase.Num)
		models.LowCommodity(good.GoodsId)
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
		})
	}else{
		code:=e.ERROR
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"商品库存不足",
		})
	}
}