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

	r.Use(static.Serve("/", static.LocalFile(global.AppSetting.StaticPagePath + "/", false)))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// swagger ui
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	r.GET("/auth", v1.GetAuth)
	r.POST("/auth/operators/create", v1.CreateAuth)

	apiv1 := r.Group("api/v1")
	{
		// todo: jwt
		//apiv1.Use(middleware.JWT())
		post := v1.NewPost()
		apiv1.GET("/posts", post.List)
		apiv1.GET("/posts/:id", post.Get)
		apiv1.POST("/posts", post.Create)
		apiv1.DELETE("/posts/:id", post.Delete)

		comment := v1.NewComment()
		apiv1.GET("/comments/:post_id", comment.List)
		apiv1.POST("/comments", comment.Create)
		apiv1.DELETE("/comments/:id", comment.Delete)
	}

	return r
}