package routers

import (
	_ "forum/docs"
	"forum/global"
	"forum/internal/middleware"
	v1 "forum/internal/routers/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"token", "Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	config.AllowMethods = []string{"GET", "PUT", "POST", "DELETE"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.Use(static.Serve("/", static.LocalFile(global.AppSetting.StaticPagePath+"/", false)))
	r.StaticFS("/upload", http.Dir(global.AppSetting.UploadSavePath))

	// swagger ui
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("api/v1")
	{
		apiV1.GET("/forum/info", v1.GetForumInfo)
		apiV1.POST("/upload", middleware.JWT(true), v1.UploadImage)

		apiV1.POST("/accounts/login", v1.Login)
		apiV1.POST("/accounts/register", v1.Register)
		apiV1.POST("/accounts/reset-password", middleware.JWT(true), v1.ResetPassword)
		apiV1.DELETE("/accounts/:id", middleware.JWT(true), v1.DeleteAccount)
		apiV1.GET("/accounts/:id", v1.GetAccountInfo)
		apiV1.PUT("/accounts", middleware.JWT(true), v1.EditAccountInfo)

		apiV1.GET("/me/account", middleware.JWT(true), v1.GetMyAccountInfo)
		apiV1.GET("/me/comments", middleware.JWT(true), v1.GetMyComments)

		post := v1.NewPost()
		apiV1.GET("/posts", middleware.JWT(false), post.List)
		apiV1.GET("/posts/:id", post.Get)
		apiV1.POST("/posts", middleware.JWT(true), post.Create)
		apiV1.DELETE("/posts/:id", middleware.JWT(true), post.Delete)
		apiV1.GET("/posts/:id/pin", middleware.JWT(true), post.Pin)

		comment := v1.NewComment()
		apiV1.GET("/comments/:post_id", middleware.JWT(false), comment.List)
		apiV1.POST("/comments", middleware.JWT(true), comment.Create)
		apiV1.PUT("/comments", middleware.JWT(true), comment.Edit)
		apiV1.DELETE("/comments/:post_id/:id", middleware.JWT(true), comment.Delete)
		apiV1.GET("/comments/:post_id/:id", middleware.JWT(false), comment.Get)
		apiV1.GET("/comments/:post_id/:id/vote/:support", middleware.JWT(true), comment.Vote)

		apiV1.GET("/checkin", middleware.JWT(true), v1.CheckIn)
		apiV1.GET("/checkin/records", middleware.JWT(true), v1.GetCheckInRecords)

		apiV1.GET("/top", v1.GetTop)

		apiV1.GET("/messages", middleware.JWT(true), v1.GetMessageList)
		apiV1.GET("/messages/unread", middleware.JWT(true), v1.GetUnreadMessageCount)
		apiV1.GET("/messages/:post_id/:comment_id", middleware.JWT(true), v1.ReadMessage)
	}

	return r
}
