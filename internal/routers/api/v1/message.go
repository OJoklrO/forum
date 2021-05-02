package v1

import "github.com/gin-gonic/gin"

// @Summary (Todo) If there is any message.
// @Produce json
// @Param token header string true "jwt token"
// @Success 200 {object} string "todo"
// @Router /api/v1/check/messages [get]
func HasMessage(c *gin.Context) {

}

// @Summary (Todo) Get unread messages(reply notifications).
// @Produce json
// @Param token header string true "jwt token"
// @Success 200 {object} string "todo"
// @Router /api/v1/messages [get]
func GetMessageList(c *gin.Context) {

}

// @Summary (Todo) Set an unread message read.
// @Produce json
// @Param token header string true "jwt token"
// @Router /api/v1/messages/{id} [get]
func ReadMessage(c *gin.Context) {

}
