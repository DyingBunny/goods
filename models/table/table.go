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
	Login    	string 	`gorm:"default:'False'"json:"login"`
	LastTime	*time.Time	`json:"last_time"`
	Role     	string 	`json:"role"`
	DeletedAt 	*time.Time `json:"deleted_at"`
}

//买家
type Buyer struct{
	BuyerId 	uint   	`json:"buyer_id"`
	ReceiveAddress		string	`json:"receive_address"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//卖家
type Seller struct{
	SellerId 	uint   	`json:"seller_id"`
	DeliverAddress		string	`json:"deliver_address"`
	Evaluation	uint	`gorm:"default:0"json:"evaluation"`
	Count 		uint	`gorm:"default:0"json:"count"`
	Comprehensive	uint	`gorm:"default:0"json:"comprehensive"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//司机
type Driver struct{
	DriverId 	uint   	`json:"driver_id"`
	Name 		string	`json:"name"`
	Identity    string  `json:"identity"`
	Evaluation 	uint	`gorm:"default:0"json:"evaluation"`
	Count 		uint	`gorm:"default:0"json:"count"`
	AddressId 	uint	`json:"address_id"`
	DeletedAt *time.Time `json:"deleted_at"`
}
////地址
//type Address struct{
//	AddressId	uint	`json:"address_id"`
//	Address		string	`json:"address"`
//	Lng			float64	`json:"lng"`
//	Lat  		float64	`jsn:"lat"`
//	DeletedAt 	*time.Time `json:"deleted_at"`
//}

