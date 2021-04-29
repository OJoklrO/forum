package v1

type AccountInfoResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Level       uint32 `json:"level"`
}

// @Summary (Todo) Get account information(name, avatar...).
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} AccountInfoResponse "success"
// @Router /api/v1/accounts/{id} [get]
func GetAccountInfo() {

	// todo: post number of a user
	// todo: reply number of a user
}

// @Summary (Todo) Edit account information.
// @Produce json
// @Param body body AccountInfoResponse true "New account information."
// @Success 200 {object} MessageResponse "success"
// @Router /api/v1/accounts [put]
func EditAccountInfo() {

}
