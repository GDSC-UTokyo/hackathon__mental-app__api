package controller

import (
	"cmd/app/model"
	"cmd/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateReportReq struct {
	Point   int      `json:"point" binding:"required"`
	Reasons []string `json:"reasonIdList"`
}

type UpdateReportRes struct {
	Id      string   `json:"mentalPointId"`
	Point   int      `json:"point"`
	Reasons []string `json:"reasonIdList"`
}

func UpdateReport(c *gin.Context) {
	mentalPointId := c.Param("mentalPointId")
	userId := utils.GetValueFromContext(c, "userId")

	req := new(UpdateReportReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.Point < 0 || req.Point > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "out of range"})
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

	if originalReport.Point != req.Point {
		renewReport := model.MentalPoint{
			Id:     mentalPointId,
			Point:  req.Point,
			UserId: userId,
		}

		if err := renewReport.UpdateMentalPoint().Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if err := model.DeleteReportsByPointIdAndReasonId(mentalPointId, req.Reasons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "report not found"})
		return
	}

	renewReasonsOnMentalPoint := make(model.ROMPs, 0)
	for i := 0; i < len(req.Reasons); i++ {
		reasonOnPointId := utils.GenerateId()
		renewReasonsOnMentalPoint = append(renewReasonsOnMentalPoint, model.ReasonsOnMentalPoints{
			Id:            reasonOnPointId,
			ReasonId:      req.Reasons[i],
			MentalPointId: mentalPointId,
		})
	}

	if err := renewReasonsOnMentalPoint.RegisterReasonsOnMentalPoints().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &UpdateReportRes{
		Id:      mentalPointId,
		Point:   req.Point,
		Reasons: req.Reasons,
	}

	c.JSON(http.StatusOK, res)
}
