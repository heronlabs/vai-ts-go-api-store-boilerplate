package api

import (
	"fmt"
	healthCheckController "go-api-store-boilerplate/src/application/api/controllers/health-check"
	helloWorldController "go-api-store-boilerplate/src/application/api/controllers/hello-world"
	"go-api-store-boilerplate/src/application/api/presenters"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Api(ginApp *gin.Engine, db *gorm.DB) {
	jsonPresenter := &presenters.JsonPresenter{}

	healthCheckController.Handler(ginApp, db, jsonPresenter)
	helloWorldController.Handler(ginApp, db, jsonPresenter)

	fmt.Println("API OK!")
}
