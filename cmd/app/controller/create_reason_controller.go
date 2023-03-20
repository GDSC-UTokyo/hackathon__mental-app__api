package controller

import (
	"cmd/app/model"
	"cmd/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateReasonReq struct {
	Reason string `json:"reason" binding:"required"`
}

type CreateReasonRes struct {
	Id     string `json:"id"`
	Reason string `json:"reason"`
}

func CreateReason(c *gin.Context) {
	userId := c.Request.Header.Get("UserId")

	req := new(CreateReasonReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	reasonId := utils.GenerateId()
	newReason := model.Reason{
		Id:     reasonId,
		Reason: req.Reason,
		UserId: userId,
	}

	if err := newReason.CreateReason().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &CreateReasonRes{
		Id:     newReason.Id,
		Reason: newReason.Reason,
	}

	c.JSON(http.StatusOK, res)
}
