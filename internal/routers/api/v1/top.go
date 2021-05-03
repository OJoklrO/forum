package v1

import (
	"forum/internal/service"
	"forum/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TopResponse struct {
	IDs    []uint32 `json:"ids"`
	Titles []string `json:"titles"`
}

// @Summary Top posts, hot points.
// @Produce json
// @Success 200 {object} TopResponse "list of top posts"
// @Router /api/v1/top [get]
func GetTop(c *gin.Context) {
	// todo: a more reasonable top news rules("/top")
	svc := service.New(c)
	posts, _, err := svc.GetPostList(1, 5, "", false)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.GetPostList err: "+err.Error())
		return
	}
	response := TopResponse{make([]uint32, len(posts)), make([]string, len(posts))}
	for i := range posts {
		response.Titles[i] = posts[i].Title
		response.IDs[i] = posts[i].ID
	}
	c.JSON(http.StatusOK, response)
}
