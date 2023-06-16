package db_entity

type MallOrderAddress struct {
	OrderId int `json:"orderId"`

	UserName string `json:"userName"`

	UserPhone string `json:"userPhone"`

	ProvinceName string `json:"provinceName"`

	CityName string `json:"cityName"`

	RegionName string `json:"regionName"`

	DetailAddress string `json:"detailAddress"`
}
