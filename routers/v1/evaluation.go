package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goods/models"
	"goods/pkg/e"
	"net/http"
)

type BuyerEvaluation struct{
	GoodsScore    uint `json:"goods_score"`
	GoodsId  uint `json:"goods_id"`
	SellerId uint `json:"seller_id"`
	BuyerId  uint `json:"buyer_id"`
	DriverId uint `json:"driver_id"`
	Comment  string `json:"comment"`
	SellerScore uint `json:"seller_score"`
	DriverScore uint `json:"driver_score"`
	DistributionId	uint	`json:"distribution_id"`
	OrderId		uint	`json:"order_id"`
}

type GoodsEvaluation struct{
	GoodsId	uint	`json:"goods_id"`
	Page 		int		`json:"page"`
	PageSize	int		`json:"page_size"`
}

//买家评论商品与卖家
func BuyerEvaluate(context *gin.Context){
	var evaluation BuyerEvaluation
	_=context.ShouldBindBodyWith(&evaluation,binding.JSON)
	models.AddEvaluation(evaluation.GoodsScore,evaluation.GoodsId,evaluation.SellerId,evaluation.BuyerId,evaluation.Comment)
	models.BuyerEvaluateSeller(evaluation.SellerScore,evaluation.SellerId)
	models.BuyerEvaluateDriver(evaluation.DriverScore,evaluation.DriverId)
	code:=e.SUCCESS
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
	})
}
//查看商品的评论
func AllGoodsEvaluation(context *gin.Context){
	var goodsEvaluation GoodsEvaluation
	_=context.ShouldBindBodyWith(&goodsEvaluation,binding.JSON)
	goods:=models.FindGoodsEvaluation(goodsEvaluation.GoodsId,goodsEvaluation.Page,goodsEvaluation.PageSize)
	total:=models.FindGoodsEvaNum(goodsEvaluation.GoodsId)
	if len(goods)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"该商品还没有评论",
		})
	}else{
		data:=make(map[string]interface{})
		data["total"]=total
		data["goods"]=goods
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}
////评论司机
//func BuyerEvaluateDriver(context *gin.Context){
//	var evaluation BuyerEvaluation
//	_=context.ShouldBindBodyWith(&evaluation,binding.JSON)
//	models.BuyerEvaluateDriver(evaluation.DriverScore,evaluation.DriverId)
//	code:=e.SUCCESS
//	context.JSON(http.StatusOK,gin.H{
//		"code":code,
//		"msg":e.GetMsg(code),
//	})
//}