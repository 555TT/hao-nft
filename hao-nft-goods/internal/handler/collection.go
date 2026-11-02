package handler

import (
	"hao-nft-goods/internal/dto"
	"hao-nft-goods/internal/pkg/response"
	"hao-nft-goods/internal/service"

	"github.com/gin-gonic/gin"
)

func CollectionList(c *gin.Context) {
	var req dto.CollectionListReq
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, response.CodeInvalidParam, err.Error())
		return
	}
	serviceResp, err := service.GetCollectionList(req)
	if err != nil {
		response.Error(c, response.CodeError, err.Error())
		return
	}
	response.Success(c, serviceResp)
}
