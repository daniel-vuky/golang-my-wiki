package middleware

import (
	"github.com/daniel-vuky/golang-my-wiki/auth"
	"github.com/daniel-vuky/golang-my-wiki/repository"
	"github.com/gin-gonic/gin"
	ginSession "github.com/go-session/gin-session"
	"net/http"
)

type User struct {
	Repository *repository.UserRepository
}

func ValidateToken(c *gin.Context) {
	status := isTokenValid(c)
	if status == http.StatusOK {
		c.Next()
	}
	c.Redirect(http.StatusFound, "/login")
}

func IsLoggedIn(c *gin.Context) {
	status := isTokenValid(c)
	if status == http.StatusOK {
		c.Redirect(http.StatusFound, "/")
	}
	c.Next()
}

func isTokenValid(c *gin.Context) int {
	session := ginSession.FromContext(c)
	if session == nil {
		return http.StatusFound
	}
	token, tokenValid := session.Get("token")
	if token == nil || !tokenValid {
		return http.StatusFound
	}
	parsedToken, parsedTokenErr := auth.ValidateToken(token.(string))
	if parsedTokenErr != nil || !parsedToken.Valid {
		return http.StatusFound
	}
	return http.StatusOK
}

func GetUsernameFromContext(c *gin.Context) string {
	session := ginSession.FromContext(c)
	token, tokenErr := session.Get("token")
	if token == nil || !tokenErr {
		return ""
	}
	username, err := auth.GetUserInformationFromToken(
		token.(string),
		"user_name",
	)
	if err != nil {
		return ""
	}
	return username.(string)
}
