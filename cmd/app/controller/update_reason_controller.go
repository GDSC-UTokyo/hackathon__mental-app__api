package controller

import (
	"cmd/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateReasonReq struct {
	Reason string `json:"reason"`
}

type UpdateReasonRes struct {
	Id     string `json:"id"`
	Reason string `json:"reason"`
}

func UpdateReason(c *gin.Context) {
	reasonId := c.Param("reasonId")
	userId := c.Request.Header.Get("UserId")

	req := new(UpdateReasonReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var originalReason model.Reason
	originalReason.GetReasonByReasonId(reasonId)
	if userId != originalReason.UserId {
		c.JSON(http.StatusForbidden, gin.H{"message": "not permitted"})
		return
	}

	renewReason := model.Reason{
		Id:     reasonId,
		Reason: req.Reason,
		UserId: userId,
	}

	if err := renewReason.UpdateReason().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &UpdateReasonRes{
		Id:     renewReason.Id,
		Reason: renewReason.Reason,
	}

	c.JSON(http.StatusOK, res)
}
