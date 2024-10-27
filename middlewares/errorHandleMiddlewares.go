package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yosefalemu/GoCrudChallange/utils"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			utils.ResponseWithError(c, http.StatusInternalServerError, "An internal server error occurred", c.Errors.Last().Err)
		}
	}
}
