package dto

import "hao-nft-goods/internal/model"

type CollectionListResp struct {
	Total       int64              `json:"total"`
	PageSize    int                `json:"pageSize"`
	PageNum     int                `json:"pageNum"`
	Collections []model.Collection `json:"collections"`
}
