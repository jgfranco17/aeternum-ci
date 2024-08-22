package v0

import (
	"fmt"
	"net/http"
	"strconv"

	"api/errors"

	"api/models"

	"github.com/gin-gonic/gin"
)

func getPipeline() func(c *gin.Context) error {
	return func(c *gin.Context) error {
		value := c.Param("id")
		number, err := strconv.Atoi(value)
		if err != nil {
			return errors.NewInputError(c, fmt.Sprintf("Failed to parse ID: %v", err))
		}
		pipeline := &models.Pipeline{
			Id:  fmt.Sprintf("sample-id-string-%d", number),
			Url: "www.example.com",
		}
		c.JSON(http.StatusOK, pipeline)
		return nil
	}
}
