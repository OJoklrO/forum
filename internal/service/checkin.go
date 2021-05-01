package service

import (
	"forum/global"
	"forum/internal/model"
	"time"
)

func (svc *Service) GetCheckInRecords(userId string) ([]bool, error) {
	check := model.Checkin{ID: userId}
	results, err := check.GetThisMonth(global.DBEngine)
	if err != nil {
		return nil, err
	}

	// how many days in this month?
	// the answer: the 0 day of the last month
	tmpT := time.Date(time.Now().Year(), time.Now().Month()+1, 0, 0, 0, 0, 0, time.UTC)
	checks := make([]bool, tmpT.Day())
	for i := range results {
		if results[i].Month == int(time.Now().Month()) {
			checks[results[i].Day-1] = true
		}
	}
	return checks, nil
}
