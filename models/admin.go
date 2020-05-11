package models

import (
	"goods/models/table"
)
//管理员查看所有上架商品
func AdminFindAllComUp(page int,pageSize int)[]table.Goods{
	var goods []table.Goods
	Db.Where("state=?",1).Limit(pageSize).Offset((page-1)*pageSize).Find(&goods)
	return goods
}

//管理员查看所有下架商品
func AdminFindAllComLow(page int,pageSize int)[]table.Goods{
	var goods []table.Goods
	Db.Where("state=?",0).Limit(pageSize).Offset((page-1)*pageSize).Find(&goods)
	return goods
}
//查看所有上架商品的数量
func AdminFindAllComNumUp() int {
	var total=0
	Db.Model(&table.Goods{}).Where("state=?",1).Count(&total)
	return total
}

//查看所有下架商品的数量
func AdminFindAllComNumLow() int {
	var total=0
	Db.Model(&table.Goods{}).Where("state=?",0).Count(&total)
	return total
}

//管理员查看所有待付款订单
func AdminFindBAllOrderPay(page int,pageSize int)[]table.Order{
	var orders []table.Order
	Db.Where("state=?",1).Limit(pageSize).Offset((page-1)*pageSize).Find(&orders)
	return orders
}

//管理员查看所有待付款订单的数量
func AdminFindAllOrderNumPay()int{
	var total=0
	Db.Model(&table.Order{}).Where("state=?",1).Count(&total)
	return total
}

//管理员查看所有配送中订单
func AdminFindBAllOrderDel(page int,pageSize int)[]table.Order{
	var orders []table.Order
	Db.Where("state=?",2).Limit(pageSize).Offset((page-1)*pageSize).Find(&orders)
	return orders
}

//管理员查看所有配送中订单的数量
func AdminFindAllOrderNumDel()int{
	var total=0
	Db.Model(&table.Order{}).Where("state=?",2).Count(&total)
	return total
}

//管理员查看所有已完成订单
func AdminFindBAllOrderCom(page int,pageSize int)[]table.Order{
	var orders []table.Order
	Db.Where("state=?",3).Limit(pageSize).Offset((page-1)*pageSize).Find(&orders)
	return orders
}

//管理员查看所有已完成订单的数量
func AdminFindAllOrderNumCom()int{
	var total=0
	Db.Model(&table.Order{}).Where("state=?",3).Count(&total)
	return total
}

//管理员改变订单的状态
func AdminChangeOrderState(orderId uint,state uint){
	var order table.Order
	Db.Where("order_id=?",orderId).First(&order).Update("state",state)
}

//管理员查看所有配送中的配送订单
func AdminFindAllDistributionDel(page int,pageSize int)[]table.Distribution{
	var distribution []table.Distribution
	Db.Where("state=?",1).Limit(pageSize).Offset((page-1)*pageSize).Find(&distribution)
	return distribution
}

//管理员查看所有配送中的配送订单的数量
func AdminFindAllDistributionNumDel()int{
	var total=0
	Db.Model(&table.Distribution{}).Where("state=?",1).Count(&total)
	return total
}

//管理员查看所有已完成的配送订单
func AdminFindAllDistributionCom(page int,pageSize int)[]table.Distribution{
	var distribution []table.Distribution
	Db.Where("state=?",0).Limit(pageSize).Offset((page-1)*pageSize).Find(&distribution)
	return distribution
}

//管理员查看所有已完成的配送订单的数量
func AdminFindAllDistributionNumCom()int{
	var total=0
	Db.Model(&table.Distribution{}).Where("state=?",0).Count(&total)
	return total
}

//管理员改变配送订单的状态
func AdminChangeDistributionState(distributionId uint,state uint){
	var distribution table.Distribution
	Db.Where("distribution_id=?",distributionId).First(&distribution).Update("state",state)
}

//管理员查看所有用户
func AdminFindAllUser(page int,pageSize int)[]table.Login{
	var users []table.Login
	Db.Limit(pageSize).Offset((page-1)*pageSize).Find(&users)
	return users
}
//管理员查看所有用户的数量
func AdminFindAllUserNum() int {
	var total=0
	Db.Model(&table.Login{}).Count(&total)
	return total
}