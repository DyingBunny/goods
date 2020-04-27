package routers

import (
	_ "bytes"
	_ "encoding/json"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	_ "goods/docs"
	"goods/middleware/core"
	"goods/middleware/jwt"
	"goods/models"
	logging "goods/pkg/logging/serve"
	_ "goods/pkg/sdk"
	"goods/pkg/setting"
	v1 "goods/routers/v1"
	_ "io/ioutil"
	_ "net/http/httptest"
)

//初始化路由
func InitRouter() *gin.Engine{
	gin.SetMode(setting.RunMode)
	r:=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(core.Cors())
	logging.Init()
	models.Init()
	//登录
	r.POST("/login",v1.Login)
	//注册
	r.POST("/register",v1.Register)
	//查看用户名是否重复
	r.POST("/duplicate",v1.Duplicate)
	//退出登录
	r.POST("/logout",v1.LogOut)
	api := r.Group("/api/v1")
	api.Use(jwt.JWT())
	{

		//-----auth

		//查看买家信息
		api.POST("/profile",v1.Profile)
		//查看卖家信息
		api.POST("/seller_profile",v1.SellerProfile)
		//增加、修改地址
		api.POST("/modify_address",v1.ModifyAddress)

		//-----goods

		//卖家发布商品
		api.POST("/release_commodities",v1.ReleaseCommodities)
		//查看一个卖家目前的商品上架与下架数量
		api.POST("/seller_com_num",v1.SellerComNum)
		//查看卖家发布的所有上架商品
		api.POST("/all_seller_com_pro_up",v1.AllSellerComProUp)
		//查看卖家发布的所有下架商品
		api.POST("/all_seller_com_pro_low",v1.AllSellerComProLow)
		//按照价格降序查看所有已发布的商品
		api.POST("/view_item_desc",v1.ViewItemDesc)
		//按照价格升序查看所有已发布的商品
		api.POST("/view_item_asc",v1.ViewItemAsc)
		//买家购买商品
		api.POST("/purchase_commodities",v1.PurchaseCommodities)
		//买家查看订单的数量
		api.POST("/buyer_order_num",v1.BuyerOrderNum)
		//买家查看待付款的订单
		api.POST("/all_buyer_order_pay",v1.AllBuyerOrderPay)
		//买家查看配送中的订单
		api.POST("/all_buyer_order_deli",v1.AllBuyerOrderDeli)
		//买家查看已完成的订单
		api.POST("/all_buyer_order_complete",v1.AllBuyerOrderComplete)
	}
	return r
}
