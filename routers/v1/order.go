package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goods/models"
	"goods/models/table"
	"goods/pkg/e"
	"net/http"
)

type PurchaseBuyer struct{
	GoodsId		uint	`json:"goods_id"`
	SellerId 	uint	`json:"seller_id"`
	BuyerId 	uint	`json:"buyer_id"`
	Num 		uint	`json:"num"`
}

type GoodsBuyer struct{
	BuyerId	uint	`json:"Buyer_id"`
	Page 		int		`json:"page"`
	PageSize	int		`json:"page_size"`
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

//查看买家目前的订单数量
func BuyerOrderNum(context *gin.Context){
	var order table.Order
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	one:=models.FindBuyerOrderPay(order.BuyerId)
	two:=models.FindBuyerOrderDeli(order.BuyerId)
	three:=models.FindBuyerOrderComple(order.BuyerId)
	data:=make(map[string]interface{})
	data["to_be_paid"]=one
	data["to_be_delivered"]=two
	data["completed"]=three
	code:=e.SUCCESS
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}


//查看买家的配送中订单
func AllBuyerOrderPay(context *gin.Context){
	var order GoodsBuyer
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	orders:=models.FindBuyerAllOrderPay(order.BuyerId,order.Page,order.PageSize)
	total:=models.FindBuyerOrderPay(order.BuyerId)
	if len(orders)==0{
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"用户没有待付款的订单",
		})
	}else{
		data:=make(map[string]interface{})
		data["total"]=total
		data["orders"]=orders
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}

//查看买家的配送中订单
func AllBuyerOrderDeli(context *gin.Context){
	var order GoodsBuyer
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	orders:=models.FindBuyerAllOrderDeliv(order.BuyerId,order.Page,order.PageSize)
	total:=models.FindBuyerOrderDeli(order.BuyerId)
	if len(orders)==0{
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"用户没有配送中的订单",
		})
	}else{
		data:=make(map[string]interface{})
		data["total"]=total
		data["orders"]=orders
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}

//查看买家的已完成订单
func AllBuyerOrderComplete(context *gin.Context){
	var order GoodsBuyer
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	orders:=models.FindBuyerAllOrderComplete(order.BuyerId,order.Page,order.PageSize)
	total:=models.FindBuyerOrderComple(order.BuyerId)
	if len(orders)==0{
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"用户没有完成的订单",
		})
	}else{
		data:=make(map[string]interface{})
		data["total"]=total
		data["orders"]=orders
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}