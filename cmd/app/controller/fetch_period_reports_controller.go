package controller

import (
	"cmd/app/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FetchEachDayReport struct {
	Id    string `json:"mentalPointId"`
	Date  string `json:"createDate"`
	Point int    `json:"point"`
}

type FetchPeriodReportsRes []FetchEachDayReport

func FetchPeriodReports(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	parsedEndTime, _ := time.Parse("2006-01-02", endDate)
	parsedEndTimeNext := parsedEndTime.AddDate(0, 0, 1)
	endDateNext := parsedEndTimeNext.Format("2006-01-02")
	userId := c.Request.Header.Get("UserId")

	targetReports := make(model.MentalPoints, 0)
	if err := targetReports.GetReportsByDate(userId, startDate, endDateNext).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := make(FetchPeriodReportsRes, 0)
	for i := 0; i < len(targetReports); i++ {
		res = append(res, FetchEachDayReport{
			targetReports[i].Id,
			targetReports[i].CreatedDate,
			targetReports[i].Point,
		})
	}

	c.JSON(http.StatusOK, res)
}
