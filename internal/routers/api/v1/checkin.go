package v1

import (
	"forum/global"
	"forum/internal/model"
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
	check := model.Checkin{ID: c.Value("user_id").(string)}
	results, err := check.GetThisMonth(global.DBEngine)

	if err != nil {
		app.ResponseError(c, http.StatusInternalServerError,
			err.Error())
		return
	}

	tmpT := time.Date(time.Now().Year(), time.Now().Month()+1, 0, 0, 0, 0, 0, time.UTC)
	checkResp := make([]bool, tmpT.Day())
	for i := range results {
		if results[i].Month == int(time.Now().Month()) {
			checkResp[results[i].Day-1] = true
		}
	}

	c.JSON(http.StatusOK, checkResp)
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

	c.JSON(http.StatusOK, MessageResponse{"success."})
}
