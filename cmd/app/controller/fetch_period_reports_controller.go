package controller

import (
	"cmd/app/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FetchPeriodReportsReq struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Count     int    `json:"count"`
}

type FetchEachDayReport struct {
	Id    string `json:"mentalPointId"`
	Date  string `json:"createDate"`
	Point int    `json:"point"`
}

type FetchPeriodReportsRes []FetchEachDayReport

func FetchPeriodReports(c *gin.Context) {
	userId := c.Request.Header.Get("UserId")

	req := new(FetchPeriodReportsReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	parsedEndTime, _ := time.Parse("2006-01-02", req.EndDate)
	parsedEndTimeNext := parsedEndTime.AddDate(0, 0, 1)
	endDateNext := parsedEndTimeNext.Format("2006-01-02")

	targetReports := make(model.MentalPoints, 0)
	res := make(FetchPeriodReportsRes, 0)

	if (req.StartDate != "") && (req.EndDate != "") {
		if req.Count == 0 {
			if err := targetReports.GetReportsByDate(userId, req.StartDate, endDateNext).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			if err := targetReports.GetReportsByDateAndCount(userId, req.StartDate, endDateNext, req.Count).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	} else if (req.StartDate == "") && (req.EndDate == "") && (req.Count != 0) {
		if err := targetReports.GetReportsByCount(userId, req.Count).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	for i := 0; i < len(targetReports); i++ {
		res = append(res, FetchEachDayReport{
			targetReports[i].Id,
			targetReports[i].CreatedDate,
			targetReports[i].Point,
		})
	}

	c.JSON(http.StatusOK, res)
}
