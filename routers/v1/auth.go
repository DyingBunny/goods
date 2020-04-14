package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goods/models"
	"goods/models/table"
	"goods/pkg/e"
	logging "goods/pkg/logging/serve"
	"goods/pkg/util"
	"net/http"
	"time"
)

type address struct{
	UserId  string `json:"user_id"`
	Address string `json:"address"`
}

//登录
func Login(context *gin.Context){
	var people table.Login
	if context.BindJSON(&people)==nil{
		if models.Check(people.Username,people.Password){
			token,_ := util.GenerateToken(people.Username, people.Password)
			userid:=models.FindUserId(people.Username,people.Password)
			models.UpdateSta(userid)
			data := make(map[string]interface{})
			role:=models.FindRole(people.Username,people.Password)
			code:=e.SUCCESS
			data["token"]=token
			data["role"]=role
			data["user_id"]=userid
			data["login_time"]=time.Now()
			context.JSON(http.StatusOK,gin.H{
				"code":code,
				"msg" : e.GetMsg(code),
				"data":data,})
		}else{
			code:=e.ERROR
			context.JSON(http.StatusOK,gin.H{
				"code":code,
				"msg" : e.GetMsg(code),
			})
		}
	}
}
//查看买家信息
func Profile(context *gin.Context){
	var address address
	_ = context.ShouldBindBodyWith(&address,binding.JSON)
	auth,isExist := models.Find(address.UserId)
	if isExist==false{
		code:=e.ERROR
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
		})
	}else{
		code:=e.SUCCESS
		var buyer table.Buyer
		models.Db.First(&buyer,"buyer_id=?",address.UserId)
		data := make(map[string]interface{})
		data["user_id"]=auth.UserId
		data["username"]=auth.Username
		data["password"]=auth.Password
		data["phone_number"]=auth.PhoneNumber
		data["gender"]=auth.Gender
		data["login"]=auth.Login
		data["last_time"]=auth.LastTime
		data["role"]=auth.Role
		data["receive_address"]=buyer.ReceiveAddress
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"data":data,
			//"data":auth.UserId,auth.Username,
		})
	}
}
//注册
func Register(context *gin.Context){
	var people table.Login
	if context.BindJSON(&people)==nil{
		if models.Check(people.Username,people.Password)==false{
			models.AddUser(people.Username,people.Password,people.PhoneNumber,people.Gender,people.Role)
			userid:=models.FindUserId(people.Username,people.Password)
			if people.Role=="buyer"{
				models.AddBuyer(userid)
			}else if people.Role=="seller"{
				models.AddSeller(userid)
			}else{
				models.AddDriver(userid)
			}
			code:=e.SUCCESS
			context.JSON(http.StatusOK,gin.H{
				"code":code,
				"msg" : e.GetMsg(code),
			})
		}else{
			code:=e.ERROR
			context.JSON(http.StatusOK,gin.H{
				"code":code,
				"msg" : e.GetMsg(code),
			})
		}
	}
}
//填写/修改收货/发货地址
func ModifyAddress(context *gin.Context){
	var address address
	_ = context.ShouldBindBodyWith(&address,binding.JSON)
	auth,isExist := models.Find(address.UserId)
	if isExist==false{
		code:=e.ERROR
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
		})
	}else{
		code:=e.SUCCESS
		models.ModifyAddress(auth,address.Address)
		context.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			//"data":auth.UserId,auth.Username,
		})
	}
}

//退出登录
func LogOut(context *gin.Context){
	var address address
	_=context.ShouldBindBodyWith(&address,binding.JSON)
	auth,isExist:=models.Find(address.UserId)
	if isExist==false{
		code:=e.ERROR
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
		})
	}else{
		code:=e.SUCCESS
		models.Turnstatus(auth.UserId,auth,"False")
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
		})
	}
}

//检查用户名是否重复
func Duplicate(context *gin.Context){
	var auth table.Login
	var people table.Login
	_=context.BindJSON(&people)
	models.Db.First(&auth,"username=?",people.Username)
	valid:=validation.Validation{}
	valid.Required(auth.Username,people.Username)
	data := make(map[string]interface{})
	if valid.HasErrors(){
		for _,err:=range valid.Errors{
			logging.Info(err.Key,err.Message)
		}
		code:=e.SUCCESS
		data["isExist"]="False"
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}else{
		code:=e.SUCCESS
		data["isExist"]="True"
		context.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":e.GetMsg(code),
			"data":data,
		})
	}
}