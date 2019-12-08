package v1

import (
	"github.com/gin-gonic/gin"
	"goods/models"
	"goods/pkg/e"
	"net/http"
)

// @Summary 获取用户发货位置
// @Produce  json
// @Param name query string true "username"
// @Success 200 {string} json "{"code" : code,	"msg" : e.GetMsg(code),	"data":{"address":address.Address,"lng":address.Lng,"lat":address.Lat}}"
// @Failure 400 {string} json "{"code" : code,	"msg" : e.GetMsg(code),}"
// @Router /deliver_address [get]
func DeliverAddress(context *gin.Context){
	UserId:=context.Query("user_id")
	address,isExist := models.FindDelAddress(UserId)
	if isExist==false{
		code:=e.ERROR
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
		})
	}else{
		code:=e.SUCCESS
		data := make(map[string]interface{})
		data["address"]=address.Address
		data["lng"]=address.Lng
		data["lat"]=address.Lat
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"data":data,
			//"data":auth.UserId,auth.Username,
		})
	}
}
// @Summary 获取用户收货位置
// @Produce  json
// @Param name query string true "username"
// @Success 200 {string} json "{"code" : code,	"msg" : e.GetMsg(code),	"data":{"address":address.Address,"lng":address.Lng,"lat":address.Lat}}"
// @Failure 400 {string} json "{"code" : code,	"msg" : e.GetMsg(code),}"
// @Router /receive_address [get]
func ReceiveAddress(context *gin.Context){
	UserId:=context.Query("user_id")
	address,isExist := models.FindRecAddress(UserId)
	if isExist==false{
		code:=e.ERROR
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
		})
	}else{
		code:=e.SUCCESS
		data := make(map[string]interface{})
		data["address"]=address.Address
		data["lng"]=address.Lng
		data["lat"]=address.Lat
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"data":data,
			//"data":auth.UserId,auth.Username,
		})
	}
}