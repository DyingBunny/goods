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
		api.POST("seller_profile",v1.SellerProfile)
		//增加、修改地址
		api.POST("/modify_address",v1.ModifyAddress)
		//卖家发布商品
		api.POST("/release_commodities",v1.ReleaseCommodities)
		//查看商品信息
		api.POST("commodities_profile",v1.CommoditiesProfile)
	}
	return r
}
