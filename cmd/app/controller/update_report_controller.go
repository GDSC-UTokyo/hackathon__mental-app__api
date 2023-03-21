package controller

import (
	"cmd/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateReportReq struct {
	Point   int    `json:"point"`
	Reasons string `json:"reasonIdList" binding:"required"`
}

type UpdateReportRes struct {
	Id      string `json:"mentalPointId"`
	Point   int    `json:"point"`
	Reasons string `json:"reasonIdList" binding:"required"`
}

func UpdateReport(c *gin.Context) {
	mentalPointId := c.Param("mentalPointId")
	userId := c.Request.Header.Get("UserId")

	req := new(UpdateReportReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	originalReport := model.MentalPoint{}
	if err := originalReport.GetReportByMentalPointId(mentalPointId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "report not found"})
		return
	}

	if userId != originalReport.UserId {
		c.JSON(http.StatusForbidden, gin.H{"message": "not permitted"})
		return
	}

	renewReport := model.MentalPoint{
		Id:     mentalPointId,
		Point:  req.Point,
		UserId: userId,
	}

	if err := renewReport.UpdateMentalPoint().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	originalReasonsOnMentalPoint := model.ReasonsOnMentalPoints{}
	if err := originalReasonsOnMentalPoint.GetReportByMentalPointId(mentalPointId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "report not found"})
		return
	}

	renewReasonsOnMentalPoint := model.ReasonsOnMentalPoints{
		ReasonId:      req.Reasons,
		MentalPointId: mentalPointId,
	}

	if err := renewReasonsOnMentalPoint.UpdateReasonsOnMentalPoint().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &UpdateReportRes{
		Id:      mentalPointId,
		Point:   renewReport.Point,
		Reasons: renewReasonsOnMentalPoint.ReasonId,
	}

	c.JSON(http.StatusOK, res)
}
