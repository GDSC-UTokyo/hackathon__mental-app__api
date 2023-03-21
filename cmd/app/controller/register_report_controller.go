package controller

import (
	"cmd/app/model"
	"cmd/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateReportReq struct {
	Point   int    `json:"point" binding:"required"`
	Reasons string `json:"reasonIdList" binding:"required"`
}

type CreateReportRes struct {
	Id      string `json:"mentalPointId"`
	Point   int    `json:"point"`
	Reasons string `json:"reasonIdList"`
}

func CreateReport(c *gin.Context) {
	userId := c.Request.Header.Get("UserId")

	req := new(CreateReportReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pointId := utils.GenerateId()
	newMentalPoint := model.MentalPoint{
		Id:     pointId,
		Point:  req.Point,
		UserId: userId,
	}

	if err := newMentalPoint.RegisterPoint().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	reasonOnPointId := utils.GenerateId()
	newReasonsOnMentalPoints := model.ReasonsOnMentalPoints{
		Id:            reasonOnPointId,
		ReasonId:      req.Reasons,
		MentalPointId: pointId,
	}

	if err := newReasonsOnMentalPoints.RegisterReasonsOnMentalPoint().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &CreateReportRes{
		Id:      newMentalPoint.Id,
		Point:   newMentalPoint.Point,
		Reasons: newReasonsOnMentalPoints.MentalPointId,
	}

	c.JSON(http.StatusOK, res)
}
