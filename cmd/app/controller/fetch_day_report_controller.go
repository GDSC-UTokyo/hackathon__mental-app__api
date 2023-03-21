package controller

import (
	"cmd/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FetchDayReportRes struct {
	Id      string `json:"mentalPointId"`
	Date    string `json:"createDate"`
	Point   int    `json:"point"`
	Reasons string `json:"reasonIdList"`
}

func FetchDayReport(c *gin.Context) {
	mentalPointId := c.Param("mentalPointId")
	userId := c.Request.Header.Get("UserId")

	targetReport := model.MentalPoint{}
	if err := targetReport.GetReportByMentalPointId(mentalPointId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userId != targetReport.UserId {
		c.JSON(http.StatusForbidden, gin.H{"message": "not permitted"})
		return
	}

	targetReasons := model.ReasonsOnMentalPoints{}
	if err := targetReasons.GetReportByMentalPointId(mentalPointId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := FetchDayReportRes{
		Id:      mentalPointId,
		Date:    targetReport.CreatedDate,
		Point:   targetReport.Point,
		Reasons: targetReasons.ReasonId,
	}

	c.JSON(http.StatusOK, res)
}
