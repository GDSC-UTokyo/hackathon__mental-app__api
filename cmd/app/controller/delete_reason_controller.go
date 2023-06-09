package controller

import (
	"cmd/app/model"
	"cmd/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteReasonRes struct {
	Message string `json:"message"`
}

func DeleteReason(c *gin.Context) {
	reasonId := c.Param("reasonId")
	userId := utils.GetValueFromContext(c, "userId")

	originalReason := model.Reason{}
	if err := originalReason.GetReasonByReasonId(reasonId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "reason not found"})
		return
	}

	if userId != originalReason.UserId {
		c.JSON(http.StatusForbidden, gin.H{"message": "not permitted"})
		return
	}

	renewReason := model.Reason{
		Id:     reasonId,
		UserId: userId,
	}

	if err := renewReason.DeleteReason().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	targetROMPs := model.ROMPs{}
	if err := targetROMPs.DeleteReportsByReasonId(reasonId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &DeleteReasonRes{
		Message: "reason successfully deleted",
	}

	c.JSON(http.StatusOK, res)
}
