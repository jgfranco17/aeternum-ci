package system

import (
	"fmt"
	"net/http"
	"time"

	"api/data"
	"api/env"
	"api/logger"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Aeternum API page!",
	})
}

func ServiceInfoHandler(c *gin.Context) {
	// Send the parsed JSON data as a response
	c.JSON(http.StatusOK, data.AboutInfo{
		Name:        "Aeternum API",
		Author:      "Joaquin Franco",
		Repository:  "https://api",
		Environment: env.GetApplicationEnv(),
		License:     "MIT",
		Languages:   []string{"Go"},
	})
}

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, data.HealthStatus{
		Timestamp: time.Now().Format("Mon Jan 2 15:04:05 MST 2006"),
		Status:    "healthy",
	})
}

func NotFoundHandler(c *gin.Context) {
	log := logger.FromContext(c)
	log.Errorf("Non-existent endpoint accessed: %s", c.Request.URL.Path)
	c.JSON(http.StatusNotFound, newMissingEndpoint(c.Request.URL.Path))
}

func newMissingEndpoint(endpoint string) data.BasicErrorInfo {
	return data.BasicErrorInfo{
		StatusCode: http.StatusNotFound,
		Message:    fmt.Sprintf("Endpoint '%s' does not exist", endpoint),
	}
}
