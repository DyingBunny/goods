package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goods/models"
	"goods/models/table"
	"goods/pkg/e"
	"net/http"
)

//司机上传位置
func DisplayAddress(context *gin.Context){
	var addressDriver table.Address
	_=context.ShouldBindBodyWith(&addressDriver,binding.JSON)
	ip:=context.ClientIP()
	address,_:=models.GetAddressViaIP(ip)
	models.AddressUpdateAddressXY(addressDriver.DriverId,address.Address,address.Content.Point.X,address.Content.Point.Y)
	code:=e.SUCCESS
	data := make(map[string]interface{})
	data["address"]=address
	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data":data,
	})
}