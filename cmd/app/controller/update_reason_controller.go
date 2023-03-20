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

	req := new(UpdateReasonReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	renewReason := model.Reason{
		Id:     reasonId,
		Reason: req.Reason,
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
