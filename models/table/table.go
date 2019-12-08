package table

import (
	"time"
)

//验证登录
type Login struct{
	UserId   	uint	`json:"user_id"`
	Username 	string 	`json:"username"`
	Password 	string 	`json:"password"`
	PhoneNumber string 	`json:"phone_number"`
	Gender   	string 	`json:"gender"`
	Email    	string 	`json:"email"`
	Login    	string 	`gorm:"default:'False'"json:"login"`
	Role     	string 	`json:"role"`
	RoleId 	 	uint   	`json:"role_id"`
	DeletedAt 	*time.Time `json:"deleted_at"`
}
//司机
type Driver struct{
	DriverId 	uint   	`json:"driver_id"`
	Evaluation 	uint	`gorm:"default:0"json:"evaluation"`
	Count 		uint	`gorm:"default:0"json:"count"`
	AddressId 	uint	`json:"address_id"`
	DeletedAt *time.Time `json:"deleted_at"`
}
//收货
type ReceiveMapping struct {
	UserAddressId	uint	`json:"user_address_id"`
	Pid 			uint	`json:"pid"`
	Cid 			uint	`json:"cid"`
	DeletedAt 	*time.Time `json:"deleted_at"`
}
//发货
type DeliverMapping struct {
	UserAddressId	uint	`json:"user_address_id"`
	Pid				uint	`json:"pid"`
	Cid				uint	`json:"cid"`
	DeletedAt 	*time.Time 	`json:"deleted_at"`
}
//地址
type Address struct{
	AddressId	uint	`json:"address_id"`
	Address		string	`json:"address"`
	Lng			float64	`json:"lng"`
	Lat  		float64	`jsn:"lat"`
	DeletedAt 	*time.Time `json:"deleted_at"`
}
//订单
type Order struct{
	OrderId		uint	`json:"order_id"`
	Text		string	`json:"text"`
	Complete	string 	`gorm:"default:'False'"json:"complete"`
	CreateTime	time.Time	`json:"create_time"`
	CompleteTime	time.Time	`json:"complete_time"`
	Price		uint	`json:"price"`
	Distance 	float64	`json:"distance"`
	DeletedAt 	*time.Time `json:"deleted_at"`
}
