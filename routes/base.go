package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes registers all routes
func RegisterRoutes(r *gin.Engine) {
	RegisterPeopleRoutes(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Route Not Found"})
	})
}
