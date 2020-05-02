package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goods/models"
	"goods/models/table"
	"goods/pkg/e"
	"net/http"
)

type BuyerEvaluation struct{
	GoodsScore    uint `json:"goods_score"`
	GoodsId  uint `json:"goods_id"`
	SellerId uint `json:"seller_id"`
	BuyerId  uint `json:"buyer_id"`
	Comment  string `json:"comment"`
	SellerScore uint `json:"seller_score"`
}

//买家评论商品与卖家
func BuyerEvaluate(context *gin.Context){
	var evaluation BuyerEvaluation
	_=context.ShouldBindBodyWith(&evaluation,binding.JSON)
	models.AddEvaluation(evaluation.GoodsScore,evaluation.GoodsId,evaluation.SellerId,evaluation.BuyerId,evaluation.Comment)
	models.BuyerEvaluateSeller(evaluation.SellerScore,evaluation.SellerId)
	code:=e.SUCCESS
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
	})
}
//买家评论卖家
func BuyerEvaluateSeller(context *gin.Context){
	var evaluation table.Evaluation
	_=context.ShouldBindBodyWith(&evaluation,binding.JSON)
	models.BuyerEvaluateSeller(evaluation.Score,evaluation.SellerId)
	code:=e.SUCCESS
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
	})
}