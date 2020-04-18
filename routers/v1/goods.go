package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goods/models"
	"goods/models/table"
	"goods/pkg/e"
	"net/http"
)

//卖家发布商品
func ReleaseCommodities(context *gin.Context){
	var good table.Goods
	_=context.ShouldBindBodyWith(&good,binding.JSON)
	_,isExist:=models.Find(good.SellerId)
	if isExist==false{
		code:=e.ERROR
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
		})
	}else{
		code:=e.SUCCESS
		a:=models.AddCommodities(good.Text,good.SellerId,good.SellerUsername,good.TotalNum,
			good.Price,good.TransPrice,good.DeliverAddress)
		data := make(map[string]interface{})
		goodsId:=models.FindCommodities(good.SellerId,a)
		data["goods_id"]=goodsId
		models.AddSellerGoods(good.SellerId,goodsId)
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}

//查看商品
func CommoditiesProfile(context *gin.Context){
	var good table.Goods
	_=context.ShouldBindBodyWith(&good,binding.JSON)
	models.Db.First(&good,"goods_id=?",good.GoodsId)
	code:=e.SUCCESS
	data := make(map[string]interface{})
	data["goods_id"]=good.GoodsId
	data["text"]=good.Text
	data["seller_id"]=good.SellerId
	data["seller_username"]=good.SellerUsername
	data["total_num"]=good.TotalNum
	data["remain_num"]=good.RemainNum
	data["price"]=good.Price
	data["trans_price"]=good.TransPrice
	data["deliver_address"]=good.DeliverAddress
	data["create_time"]=good.CreateTime
	data["identification"]=good.Identification
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}