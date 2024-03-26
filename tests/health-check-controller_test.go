package tests

import (
	"bytes"
	"encoding/json"
	healthCheckController "go-api-store-boilerplate/src/application/api/controllers/health-check"
	healthCheckControllerModels "go-api-store-boilerplate/src/application/api/controllers/health-check/models"
	presenterJsonModels "go-api-store-boilerplate/src/application/api/presenters/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckController(t *testing.T) {
	_assert := assert.New(t)

	healthCheckController := &healthCheckController.HealthCheckController{}

	t.Run("Should receive 200 from GET /", func(t *testing.T) {
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.GET("/", healthCheckController.GetIndex)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)
		var response presenterJsonModels.JsonModel
		json.NewDecoder(w.Body).Decode(&response)

		var expectedResponsePayload = presenterJsonModels.JsonModel{
			Payload: "Server running!",
		}
		_assert.Equal(expectedResponsePayload, response)
	})

	t.Run("Should receive 200 from POST /webhook", func(t *testing.T) {
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.POST("/", healthCheckController.PostWebHook)

		body := map[string]interface{}{
			"testOne": "hello world!",
		}

		jsonBytes, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/", bytes.NewReader(jsonBytes))
		router.ServeHTTP(w, req)
		var response presenterJsonModels.JsonModel
		json.NewDecoder(w.Body).Decode(&response)

		var expectedResponsePayload = presenterJsonModels.JsonModel{
			Payload: healthCheckControllerModels.WebHookModel{
				Method:  "POST",
				Content: body,
			},
		}

		assert.ObjectsAreEqual(expectedResponsePayload, response)
	})

}
