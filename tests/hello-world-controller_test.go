package tests

import (
	"encoding/json"
	helloWorldController "go-api-store-boilerplate/src/application/api/controllers/hello-world"
	presenterJsonModels "go-api-store-boilerplate/src/application/api/presenters/models"
	"go-api-store-boilerplate/src/core/entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorldController(t *testing.T) {
	_assert := assert.New(t)

	helloWorldController := &helloWorldController.HelloWorldController{
		DB: db,
	}

	t.Run("Should receive 200 with empty users from GET /hello-world", func(t *testing.T) {
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.GET("/hello-world", helloWorldController.List)

		expectedUsers := []interface{}([]interface{}{})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/hello-world", nil)
		router.ServeHTTP(w, req)
		var response presenterJsonModels.JsonModel
		json.NewDecoder(w.Body).Decode(&response)

		_assert.Equal(expectedUsers, response.Payload)
	})

	t.Run("Should receive 200 with one user from GET /hello-world", func(t *testing.T) {
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.GET("/hello-world", helloWorldController.List)

		createdUser := &entities.HelloWorldEntity{
			Name: "Test 01",
			BaseEntity: entities.BaseEntity{
				UUID: "40318fcf-f7d8-455d-9a1d-6f01f4162a16",
			},
		}
		db.Create(createdUser)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/hello-world", nil)
		router.ServeHTTP(w, req)
		var response presenterJsonModels.JsonModel
		json.NewDecoder(w.Body).Decode(&response)

		assert.ObjectsAreEqual([]entities.HelloWorldEntity{*createdUser}, response.Payload)
	})

}
