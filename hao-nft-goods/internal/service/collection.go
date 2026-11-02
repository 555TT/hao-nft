package service

import (
	"hao-nft-goods/internal/dto"
	"hao-nft-goods/internal/infra/database"
	"hao-nft-goods/internal/model"
)

func GetCollectionList(req dto.CollectionListReq) (resp dto.CollectionListResp, err error) {
	var collections []model.Collection
	resp.PageSize = req.PageSize
	resp.PageNum = req.PageNum
	db := database.DB.Model(&model.Collection{}).Where("state = ?", req.State)
	if req.Keyword != "" {
		db = db.Where("name like ?", "%"+req.Keyword+"%")
	}
	err = db.Order("updated_at desc").
		Limit(req.PageSize).
		Offset((req.PageNum - 1) * req.PageSize).
		Find(&collections).Error
	if err != nil {
		return resp, err
	}
	resp.Total = int64(len(collections))
	resp.Collections = collections
	return resp, nil
}
