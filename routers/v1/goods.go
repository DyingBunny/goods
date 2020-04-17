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
		models.AddCommodities(good.Text,good.SellerId,good.SellerUsername,good.TotalNum,
			good.Price,good.TransPrice,good.DeliverAddress, good.ReceiveAddress)
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
		})
	}
}

//查看商品
