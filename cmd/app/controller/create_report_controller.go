package controller

import (
	"cmd/app/model"
	"cmd/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateReportReq struct {
	Date    string   `json:"createdDate" binding:"required"`
	Point   int      `json:"point" binding:"required"`
	Reasons []string `json:"reasonIdList"`
}

type CreateReportRes struct {
	Id      string   `json:"mentalPointId"`
	Point   int      `json:"point"`
	Reasons []string `json:"reasonIdList"`
}

func CreateReport(c *gin.Context) {
	userId := c.Request.Header.Get("UserId")

	req := new(CreateReportReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.Point < 0 || req.Point > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "out of range"})
		return
	}

	pointId := utils.GenerateId()
	newMentalPoint := model.MentalPoint{
		Id:          pointId,
		Point:       req.Point,
		UserId:      userId,
		CreatedDate: req.Date,
	}

	if err := newMentalPoint.RegisterPoint().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newReasonsOnMentalPoints := make(model.ROMPs, 0)
	for i := 0; i < len(req.Reasons); i++ {
		reasonOnPointId := utils.GenerateId()
		newReasonsOnMentalPoints = append(newReasonsOnMentalPoints, model.ReasonsOnMentalPoints{
			Id:            reasonOnPointId,
			ReasonId:      req.Reasons[i],
			MentalPointId: pointId,
		})
	}

	if err := newReasonsOnMentalPoints.RegisterReasonsOnMentalPoints().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	registeredReasonIdList := make(model.ReasonIdList, 0)
	if err := registeredReasonIdList.GetReasonIdsByMentalPointId(newReasonsOnMentalPoints[0].MentalPointId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &CreateReportRes{
		Id:      newMentalPoint.Id,
		Point:   newMentalPoint.Point,
		Reasons: registeredReasonIdList,
	}

	c.JSON(http.StatusOK, res)
}
