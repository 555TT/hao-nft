package main

import (
	"hao-nft-goods/internal/app"
	"hao-nft-goods/internal/infra/database"
	"hao-nft-goods/internal/model"
)

func main() {
	if err := database.NewMySQL("conf/conf.yaml"); err != nil {
		panic(err)
	}
	database.DB.AutoMigrate(&model.Collection{})
	if err := app.NewRouter().Run(":8080"); err != nil {
		panic(err)
	}
}
