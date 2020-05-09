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
	BuyerId	uint	`json:"buyer_id"`
	Page 		int		`json:"page"`
	PageSize	int		`json:"page_size"`
}

//*****买家*****\\

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
		code:=e.DISPLAY
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
		code:=e.DISPLAY
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
		code:=e.DISPLAY
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
		code:=e.DISPLAY
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
//买家对一个订单付款
func BuyerPay(context *gin.Context){
	var order table.Order
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	models.BuyerPay(order.OrderId)
	code:=e.SUCCESS
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
	})
}
//买家确认收货
func BuyerComplete(context *gin.Context){
	var order BuyerEvaluation
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	if models.ConfirmAllDistribution(order.OrderId)==1{
		models.BuyerComplete(order.OrderId)
		models.AddEvaluation(order.GoodsScore,order.GoodsId,order.SellerId,order.BuyerId,order.Comment)
		models.BuyerEvaluateSeller(order.SellerScore,order.SellerId)
		models.BuyerEvaluateDriver(order.DriverScore,order.DriverId)
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
		})
	}else{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"您有尚未确认收货的配送订单",
		})
	}
}

//*****卖家*****\\

//看卖家目前的订单数量
func BuyerSellerNum(context *gin.Context){
	var order table.Order
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	one:=models.FindSellerOrderPay(order.SellerId)
	two:=models.FindSellerOrderDeli(order.SellerId)
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


//查看卖家的买家待付款订单
func AllBuyerSellerPay(context *gin.Context){
	var order GoodsSeller
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	orders:=models.FindSellerAllOrderPay(order.SellerId,order.Page,order.PageSize)
	total:=models.FindSellerOrderPay(order.SellerId)
	if len(orders)==0{
		code:=e.DISPLAY
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

//查看卖家的配送中订单
func AllBuyerSellerDeli(context *gin.Context){
	var order GoodsSeller
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	orders:=models.FindSellerAllOrderDeliv(order.SellerId,order.Page,order.PageSize)
	total:=models.FindSellerOrderDeli(order.SellerId)
	if len(orders)==0{
		code:=e.DISPLAY
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

//查看卖家的已完成订单
func AllBuyerSellerComplete(context *gin.Context){
	var order GoodsSeller
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	orders:=models.FindSellerAllOrderComplete(order.SellerId,order.Page,order.PageSize)
	total:=models.FindSellerOrderComple(order.SellerId)
	if len(orders)==0{
		code:=e.DISPLAY
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

//*****司机*****\\

//司机查看订单
func DriverFindAllOrder(context *gin.Context){
	var order GoodsSeller
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	orders:=models.DriverFindAllOrder(order.Page,order.PageSize)
	total:=models.AllOrderNum()
	if len(orders)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"尚未有待配送订单",
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
