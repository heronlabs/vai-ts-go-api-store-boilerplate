package healthCheckController

import (
	models "go-api-store-boilerplate/src/application/api/controllers/health-check/models"
	"go-api-store-boilerplate/src/application/api/presenters"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthCheckController struct {
	JsonPresenter presenters.JsonPresenter
}

func (hc HealthCheckController) GetIndex(ctx *gin.Context) {
	response := hc.JsonPresenter.Envelope("Server running!")

	ctx.JSON(200, response)
}

func (hc HealthCheckController) PostWebHook(ctx *gin.Context) {
	var req models.WebHookModel
	req.Method = "POST"
	ctx.BindJSON(&req.Content)

	response := hc.JsonPresenter.Envelope(req)

	ctx.JSON(200, response)
}

func Handler(
	ginApp *gin.Engine,
	db *gorm.DB,
	jsonPresenter *presenters.JsonPresenter,
) {
	healthCheckController := HealthCheckController{
		JsonPresenter: *jsonPresenter,
	}

	ginApp.GET("/", healthCheckController.GetIndex)
	ginApp.POST("/", healthCheckController.PostWebHook)
}
