package middleware

import (
	"github.com/daniel-vuky/golang-my-wiki/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateCategoryInput(c *gin.Context) {
	var categoryInput controller.CategoryInput
	bindErr := c.ShouldBindJSON(&categoryInput)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": bindErr.Error()})
		return
	}
	if len(categoryInput.Name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Name is missing"})
		return
	}
	c.Set("category_input", categoryInput)
	c.Next()
}
