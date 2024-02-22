package controller

import (
	"fmt"
	"github.com/daniel-vuky/golang-my-wiki/auth"
	"github.com/daniel-vuky/golang-my-wiki/model"
	"github.com/daniel-vuky/golang-my-wiki/repository"
	"github.com/gin-gonic/gin"
	ginSession "github.com/go-session/gin-session"
	"net/http"
)

type User struct {
	Repository *repository.UserRepository
}

func (user *User) Register(c *gin.Context) {
	hashedPassword, hashedPasswordErr := auth.Hash(c.PostForm("password"))
	if hashedPasswordErr != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": hashedPasswordErr.Error()},
		)
	}
	newUser := model.User{
		UserName: c.PostForm("username"),
		Email:    c.PostForm("email"),
		Password: string(hashedPassword),
	}
	createdUserErr := user.Repository.CreateUser(&newUser)
	if createdUserErr != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": fmt.Sprintf(
					"Can not create new user, %s",
					createdUserErr.Error(),
				),
			},
		)
		return
	}
	// Generate token and move end user to homepage
	createSessionErr := user.CreateSession(c, newUser.UserName, newUser.UserId)
	if createSessionErr != nil {
		c.JSON(
			http.StatusForbidden,
			gin.H{"message": "Can not login, please try again!"},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (user *User) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	inputUser := model.User{
		Email: email,
	}
	getUserErr := user.Repository.GetUser(&inputUser)
	if getUserErr != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Can not find any user with this email"})
		return
	}
	userPassword := inputUser.Password
	if comparePasswordErr := auth.CompareHashAndPassword(userPassword, password); comparePasswordErr != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Email or password incorrect"})
		return
	}
	createSessionError := user.CreateSession(c, inputUser.UserName, inputUser.UserId)
	if createSessionError != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Unknown issue, please try again!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
	return
}

func (user *User) Logout(c *gin.Context) {
	session := ginSession.FromContext(c)
	session.Delete("token")
	session.Delete("user_id")
	session.Delete("user_name")
	session.Save()
	c.Redirect(http.StatusFound, "/login")
	return
}

func (user *User) CheckUserExisted(c *gin.Context) {
	email := c.PostForm("email")
	inputUser := model.User{Email: email}
	_ = user.Repository.GetUser(&inputUser)
	if inputUser.UserId > 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "This email has been registered!"},
		)
	}
	c.Next()
}

func (user *User) IsUserExisted(c *gin.Context) {
	email := c.Query("email")
	inputUser := model.User{Email: email}
	_ = user.Repository.GetUser(&inputUser)
	if inputUser.UserId > 0 {
		c.JSON(
			http.StatusOK,
			gin.H{"existed": true},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"existed": false},
	)
}

func (user *User) CreateSession(c *gin.Context, username string, userId uint64) error {
	token, tokenErr := auth.CreateToken(username)
	if tokenErr != nil {
		return tokenErr
	}
	session := ginSession.FromContext(c)
	session.Set("token", token)
	session.Set("user_id", userId)
	session.Set("user_name", username)
	return session.Save()
}
