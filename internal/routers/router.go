package routers

import (
	_ "github.com/OJoklrO/forum/docs"
	"github.com/OJoklrO/forum/global"
	"github.com/OJoklrO/forum/internal/middleware"
	v1 "github.com/OJoklrO/forum/internal/routers/api/v1"
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
		r.POST("/account/login", v1.Login)
		r.POST("/account/register", v1.Register)
		r.DELETE("/account/delete/:id", v1.DeleteAccount)

		// todo: jwt
		//apiV1.Use(middleware.JWT())
		post := v1.NewPost()
		apiV1.GET("/posts", post.List)
		apiV1.GET("/posts/:id", post.Get)
		apiV1.POST("/posts", post.Create)
		apiV1.DELETE("/posts/:id", post.Delete)

		comment := v1.NewComment()
		apiV1.GET("/comments/:post_id", comment.List)
		apiV1.POST("/comments/:post_id", comment.Create)
		apiV1.DELETE("/comments/:post_id/:id", comment.Delete)
		// todo: edit comment
		//apiV1.PUT("/comments/:post_id/:id", comment.Edit)
		//apiV1.GET("/comments/:post_id/:id", comment.Get)
	}

	return r
}
