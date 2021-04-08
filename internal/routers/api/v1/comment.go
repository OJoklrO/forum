package v1

import "github.com/gin-gonic/gin"

type Comment struct {

}

func NewComment() Comment {
	return Comment{}
}

// @Summary get comment by id
// @Produce json
// @Param id path int true "comment id"
// @Success 200 {object} model.Comment "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /api/v1/comments/{id} [get]
func (c Comment) Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"id": ctx.Param("id"),
		"msg": "success",
	})
}

// @Summary get comment list with post id
// @Produce json
// @Param id query int true "post id"
// @Param page query int false "page number"
// @Param page_size query int false "page size"
// @Success 200 {object} model.Comment "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /api/v1/comments [get]
func (c Comment) List(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"id": ctx.Query("id"),
		"msg": "success",
	})
}

// @Summary create comment
// @Produce json
// @Param content body string true "comment content"
// @Param created_by body int true "creator id"
// @Success 200 {object} model.Comment "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /api/v1/comments/create [post]
func (c Comment) Create(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "success",
	})
}

// @Summary delete comment need adm cookie
// @Produce json
// @Param id body int true "comment id"
// @Success 200 {object} model.Comment "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "server error"
// @Router /api/v1/comments/delete [post]
func (c Comment) Delete(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "success",
	})
}