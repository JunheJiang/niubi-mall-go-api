package req_param

import (
	"niubi-mall/model/admin/db_entity"
	"niubi-mall/model/common/req_param"
)

type MallCarouselSearch struct {
	db_entity.MallCarousel
	req_param.PageInfo
}

type MallCarouselAddParam struct {
	CarouselUrl  string `json:"carouselUrl"`
	RedirectUrl  string `json:"redirectUrl"`
	CarouselRank string `json:"carouselRank"`
}

type MallCarouselUpdateParam struct {
	CarouselId   int    `json:"carouselId"`
	CarouselUrl  string `json:"carouselUrl"`
	RedirectUrl  string `json:"redirectUrl"`
	CarouselRank string `json:"carouselRank" `
}
