package app

import (
	"hao-nft-goods/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err)
	}
	collectionRouter := r.Group("/collection")
	collectionRouter.GET("/list", handler.CollectionList)
	return r
}
