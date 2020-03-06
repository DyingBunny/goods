package v1

import (
	"github.com/gin-gonic/gin"
	"goods/pkg/e"
	"goods/pkg/sdk"
	"net/http"
)

// @Summary 获取地址
// @Produce  json
// @Param name query string true "address"
// @Success 200 {string} json "{"code":code,	"msg" : e.GetMsg(code),"data":{"address":res.Address}}"
// @Failure 400 {string} json "{"code":code,	"msg" : e.GetMsg(code)}"
// @Router /get_address [get]

func GetAddress(context *gin.Context){
	address := context.Query("address")
	res,err:= sdk.GetAddressByIP(address)
	data:=make(map[string]interface{})
	if err==nil {
		code:=e.SUCCESS
		data["address"]=res.Address
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"data":data,
		})
	} else{
		code:=e.ERROR
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg" : e.GetMsg(code),
		})
	}
}