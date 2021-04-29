package v1

import "github.com/gin-gonic/gin"

type UploadImageResponse struct {
	Url string `json:"url"`
}

type UploadImageRequest struct {
	Todo string `json:"todo"`
}

// @Summary (Todo) Upload images.
// @Produce json
// @Param body body UploadImageRequest true "todo..."
// @Success 200 {object} UploadImageResponse "success"
// @Router /api/v1/images [post]
func UploadImage(c *gin.Context) {

}
