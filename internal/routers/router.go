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
	apiv1 := r.Group("/api/v1")
	{
		p := apiv1.Group("/posts")
		p.GET("/", post.List)
		p.GET("/:id", post.Get)
		p.POST("/create", post.Create)
		p.POST("/delete", post.Delete)

		c := apiv1.Group("/comments")
		c.GET("/", comment.List)
		c.GET("/:id", comment.Get)
		c.POST("/create", comment.Create)
		c.POST("/delete", comment.Delete)
	}


	return r
}