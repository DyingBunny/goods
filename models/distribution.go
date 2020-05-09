package models

import "goods/models/table"

//增加配送订单
func AddDistribution(orderId uint,driverId uint,num uint){
	var order table.Order
	var driver table.Driver
	var user table.Login
	Db.Where("order_id=?",orderId).First(&order)
	Db.Where("driver_id=?",driverId).First(&driver)
	Db.Where("user_id=?",driverId).First(&user)
	Db.Create(&table.Distribution{
		BuyerId:order.BuyerId,
		SellerId:order.SellerId,
		DriverId:driverId,
		OrderId:order.OrderId,
		DeliverAddress:order.DeliverAddress,
		ReceiveAddress:order.ReceiveAddress,
		Num:num,
		TransPrice:order.TransPrice*num,
		Name:driver.Name,
		PhoneNumber:user.PhoneNumber,
		State:1})
}

//订单减去数量
func AdjustOrderNum(orderId uint,num uint){
	var order table.Order
	var tmp table.Order
	Db.Where("order_id=?",orderId).First(&tmp)
	result:=tmp.RemainNum-num
	Db.Model(&order).Where("order_id=?",orderId).Update("remain_num",result)
}

//司机查看自己配送的订单
func DriverFindDistribution(driverId uint)[]table.Distribution{
	var distribution []table.Distribution
	Db.Where("driver_id=? AND state=?",driverId,1).Find(&distribution)
	return distribution
}

//买家卖家按照订单查看配送订单的数量
func SellerBuyerFindDistributionNum(orderId uint)int{
	var total=0
	Db.Model(&table.Distribution{}).Where("order_id=? AND state=?",orderId,1).Count(&total)
	return total
}
//买家卖家按照订单查看配送订单
func SellerBuyerFindDistribution(orderId uint,page int,pageSize int)[]table.Distribution{
	var distribution []table.Distribution
	Db.Where("order_id=? AND state=?",orderId,1).Limit(pageSize).Offset((page-1)*pageSize).Find(&distribution)
	return distribution
}
//买家确认配送订单收货
func ConfirmDistribution(distributionId uint){
	var distribution table.Distribution
	Db.Model(&distribution).Where("distribution_id=?",distributionId).Update("state",0)
}
//查看买家配送订单是否全部确认
func ConfirmAllDistribution(orderId uint)int{
	var total=0
	Db.Model(&table.Distribution{}).Where("order_id=? AND state=?",orderId,1).Count(&total)
	if total==0{
		return 1
	}else{
		return 0
	}
}

//司机接单数量查询
func DriverFindDistributionNum(driverId uint)int{
	var total=0
	Db.Model(&table.Distribution{}).Where("driver_id=? AND state=?",driverId,1).Count(&total)
	return total
}