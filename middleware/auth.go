package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func ValidateEmailAndPassword(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	emailRex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if len(email) == 0 || len(password) == 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "missing required fields"},
		)
		return
	}
	if len(password) < 6 || len(password) > 12 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "password need greater than 6 and less than 12 chars"},
		)
		return
	}
	if !emailRex.MatchString(email) {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "wrong email format, please try again"},
		)
		return
	}
	c.Next()
}
