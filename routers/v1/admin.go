package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goods/models"
	"goods/models/table"
	"goods/pkg/e"
	"net/http"
)

//商品管理

//管理员查看所有上架商品
func AdminFindAllGoodsUp(context *gin.Context){
	var good GoodsSeller
	_=context.ShouldBindBodyWith(&good,binding.JSON)
	goods:=models.AdminFindAllComUp(good.Page,good.PageSize)
	total:=models.AdminFindAllComNumUp()
	if len(goods)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"还未有卖家发布商品",
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

//管理员查看所有下架商品
func AdminFindAllGoodsLow(context *gin.Context){
	var good GoodsSeller
	_=context.ShouldBindBodyWith(&good,binding.JSON)
	goods:=models.AdminFindAllComLow(good.Page,good.PageSize)
	total:=models.AdminFindAllComNumLow()
	if len(goods)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"还未有下架的商品",
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

//管理员下架商品
func AdminLowCommodity(context *gin.Context){
	var good table.Goods
	_=context.ShouldBindBodyWith(&good,binding.JSON)
	models.LowCommodity(good.GoodsId)
	code:=e.SUCCESS
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
	})
}

//订单管理

//管理员查看待付款订单
func AdminFindAllOrderPay(context *gin.Context){
	var order GoodsSeller
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	orders:=models.AdminFindBAllOrderPay(order.Page,order.PageSize)
	total:=models.AdminFindAllOrderNumPay()
	if len(orders)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"没有配送中的订单",
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
//管理员查看正在配送订单
func AdminFindAllOrderDel(context *gin.Context){
	var order GoodsSeller
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	orders:=models.AdminFindBAllOrderDel(order.Page,order.PageSize)
	total:=models.AdminFindAllOrderNumDel()
	if len(orders)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"没有配送中的订单",
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
//管理员查看已完成订单
func AdminFindAllOrderCom(context *gin.Context){
	var order GoodsSeller
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	orders:=models.AdminFindBAllOrderCom(order.Page,order.PageSize)
	total:=models.AdminFindAllOrderNumCom()
	if len(orders)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"没有配送中的订单",
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
//管理员改变订单状态
func AdminChangeOrderState(context *gin.Context){
	var order table.Order
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	models.AdminChangeOrderState(order.OrderId,order.State)
	code:=e.SUCCESS
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
	})
}

//配送订单管理

//管理员查看所有配送中的配送订单
func AdminFindDistributionDel(context *gin.Context){
	var order GoodsSeller
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	distribution:=models.AdminFindAllDistributionDel(order.Page,order.PageSize)
	total:=models.AdminFindAllDistributionNumDel()
	if len(distribution)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"没有配送中的的配送订单",
		})
	}else{
		data:=make(map[string]interface{})
		data["total"]=total
		data["distribution"]=distribution
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}
//管理员查看所有已完成的配送订单
func AdminFindDistributionCom(context *gin.Context){
	var order GoodsSeller
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	distribution:=models.AdminFindAllDistributionCom(order.Page,order.PageSize)
	total:=models.AdminFindAllDistributionNumCom()
	if len(distribution)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"没有已完成的的配送订单",
		})
	}else{
		data:=make(map[string]interface{})
		data["total"]=total
		data["distribution"]=distribution
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}

//管理员改变配送订单状态
func AdminChangeDistributionState(context *gin.Context){
	var distribution table.Distribution
	_=context.ShouldBindBodyWith(&distribution,binding.JSON)
	models.AdminChangeDistributionState(distribution.DistributionId,distribution.State)
	code:=e.SUCCESS
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
	})
}

//用户信息管理
func AdminFindAllUser(context *gin.Context){
	var order GoodsSeller
	_=context.ShouldBindBodyWith(&order,binding.JSON)
	users:=models.AdminFindAllUser(order.Page,order.PageSize)
	total:=models.AdminFindAllUserNum()
	if len(users)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"没有已注册的用户",
		})
	}else{
		data:=make(map[string]interface{})
		data["total"]=total
		data["users"]=users
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}