package v1

import (
	"forum/global"
	"forum/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ForumInfo struct {
	UserCount   int    `json:"user_count"`
	PostCount   int    `json:"post_count"`
	HeaderImage string `json:"header_image"`
}

// @Summary Get forum information.
// @Produce json
// @Success 200 {object} ForumInfo "success"
// @Router /api/v1/forum/info [get]
func GetForumInfo(c *gin.Context) {
	// todo: header image
	var info ForumInfo
	global.DBEngine.Model(model.Account{}).Where("is_del = 0").Count(&info.UserCount)
	global.DBEngine.Model(model.Post{}).Where("is_del = 0").Count(&info.PostCount)
	c.JSON(http.StatusOK, info)
}
