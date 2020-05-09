package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"goods/models/table"
	"io/ioutil"
	"net/http"
)

const (
	//AppKey
	AppKey string = "5dV2K1AEafWxaVhUyiOCjdWKb2wxCUow"
	//通过 IP 获取地址
	reqURLForIP string = "http://api.map.baidu.com/location/ip?ak="
)
//映射结构体
type StructIPToAddress struct {
	Address string `json:"address"`
	Content struct {
		Address       string `json:"address"`
		AddressDetail struct {
			City         string `json:"city"`
			CityCode     int64  `json:"city_code"`
			District     string `json:"district"`
			Province     string `json:"province"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
		} `json:"address_detail"`
		Point        struct {
			X string `json:"x"`
			Y string `json:"y"`
		} `json:"point"`
	} `json:"content"`
	Status  int64  `json:"status"`
	Message string `json:"message"`
}

//处理百度API返回数据，映射到结构体中
func getResStruct(reqType string) (interface{}, error) {
	var res interface{}
	if reqType == "GetAddressViaIP" {
		return new(StructIPToAddress), nil
	}
	return res, errors.New("结构体请求错误")
}
//构建HTTP请求
func requestBaidu(reqType, reqURL string) (interface{}, error) {
	res, err := getResStruct(reqType)
	if err != nil {
		return res, err
	}
	httpClient := http.Client{}
	resp, err := httpClient.Get(reqURL)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		err := json.Unmarshal(bytes, &res)
		if err != nil {
			return res, err
		}
	} else {
		return res, errors.New("请求百度API失败，状态码不等于200")
	}
	return res, nil
}
//通过ip获取地址
func GetAddressViaIP(address string) (*StructIPToAddress, error) {
	res := new(StructIPToAddress)
	parameter := fmt.Sprintf("&ip=%s&&coor=bd09ll", address)
	reqURL := fmt.Sprintf("%s%s%s", reqURLForIP, AppKey, parameter)
	res2, err := requestBaidu("GetAddressViaIP", reqURL)
	if err != nil {
		return res, err
	}
	if res2.(*StructIPToAddress).Status != 0 {
		message := fmt.Sprintf("百度 API 报错：%s", res2.(*StructIPToAddress).Message)
		return res, errors.New(message)
	}
	res3 := res2.(*StructIPToAddress)
	return res3, nil

}
//插入司机信息
func AddressInsertDriver(driverId uint){
	Db.Create(&table.Address{
		DriverId:driverId})
}
//司机接单之后插入发货收货地址
func AddressUpdateDelRec(orderId uint,driverId uint){
	var order table.Order
	var address table.Address
	Db.Where("order_id=?",orderId).First(&order)
	Db.Model(&address).Where("driver_id=?",driverId).Update("deliver_address",order.DeliverAddress)
	Db.Model(&address).Where("driver_id=?",driverId).Update("receive_address",order.ReceiveAddress)
}

//司机上传数据插入地址和经纬度
func AddressUpdateAddressXY(driverId uint,address string,lng string,lat string){
	var driverAddress table.Address
	Db.Model(&driverAddress).Where("driver_id=?",driverId).Update("address",address)
	Db.Model(&driverAddress).Where("driver_id=?",driverId).Update("lng",lng)
	Db.Model(&driverAddress).Where("driver_id=?",driverId).Update("lat",lat)
}

