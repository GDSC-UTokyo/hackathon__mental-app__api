package controller

import (
	"cmd/app/model"
	"cmd/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FetchDayReportRes struct {
	Id      string   `json:"mentalPointId"`
	Date    string   `json:"createDate"`
	Point   int      `json:"point"`
	Reasons []string `json:"reasonIdList"`
}

func FetchDayReport(c *gin.Context) {
	mentalPointId := c.Param("mentalPointId")
	userId := utils.GetValueFromContext(c, "userId")

	targetReport := model.MentalPoint{}
	if err := targetReport.GetReportByMentalPointId(mentalPointId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userId != targetReport.UserId {
		c.JSON(http.StatusForbidden, gin.H{"message": "not permitted"})
		return
	}

	targetReasons := make(model.ReasonIdList, 0)
	if err := targetReasons.GetReasonIdsByMentalPointId(mentalPointId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := FetchDayReportRes{
		Id:      mentalPointId,
		Date:    targetReport.CreatedDate,
		Point:   targetReport.Point,
		Reasons: targetReasons,
	}

	c.JSON(http.StatusOK, res)
}
