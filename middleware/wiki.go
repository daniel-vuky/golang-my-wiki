package middleware

import (
	"github.com/daniel-vuky/golang-my-wiki/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateWikiInput(c *gin.Context) {
	var wikiInput controller.WikiInput
	bindErr := c.ShouldBindJSON(&wikiInput)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": bindErr.Error()})
		return
	}
	if len(wikiInput.Title) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Title is missing"})
		return
	}
	if wikiInput.CategoryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can not insert wiki without category"})
		return
	}
	c.Set("wiki_input", wikiInput)
	c.Next()
}
