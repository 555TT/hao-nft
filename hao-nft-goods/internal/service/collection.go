package service

import (
	"hao-nft-goods/internal/dto"
	"hao-nft-goods/internal/infra/database"
	"hao-nft-goods/internal/model"
)

func GetCollectionList(req dto.CollectionListReq) ([]model.Collection, error) {
	var collections []model.Collection
	err := database.DB.Model(&model.Collection{}).
		Where("state = ?", req.State).
		Where("name like ?", "%"+req.Keyword+"%").
		Limit(req.PageSize).
		Offset((req.PageNum - 1) * req.PageSize).
		Find(&collections).Error
	if err != nil {
		return nil, err
	}
	return collections, err
}
