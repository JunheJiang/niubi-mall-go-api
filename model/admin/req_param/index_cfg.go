package req_param

import (
	"niubi-mall/model/admin/db_entity"
	"niubi-mall/model/common/req_param"
)

type MallIndexConfigSearch struct {
	db_entity.MallIndexConfig
	req_param.PageInfo
}

type MallIndexConfigAddParams struct {
	ConfigName  string `json:"configName"`
	ConfigType  int    `json:"configType"`
	GoodsId     string `json:"goodsId"`
	RedirectUrl string `json:"redirectUrl"`
	ConfigRank  string `json:"configRank"`
}

type MallIndexConfigUpdateParams struct {
	ConfigId    int    `json:"configId"`
	ConfigName  string `json:"configName"`
	RedirectUrl string `json:"redirectUrl"`
	ConfigType  int    `json:"configType"`
	GoodsId     int    `json:"goodsId"`
	ConfigRank  string `json:"configRank"`
}
