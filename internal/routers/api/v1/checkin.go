package v1

import "github.com/gin-gonic/gin"

type CheckInResponse struct {
	Todo string `json:"todo"`
}

// @Summary (Todo) Get check in records
// @Produce json
// @Param token header string true "jwt token"
// @Success 200 {object} CheckInResponse "success"
// @Router /api/v1/checkin/records [get]
func GetCheckInRecords(c *gin.Context) {
	// todo: return records of this month

}

// @Summary (Todo) Check in.
// @Produce json
// @Param token header string true "jwt token"
// @Success 200 {object} MessageResponse "success"
// @Router /api/v1/checkin [get]
func CheckIn(c *gin.Context) {

}
