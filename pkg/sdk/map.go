package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const(
	//百度AK
	AppKey string ="5dV2K1AEafWxaVhUyiOCjdWKb2wxCUow"
	//通过ip获取地址
	//https
	reqURLForIP string ="https://api.map.baidu.com/location/ip?ak="
	//坐标类型
	coor string="&coor=bd09ll"
)

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
			Point        struct {
				X string `json:"x"`
				Y string `json:"y"`
			} `json:"point"`
		} `json:"address_detail"`
	} `json:"content"`
	Status  int64  `json:"status"`
	Message string `json:"message"`
}

//通过ip获取地址
func GetAddressByIP(address string) (*StructIPToAddress, error) {
	res := new(StructIPToAddress)
	parameter := fmt.Sprintf("&ip=%s&output=json&pois=0", address)
	reqURL := fmt.Sprintf("%s%s%s", reqURLForIP, AppKey, parameter)
	res2, err := requestBaidu("GetAddressViaIP", reqURL)
	if err != nil {
		return res,err
	}
	if res2.(*StructIPToAddress).Status != 0 {
		message := fmt.Sprintf("百度 API 报错：%s", res2.(*StructIPToAddress).Message)
		return res, errors.New(message)
	}
	res3 := res2.(*StructIPToAddress)
	return res3, nil
}

//构建http请求
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

// getResStruct 处理百度 API 返回数据，映射到结构体中
func getResStruct(reqType string) (interface{}, error) {
	var res interface{}
	if reqType == "GetAddressViaIP" {
		return new(StructIPToAddress), nil
	}
	return res, errors.New("结构体请求错误")
}