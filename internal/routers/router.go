package routers

import (
	_ "github.com/OJoklrO/forum/docs"
	"github.com/OJoklrO/forum/internal/middleware"
	v1 "github.com/OJoklrO/forum/internal/routers/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.Translations())

	r.Use(cors.Default())

	// swagger ui
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	post := v1.NewPost()
	comment := v1.NewComment()

	api := r.Group("/api")
	{
		api.GET("/auth", v1.GetAuth)
		api.POST("/auth/operators/create", v1.CreateAuth)
	}

	apiv1 := api.Group("/v1")
	{
		//apiv1.Use(middleware.JWT())

		apiv1.GET("/posts", post.List)
		apiv1.GET("/posts/:id", post.Get)
		apiv1.POST("/posts", post.Create)
		apiv1.DELETE("/posts/:id", post.Delete)

		apiv1.GET("/comments/:post_id", comment.List)
		apiv1.POST("/comments", comment.Create)
		apiv1.DELETE("/comments/:id", comment.Delete)
	}

	return r
}