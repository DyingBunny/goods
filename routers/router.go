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

		//查看买家信息
		api.POST("/profile",v1.Profile)
		//查看卖家信息
		api.POST("/seller_profile",v1.SellerProfile)
		//增加、修改地址
		api.POST("/modify_address",v1.ModifyAddress)

			//买家

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
		//买家订单付款
		api.POST("/buyer_pay",v1.BuyerPay)
		//买家评论商品和卖家
		api.POST("/buyer_evaluate_seller_goods",v1.BuyerEvaluateSellerGoods)
		//查看商品的评价
		api.POST("/all_goods_evaluation",v1.AllGoodsEvaluation)
		//买家查看司机的配送状态
		api.POST("/buyer_find_distribution",v1.BuyerFindDistribution)
		//买家查看已完成的配送
		api.POST("/buyer_find_distribution_com",v1.BuyerFindDistributionCom)
		//买家确认司机配送送达并评论司机
		api.POST("/buyer_confirm_distribution",v1.BuyerConfirmDistribution)

			//卖家

		//卖家发布商品
		api.POST("/release_commodities",v1.ReleaseCommodities)
		//卖家修改价格
		api.POST("/seller_change_price",v1.SellerChangePrice)
		//查看一个卖家目前的商品上架与下架数量
		api.POST("/seller_com_num",v1.SellerComNum)
		//查看卖家发布的所有上架商品
		api.POST("/all_seller_com_pro_up",v1.AllSellerComProUp)
		//查看卖家发布的所有下架商品
		api.POST("/all_seller_com_pro_low",v1.AllSellerComProLow)
		//查看卖家目前的订单数量
		api.POST("/buyer_seller_num",v1.BuyerSellerNum)
		//查看卖家的待买家付款订单
		api.POST("/all_buyer_seller_pay",v1.AllBuyerSellerPay)
		//查看卖家的待配送订单
		api.POST("/all_buyer_seller_deli",v1.AllBuyerSellerDeli)
		//查看卖家的已完成订单
		api.POST("/all_buyer_seller_complete",v1.AllBuyerSellerComplete)
		//卖家查看司机的配送状态
		api.POST("/seller_find_distribution",v1.SellerFindDistribution)
		//卖家查看已完成的配送
		api.POST("/seller_find_distribution_com",v1.SellerFindDistributionCom)

			//司机

		//查看司机信息
		api.POST("/driver_profile",v1.DriverProfile)
		//填写姓名和身份证
		api.POST("/modify_identity",v1.ModifyIdentity)
		//司机查看订单
		api.POST("/driver_find_all_order",v1.DriverFindAllOrder)
		//司机配送
		api.POST("/driver_distribute",v1.DriverDistribute)
		//司机查看自己的配送订单
		api.POST("/driver_find_distribution",v1.DriverFindDistribution)
		//司机更新自己的地址
		api.POST("/driver_upload_address",v1.DriverUploadAddress)

			//管理员
		//管理员下架商品
		api.POST("/admin_low_commodity",v1.AdminLowCommodity)
		//管理员查看所有上架商品
		api.POST("/admin_find_all_goods_up",v1.AdminFindAllGoodsUp)
		//管理员查看所有下架商品
		api.POST("/admin_find_all_goods_low",v1.AdminFindAllGoodsLow)
		//管理员查看待付款订单
		api.POST("/admin_find_all_order_pay",v1.AdminFindAllOrderPay)
		//管理员查看正在配送订单
		api.POST("/admin_find_all_order_del",v1.AdminFindAllOrderDel)
		//管理员查看已完成订单
		api.POST("/admin_find_all_order_com",v1.AdminFindAllOrderCom)
		//管理员改变订单状态
		api.POST("/admin_change_order_state",v1.AdminChangeOrderState)
		//管理员查看所有配送中的配送订单
		api.POST("/admin_find_distribution_del",v1.AdminFindDistributionDel)
		//管理员查看所有已完成的配送订单
		api.POST("/admin_find_distribution_com",v1.AdminFindDistributionCom)
		//管理员改变配送订单状态
		api.POST("/admin_change_distribution_state",v1.AdminChangeDistributionState)
		//用户信息管理
		api.POST("/admin_find_all_user",v1.AdminFindAllUser)
	}
	return r
}
