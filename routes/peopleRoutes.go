package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yosefalemu/GoCrudChallange/controllers"
	"github.com/yosefalemu/GoCrudChallange/initializers"
)

func RegisterPeopleRoutes(r *gin.Engine) {
	// Create a group of routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/person", func(ctx *gin.Context) {
			controllers.CreatePerson(ctx, initializers.DB)
		})
		v1.GET("/person", func(ctx *gin.Context) {
			controllers.GetPerson(ctx, initializers.DB)
		})
		v1.GET("/person/:id", func(ctx *gin.Context) {
			controllers.GetPersonById(ctx, initializers.DB)
		})
		v1.PUT("/person/:id", func(ctx *gin.Context) {
			controllers.UpdatePersonById(ctx, initializers.DB)
		})
		v1.DELETE("/person/:id", func(ctx *gin.Context) {
			controllers.DeletePersonById(ctx, initializers.DB)
		})
	}
}
