package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goods/models"
	"goods/models/table"
	"goods/pkg/e"
	"net/http"
)

type DistributeDriver struct{
	OrderId 	uint	`json:"order_id"`
	DriverId 	uint	`json:"driver_id"`
	Num 		uint	`json:"num"`
}

type DistributeSellerBuyer struct{
	SellerId	uint	`json:"seller_id"`
	BuyerId	uint	`json:"buyer_id"`
	Page 		int		`json:"page"`
	PageSize	int		`json:"page_size"`
}

//司机接单配送
func DriverDistribute(context *gin.Context){
	var distributeDriver DistributeDriver
	_=context.ShouldBindBodyWith(&distributeDriver,binding.JSON)
	var order table.Order
	models.Db.Where("order_id=?",distributeDriver.OrderId).First(&order)
	if order.RemainNum>=distributeDriver.Num{
		total:=models.DriverFindDistributionNum(distributeDriver.DriverId)
		if total==0{
			models.AddDistribution(distributeDriver.OrderId,distributeDriver.DriverId,distributeDriver.Num)
			models.AdjustOrderNum(distributeDriver.OrderId,distributeDriver.Num)
			models.AddressUpdateDelRec(distributeDriver.OrderId,distributeDriver.DriverId)
			code:=e.SUCCESS
			context.JSON(http.StatusOK,gin.H{
				"code":code,
				"msg":e.GetMsg(code),
			})
		}else{
			code:=e.ERROR
			context.JSON(http.StatusOK,gin.H{
				"code":code,
				"msg":"您有未完成的订单",
			})
		}
	}else{
		code:=e.ERROR
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"待配送数量不足",
		})
	}
}

//司机查看自己的配送订单
func DriverFindDistribution(context *gin.Context){
	var distributeDriver DistributeDriver
	_=context.ShouldBindBodyWith(&distributeDriver,binding.JSON)
	distribution:=models.DriverFindDistribution(distributeDriver.DriverId)
	if len(distribution)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"您尚未有正在配送中的订单",
		})
	}else{
		data:=make(map[string]interface{})
		data["distribution"]=distribution
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}
//买家确认配送订单
func BuyerConfirmDistribution(context *gin.Context){
	var distribution BuyerEvaluation
	_=context.ShouldBindBodyWith(&distribution,binding.JSON)
	models.ConfirmDistribution(distribution.DistributionId)
	models.BuyerEvaluateDriver(distribution.DriverScore,distribution.DriverId)
	if models.ConfirmAllDistribution(distribution.OrderId)==1 {
		models.BuyerComplete(distribution.OrderId)
	}
	code:=e.SUCCESS
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
	})
}

//卖家查看司机的配送状态
func SellerFindDistribution(context *gin.Context){
	var distribution DistributeSellerBuyer
	_=context.ShouldBindBodyWith(&distribution,binding.JSON)
	distributions:=models.SellerFindDistribution(distribution.SellerId,distribution.Page,distribution.PageSize)
	total:=models.SellerFindDistributionNum(distribution.SellerId)
	if len(distributions)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"没有配送中的订单",
		})
	}else{
		data:=make(map[string]interface{})
		data["total"]=total
		data["distributions"]=distributions
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}
//卖家查看已完成的配送
func SellerFindDistributionCom(context *gin.Context){
	var distribution DistributeSellerBuyer
	_=context.ShouldBindBodyWith(&distribution,binding.JSON)
	distributions:=models.SellerFindDistributionCom(distribution.SellerId,distribution.Page,distribution.PageSize)
	total:=models.SellerFindDistributionComNum(distribution.SellerId)
	if len(distributions)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"没有配送中的订单",
		})
	}else{
		data:=make(map[string]interface{})
		data["total"]=total
		data["distributions"]=distributions
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}

//买家查看司机的配送状态
func BuyerFindDistribution(context *gin.Context){
	var distribution DistributeSellerBuyer
	_=context.ShouldBindBodyWith(&distribution,binding.JSON)
	distributions:=models.BuyerFindDistribution(distribution.BuyerId,distribution.Page,distribution.PageSize)
	total:=models.BuyerFindDistributionNum(distribution.BuyerId)
	if len(distributions)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"没有配送中的订单",
		})
	}else{
		data:=make(map[string]interface{})
		data["total"]=total
		data["distributions"]=distributions
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}
//买家查看已完成的配送
func BuyerFindDistributionCom(context *gin.Context){
	var distribution DistributeSellerBuyer
	_=context.ShouldBindBodyWith(&distribution,binding.JSON)
	distributions:=models.BuyerFindDistributionCom(distribution.BuyerId,distribution.Page,distribution.PageSize)
	total:=models.BuyerFindDistributionComNum(distribution.BuyerId)
	if len(distributions)==0{
		code:=e.DISPLAY
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":"没有配送中的订单",
		})
	}else{
		data:=make(map[string]interface{})
		data["total"]=total
		data["distributions"]=distributions
		code:=e.SUCCESS
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}