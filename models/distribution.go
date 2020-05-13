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

//卖家查看司机的配送状态
func SellerFindDistribution(sellerId uint,page int,pageSize int)[]table.Distribution{
	var distribution []table.Distribution
	Db.Where("seller_id=? AND state=?",sellerId,1).Limit(pageSize).Offset((page-1)*pageSize).Find(&distribution)
	return distribution
}
//卖家查看司机的配送状态的数量
func SellerFindDistributionNum(sellerId uint)int{
	var total=0
	Db.Model(&table.Distribution{}).Where("seller_id=? AND state=?",sellerId,1).Count(&total)
	return total
}

//卖家查看已完成的配送
func SellerFindDistributionCom(sellerId uint,page int,pageSize int)[]table.Distribution{
	var distribution []table.Distribution
	Db.Where("seller_id=? AND state=?",sellerId,0).Limit(pageSize).Offset((page-1)*pageSize).Find(&distribution)
	return distribution
}
//卖家查看已完成的配送的数量
func SellerFindDistributionComNum(sellerId uint)int{
	var total=0
	Db.Model(&table.Distribution{}).Where("seller_id=? AND state=?",sellerId,0).Count(&total)
	return total
}

//买家查看司机的配送状态
func BuyerFindDistribution(buyerId uint,page int,pageSize int)[]table.Distribution{
	var distribution []table.Distribution
	Db.Where("buyer_id=? AND state=?",buyerId,1).Limit(pageSize).Offset((page-1)*pageSize).Find(&distribution)
	return distribution
}
//买家查看司机的配送状态的数量
func BuyerFindDistributionNum(buyerId uint)int{
	var total=0
	Db.Model(&table.Distribution{}).Where("buyer_id=? AND state=?",buyerId,1).Count(&total)
	return total
}

//买家查看已完成的配送
func BuyerFindDistributionCom(buyerId uint,page int,pageSize int)[]table.Distribution{
	var distribution []table.Distribution
	Db.Where("buyer_id=? AND state=?",buyerId,0).Limit(pageSize).Offset((page-1)*pageSize).Find(&distribution)
	return distribution
}
//买家查看已完成的配送的数量
func BuyerFindDistributionComNum(buyerId uint)int{
	var total=0
	Db.Model(&table.Distribution{}).Where("buyer_id=? AND state=?",buyerId,0).Count(&total)
	return total
}