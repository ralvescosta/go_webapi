package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": message})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": message})
}

func ConflictError(c *gin.Context, message string) {
	c.JSON(http.StatusConflict, gin.H{"message": message})
}

func Created(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{})
}
