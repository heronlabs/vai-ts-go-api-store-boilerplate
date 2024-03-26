package helloWorldController

import (
	"go-api-store-boilerplate/src/application/api/presenters"
	"go-api-store-boilerplate/src/core/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HelloWorldController struct {
	DB            *gorm.DB
	JsonPresenter presenters.JsonPresenter
}

func (hw HelloWorldController) List(ctx *gin.Context) {
	var users []entities.HelloWorldEntity
	hw.DB.Find(&users)

	response := hw.JsonPresenter.Envelope(users)

	ctx.JSON(200, response)
}

func Handler(
	ginApp *gin.Engine,
	db *gorm.DB,
	jsonPresenter *presenters.JsonPresenter,
) {
	helloWorldController := HelloWorldController{
		DB:            db,
		JsonPresenter: *jsonPresenter,
	}

	ginApp.GET("/hello-world", helloWorldController.List)
}
