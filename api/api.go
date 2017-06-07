package api

import (
	"log"

	"github.com/fkdldjs02/household/api/model"
	"github.com/fkdldjs02/household/conf"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// App inject configuration
type App struct {
	Env *viper.Viper
}

var (
	app *App
	db  *model.DBPool
)

// SetRouter set up router group
func SetRouter(router *gin.Engine) {
	auth := router.Group("auth")
	{
		auth.POST("/user", SetUserContext)
		auth.POST("/login", AuthContext)
	}

	category := router.Group("category")
	{
		category.POST("", SetCategoryContext)
		category.PUT("/:id", PutCategoryContext)
		category.GET("", GetCategoryListContext)
		category.GET("/:id", GetCategoryContext)
		category.DELETE("/:id", DeleteCategoryContext)
	}

	household := router.Group("household")
	{
		household.POST("", SetHouseholdConext)
		household.PUT("/:id", PutHouseholdContext)
		household.GET("/list/:name", GetHouseholdListContext)
		household.GET("/row/:id", GetHouseholdContext)
		household.DELETE("/row/:id", DeleteHouseholdContext)
	}
}

// NewApp create app
func NewApp(caseOne *conf.CaseOne) {
	log.Println("init app")
	app = &App{
		Env: caseOne.Env,
	}

	db = &model.DBPool{
		Master: caseOne.DBMaster,
	}

}
