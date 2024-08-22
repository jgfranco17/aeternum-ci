package v0

import (
	errors "api/router/error_handling"

	"github.com/gin-gonic/gin"
)

// Adds v0 routes to the router.
func SetRoutes(route *gin.Engine) {
	v0 := route.Group("/v0")
	{
		ciRoutes := v0.Group("/pipelines")
		{
			ciRoutes.GET("/:id", errors.WithErrorHandling(getPipeline()))
		}
	}
}
