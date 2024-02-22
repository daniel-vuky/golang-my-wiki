package application

import (
	"github.com/daniel-vuky/golang-my-wiki/controller"
	"github.com/daniel-vuky/golang-my-wiki/middleware"
	"github.com/daniel-vuky/golang-my-wiki/repository"
	"github.com/gin-gonic/gin"
	ginSession "github.com/go-session/gin-session"
	"net/http"
)

func (app *App) LoadRoutes() {
	router := gin.Default()
	router.Use(ginSession.New())

	router.Static("/static", "./public/static")
	router.LoadHTMLGlob("./public/templates/*/*")

	loadIndexRoute(app, router)
	loadAuthRoutes(app, router)
	loadWikiRoutes(app, router)
	loadCategoryRoutes(app, router)

	app.router = router
}

func loadIndexRoute(app *App, router *gin.Engine) {
	router.GET("/", middleware.ValidateToken, func(c *gin.Context) {
		username := middleware.GetUsernameFromContext(c)
		c.HTML(http.StatusOK, "index.html", gin.H{"username": username})
	})
}

func loadAuthRoutes(app *App, router *gin.Engine) {
	wikiController := &controller.User{
		Repository: &repository.UserRepository{
			Db: app.rdb,
		},
	}
	router.GET("/login", middleware.IsLoggedIn, func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	router.GET("/register", middleware.IsLoggedIn, func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})
	router.POST(
		"/login",
		middleware.ValidateEmailAndPassword,
		wikiController.Login,
	)
	router.POST(
		"/register",
		middleware.ValidateEmailAndPassword,
		wikiController.CheckUserExisted,
		wikiController.Register,
	)
	router.GET(
		"/logout",
		wikiController.Logout,
	)
	router.GET(
		"/user-existed",
		wikiController.IsUserExisted,
	)
}

func loadWikiRoutes(app *App, router *gin.Engine) {
	wikiController := &controller.Wiki{
		Repository: &repository.WikiRepository{
			Db: app.rdb,
		},
	}
	wikiGroup := router.Group("/wiki")
	wikiGroup.Use(middleware.ValidateToken)
	{
		wikiGroup.GET(
			"/",
			wikiController.GetListWiki,
		)
		wikiGroup.POST(
			"/",
			middleware.ValidateWikiInput,
			wikiController.CreateWiki,
		)
		wikiGroup.GET(
			"/:id",
			wikiController.GetWiki,
		)
		wikiGroup.PUT(
			"/:id",
			middleware.ValidateWikiInput,
			wikiController.UpdateWiki,
		)
		wikiGroup.DELETE(
			"/:id",
			wikiController.DeleteWiki,
		)
		wikiGroup.GET(
			"/add/:category_id",
			wikiController.LoadWikiTemplate,
		)
		wikiGroup.GET(
			"/edit/:id",
			wikiController.LoadWikiTemplate,
		)
		wikiGroup.GET(
			"/view/:id",
			wikiController.LoadWikiTemplate,
		)
	}
}

func loadCategoryRoutes(app *App, router *gin.Engine) {
	categoryController := &controller.Category{
		Repository: &repository.CategoryRepository{
			Db: app.rdb,
		},
	}
	categoryGroup := router.Group("/category")
	categoryGroup.Use(middleware.ValidateToken)
	{
		categoryGroup.GET(
			"/",
			categoryController.GetListCategory,
		)
		categoryGroup.POST(
			"/",
			middleware.ValidateCategoryInput,
			categoryController.CreateCategory,
		)
		categoryGroup.GET(
			"/:id",
			categoryController.GetCategory,
		)
		categoryGroup.PUT(
			"/:id",
			middleware.ValidateCategoryInput,
			categoryController.UpdateCategory,
		)
		categoryGroup.DELETE(
			"/:id",
			categoryController.DeleteCategory,
		)
		categoryGroup.GET(
			"/add/",
			categoryController.LoadCategoryTemplate,
		)
		categoryGroup.GET(
			"/edit/:id",
			categoryController.LoadCategoryTemplate,
		)
		categoryGroup.GET(
			"/view/:id",
			categoryController.LoadCategoryTemplate,
		)
	}
}
