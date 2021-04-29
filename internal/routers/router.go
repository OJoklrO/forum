package routers

import (
	_ "forum/docs"
	"forum/global"
	"forum/internal/middleware"
	v1 "forum/internal/routers/api/v1"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.Translations())

	// todo: an API to return the user count, post number, big image.

	// todo: cors
	//config := cors.DefaultConfig()
	//config.AllowAllOrigins = true
	//config.AllowHeaders = []string{"token", "Authorization", "Content-Type", "Upgrade", "Origin",
	//	"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	//config.AllowMethods = []string{"GET", "PUT", "POST", "DELETE"}
	//config.AllowCredentials = true
	//r.Use(cors.New(config))

	r.Use(static.Serve("/", static.LocalFile(global.AppSetting.StaticPagePath+"/", false)))

	// swagger ui
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("api/v1")
	{
		apiV1.POST("/accounts/login", v1.Login)
		apiV1.POST("/accounts/register", v1.Register)
		apiV1.DELETE("/accounts/:id", middleware.JWT(), v1.DeleteAccount)

		post := v1.NewPost()
		apiV1.GET("/posts", post.List)
		apiV1.GET("/posts/:id", post.Get)
		apiV1.POST("/posts", middleware.JWT(), post.Create)
		apiV1.DELETE("/posts/:id", middleware.JWT(), post.Delete)

		comment := v1.NewComment()
		apiV1.GET("/comments/:post_id", comment.List)
		apiV1.POST("/comments", middleware.JWT(), comment.Create)
		apiV1.PUT("/comments", middleware.JWT(), comment.Edit)
		apiV1.DELETE("/comments/:post_id/:id", middleware.JWT(), comment.Delete)
		apiV1.GET("/comments/:post_id/:id", comment.Get)
	}

	return r
}
