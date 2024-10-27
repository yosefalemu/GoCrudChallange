package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ResponseWithError(ctx *gin.Context, status int, message string, err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
	ctx.JSON(status, gin.H{"error": message + " OR an internal server error occurred please try again later"})
}
