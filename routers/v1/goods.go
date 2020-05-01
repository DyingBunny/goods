package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "golang.org/x/tools/godoc"
	"goods/models"
	"goods/models/table"
	"goods/pkg/e"
	"net/http"
	_ "time"
)

type GoodsSeller struct{
	SellerId	uint	`json:"seller_id"`
	Page 		int		`json:"page"`
	PageSize	int		`json:"page_size"`
}

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
			good.Price,good.TransPrice,good.DeliverAddress,good.Detail)
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
	data["state"]=good.State
	data["detail"]=good.Detail
	data["identification"]=good.Identification
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}

//查看一个卖家目前的商品上架与下架数量
func SellerComNum(context *gin.Context){
	var good GoodsSeller
	_=context.ShouldBindBodyWith(&good,binding.JSON)
	up:=models.FindSellerComUp(good.SellerId)
	low:=models.FindSellerComLow(good.SellerId)
	data:=make(map[string]interface{})
	data["up"]=up
	data["low"]=low
	code:=e.SUCCESS
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}

//查看卖家发布的所有上架商品
func AllSellerComProUp(context *gin.Context){
	var good GoodsSeller
	_=context.ShouldBindBodyWith(&good,binding.JSON)
	goods:=models.FindSellerAllComUp(good.SellerId,good.Page,good.PageSize)
	total:=models.FindSellerComUp(good.SellerId)
	if len(goods)==0{
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"该用户尚未发布商品",
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

//查看卖家发布的所有下架商品
func AllSellerComProLow(context *gin.Context){
	var good GoodsSeller
	_=context.ShouldBindBodyWith(&good,binding.JSON)
	goods:=models.FindSellerAllComLow(good.SellerId,good.Page,good.PageSize)
	total:=models.FindSellerComLow(good.SellerId)
	if len(goods)==0{
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"该用户尚未发布商品",
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

//买家查看已上架的商品，按照价格降序
func ViewItemDesc(context *gin.Context){
	var good GoodsSeller
	_=context.ShouldBindBodyWith(&good,binding.JSON)
	goods:=models.FindAllComByPriceDesc(good.Page,good.PageSize)
	total:=models.FindAllComPageNum()
	if len(goods)==0{
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"还未有买家发布商品",
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

//买家查看已上架的商品，按照价格升序
func ViewItemAsc(context *gin.Context){
	var good GoodsSeller
	_=context.ShouldBindBodyWith(&good,binding.JSON)
	goods:=models.FindAllComByPriceAsc(good.Page,good.PageSize)
	total:=models.FindAllComPageNum()
	if len(goods)==0{
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"还未有买家发布商品",
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
//卖家主动下架商品
func LowCommodity(context *gin.Context){
	var good table.Goods
	_=context.ShouldBindBodyWith(&good,binding.JSON)
	models.LowCommodity(good.GoodsId)
	code:=e.SUCCESS
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
	})
}