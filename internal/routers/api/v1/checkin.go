package v1

import (
	"forum/global"
	"forum/internal/model"
	"forum/internal/service"
	"forum/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Summary Get check in records in this month
// @Produce json
// @Param token header string true "jwt token"
// @Success 200 {object} []bool "success"
// @Router /api/v1/checkin/records [get]
func GetCheckInRecords(c *gin.Context) {
	svc := service.New(c)
	checks, err := svc.GetCheckInRecords(c.Value("user_id").(string))
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.GetCheckInRecords: "+err.Error())
	}

	c.JSON(http.StatusOK, checks)
}

// @Summary Check in.
// @Produce json
// @Param token header string true "jwt token"
// @Success 200 {object} MessageResponse "success"
// @Router /api/v1/checkin [get]
func CheckIn(c *gin.Context) {
	check := model.Checkin{
		ID:    c.Value("user_id").(string),
		Year:  time.Now().Year(),
		Month: int(time.Now().Month()),
		Day:   time.Now().Day(),
	}

	err := check.Set(global.DBEngine)
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			err.Error())
		return
	}

	svc := service.New(c)
	err = svc.UpdateUserLevel()
	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			"svc.UpdateUserLevel: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, MessageResponse{"success."})
}
