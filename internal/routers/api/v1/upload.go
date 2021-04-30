package v1

import (
	"forum/internal/service"
	"forum/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UploadImageResponse struct {
	Url string `json:"url"`
}

type UploadImageRequest struct {
	Todo string `json:"todo"`
}

// @Summary Upload.
// @Produce json
// @Param file formData file true "image file"
// @Param token header string true "jwt token"
// @Success 200 {object} UploadImageResponse "success"
// @Router /api/v1/upload [post]
func UploadImage(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		app.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	if fileHeader == nil {
		app.ResponseError(c, http.StatusBadRequest, "fileHeader is null.")
		return
	}

	svc := service.New(c)
	fileInfo, err := svc.Upload(file, fileHeader)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fileInfo)
}
