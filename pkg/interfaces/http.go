package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": message})
}

func InvalidBody(c *gin.Context, errors []string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong Body", "errors": errors})
}

func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": message})
}

func Forbiden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, gin.H{"message": message})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": message})
}

func Conflict(c *gin.Context, message string) {
	c.JSON(http.StatusConflict, gin.H{"message": message})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"message": message})
}

func Created(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{})
}

func Ok(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, body)
}
