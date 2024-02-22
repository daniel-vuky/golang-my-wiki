package controller

import (
	"github.com/daniel-vuky/golang-my-wiki/model"
	"github.com/daniel-vuky/golang-my-wiki/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	Repository *repository.CategoryRepository
}

type CategoryInput struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
}

func (category *Category) GetListCategory(c *gin.Context) {
	userId := getUserId(c)
	listCategory, listCategoryErr := category.Repository.GetListCategory(userId)
	if listCategoryErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": listCategoryErr.Error()})
		return
	}
	c.JSON(http.StatusOK, listCategory)
	return
}

func (category *Category) CreateCategory(c *gin.Context) {
	input, inputExisted := c.Get("category_input")
	if !inputExisted {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing input"})
		return
	}
	categoryInput, ok := input.(CategoryInput)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing input"})
		return
	}
	userId := getUserId(c)
	newCategory := model.Category{
		Name:             categoryInput.Name,
		UserId:           userId,
		ShortDescription: categoryInput.ShortDescription,
	}
	createdErr := category.Repository.CreateCategory(&newCategory)
	if createdErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": createdErr.Error()})
		return
	}
	c.JSON(http.StatusOK, newCategory)
	return
}

func (category *Category) GetCategory(c *gin.Context) {
	categoryId, convertErr := strconv.ParseUint(c.Param("id"), 10, 64)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": convertErr.Error()})
		return
	}
	userId := getUserId(c)
	collectedCategory, collectedErr := category.Repository.GetCategoryById(userId, categoryId)
	if collectedErr != nil || collectedCategory.CategoryId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"category": collectedCategory})
	return
}

func (category *Category) UpdateCategory(c *gin.Context) {
	input, inputExisted := c.Get("category_input")
	if !inputExisted {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing input"})
		return
	}
	categoryInput, ok := input.(CategoryInput)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing input"})
		return
	}
	categoryId, convertErr := strconv.ParseUint(c.Param("id"), 10, 64)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": convertErr.Error()})
		return
	}
	userId := getUserId(c)
	collectedCategory, collectedErr := category.Repository.GetCategoryById(userId, categoryId)
	if collectedErr != nil || collectedCategory.CategoryId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "category not found"})
		return
	}
	collectedCategory.Name = categoryInput.Name
	collectedCategory.ShortDescription = categoryInput.ShortDescription
	updatedErr := category.Repository.UpdateCategory(&collectedCategory)
	if updatedErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": updatedErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
	return
}

func (category *Category) DeleteCategory(c *gin.Context) {
	categoryId, convertErr := strconv.ParseUint(c.Param("id"), 10, 64)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": convertErr.Error()})
		return
	}
	userId := getUserId(c)
	collectedCategory, collectedErr := category.Repository.GetCategoryById(userId, categoryId)
	if collectedErr != nil || collectedCategory.CategoryId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "category not found"})
		return
	}
	deletedErr := category.Repository.DeleteCategory(&collectedCategory)
	if deletedErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": deletedErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted!"})
}

func (category *Category) LoadCategoryTemplate(c *gin.Context) {
	path := c.Request.URL.Path
	categoryId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if strings.HasPrefix(path, "/category/add") {
		c.HTML(http.StatusOK, "category-form.html", gin.H{"username": getUserName(c)})
		return
	}
	userId := getUserId(c)
	existedCategory, queryErr := category.Repository.GetCategoryById(userId, categoryId)
	if queryErr != nil || existedCategory.CategoryId == 0 {
		c.Redirect(http.StatusFound, "/")
		return
	}
	defaultTemplate := "category-form.html"
	if strings.HasPrefix(path, "/category/view/") {
		defaultTemplate = "category-detail.html"
	}
	c.HTML(http.StatusOK, defaultTemplate, gin.H{"username": getUserName(c), "category": existedCategory})
}
