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
	Name	string	`json:"name"`
	ReceiveNumber	string	`json:"receive_number"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//卖家
type Seller struct{
	SellerId 	uint   	`json:"seller_id"`
	DeliverAddress		string	`json:"deliver_address"`
	Name	string	`json:"name"`
	DeliverNumber	string	`json:"deliver_number"`
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
	Comprehensive	uint	`gorm:"default:0"json:"comprehensive"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//商品
type Goods struct{
	GoodsId		uint	`json:"goods_id"`
	Text		string	`json:"text"`
	SellerId	uint	`json:"seller_id"`
	SellerUsername	string	`json:"seller_username"`
	TotalNum	uint	`json:"total_num"`
	RemainNum	uint	`json:"remain_num"`
	Price		uint	`json:"price"`
	TransPrice	uint	`json:"trans_price"`
	DeliverAddress	string	`json:"deliver_address"`
	CreateTime	time.Time	`json:"create_time"`
	Detail 		string	`json:"detail"`
	State 		uint	`json:"state"`
	Identification	string	`json:"identification"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type SellerGoods struct{
	SellerGoodsID	uint	`json:"seller_goods_id"`
	GoodsId 	uint	`json:"goods_id"`
	SellerId	uint	`json:"seller_id"`
	DeletedAt *time.Time `json:"deleted_at"`
}
//订单
type Order struct {
	OrderId		uint	`json:"order_id"`
	GoodsId		uint	`json:"goods_id"`
	SellerId	uint	`json:"seller_id"`
	BuyerId		uint	`json:"buyer_id"`
	Text		string	`json:"text"`
	CreateTime	time.Time	`json:"create_time"`
	TotalNum	uint	`json:"total_num"`
	RemainNum	uint	`json:"remain_num"`
	TotalPrice	uint	`json:"total_price"`
	TotalTransPrice	uint	`json:"total_trans_price"`
	DeliverAddress	string	`json:"deliver_address"`
	ReceiveAddress	string	`json:"receive_address"`
	State		uint		`json:"state"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//评价
type Evaluation struct{
	EvaluationId	uint	`json:"evaluation_id"`
	Score		    uint	`json:"score"`
	GoodsId		uint	`json:"goods_id"`
	SellerId	uint	`json:"seller_id"`
	BuyerId		uint	`json:"buyer_id"`
	BuyerUsername 	string 	`json:"buyer_username"`
	Comment 		string	`json:"detail"`
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

