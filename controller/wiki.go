package controller

import (
	"github.com/daniel-vuky/golang-my-wiki/model"
	"github.com/daniel-vuky/golang-my-wiki/repository"
	"github.com/gin-gonic/gin"
	ginSession "github.com/go-session/gin-session"
	"net/http"
	"strconv"
	"strings"
)

type Wiki struct {
	Repository *repository.WikiRepository
}

type WikiInput struct {
	CategoryId string `json:"category_id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
}

func getUserId(c *gin.Context) uint64 {
	session := ginSession.FromContext(c)
	userId, _ := session.Get("user_id")
	return userId.(uint64)
}

func getUserName(c *gin.Context) string {
	session := ginSession.FromContext(c)
	username, _ := session.Get("user_name")
	return username.(string)
}

func (wiki *Wiki) GetListWiki(c *gin.Context) {
	userId := getUserId(c)
	categoryId, convertErr := strconv.ParseUint(c.Query("category_id"), 10, 64)
	if convertErr != nil || categoryId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	listWiki, listWikiErr := wiki.Repository.GetListWiki(userId, categoryId)
	if listWikiErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, listWiki)
	return
}

func (wiki *Wiki) CreateWiki(c *gin.Context) {
	input, inputExisted := c.Get("wiki_input")
	if !inputExisted {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing input!"})
		return
	}
	wikiInput, ok := input.(WikiInput)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can not bind the input!"})
		return
	}
	userId := getUserId(c)
	categoryId, categoryConvertErr := strconv.ParseUint(wikiInput.CategoryId, 10, 64)
	if categoryConvertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": categoryConvertErr.Error()})
		return
	}
	// TODO: validate if category belong to this user here
	newWiki := model.Wiki{
		UserId:     userId,
		CategoryId: categoryId,
		Title:      wikiInput.Title,
		Body:       wikiInput.Body,
	}
	insertErr := wiki.Repository.CreateWiki(&newWiki)
	if insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": insertErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"wiki_id": newWiki.WikiId})
	return
}

func (wiki *Wiki) GetWiki(c *gin.Context) {
	wikiId, typeErr := strconv.ParseUint(c.Param("id"), 10, 64)
	if typeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID is not found"})
		return
	}
	userId := getUserId(c)
	collectedWiki, collectedErr := wiki.Repository.GetWikiById(userId, wikiId)
	if collectedErr != nil || collectedWiki.WikiId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wiki not found"})
		return
	}
	c.JSON(http.StatusOK, collectedWiki)
}
func (wiki *Wiki) UpdateWiki(c *gin.Context) {
	input, inputExisted := c.Get("wiki_input")
	if !inputExisted {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing input!"})
		return
	}
	wikiInput, ok := input.(WikiInput)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can not bind the input!"})
		return
	}
	wikiId, typeErr := strconv.ParseUint(c.Param("id"), 10, 64)
	if typeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID is not found"})
		return
	}
	userId := getUserId(c)
	collectedWiki, collectedErr := wiki.Repository.GetWikiById(userId, wikiId)
	if collectedErr != nil || collectedWiki.WikiId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wiki not found"})
		return
	}
	collectedWiki.Title = wikiInput.Title
	collectedWiki.Body = wikiInput.Body
	updatedErr := wiki.Repository.UpdateWiki(&collectedWiki)
	if updatedErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": updatedErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated!"})
}

func (wiki *Wiki) DeleteWiki(c *gin.Context) {
	wikiId, typeErr := strconv.ParseUint(c.Param("id"), 10, 64)
	if typeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID is not found"})
		return
	}
	userId := getUserId(c)
	collectedWiki, collectedErr := wiki.Repository.GetWikiById(userId, wikiId)
	if collectedErr != nil || collectedWiki.WikiId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wiki not found"})
		return
	}
	deletedErr := wiki.Repository.DeleteWiki(&collectedWiki)
	if deletedErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": deletedErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted!"})
}

func (wiki *Wiki) LoadWikiTemplate(c *gin.Context) {
	path := c.Request.URL.Path
	wikiId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if strings.HasPrefix(path, "/wiki/add") {
		c.HTML(
			http.StatusOK,
			"wiki-form.html",
			gin.H{"username": getUserName(c), "category_id": c.Param("category_id")},
		)
		return
	}
	userId := getUserId(c)
	existedWiki, queryErr := wiki.Repository.GetWikiById(userId, wikiId)
	if queryErr != nil || existedWiki.WikiId == 0 {
		c.Redirect(http.StatusFound, "/")
		return
	}
	defaultTemplate := "wiki-form.html"
	if strings.HasPrefix(path, "/wiki/view") {
		defaultTemplate = "wiki-detail.html"
	}
	c.HTML(http.StatusOK, defaultTemplate, gin.H{"username": getUserName(c), "wiki": existedWiki})
}
