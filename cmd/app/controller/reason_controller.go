package controller

import (
	"cmd/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FetchAllReasonsRes struct {
	Id     string `json:"id"`
	Reason string `json:"reason"`
}

type CreateReasonReq struct {
	Reason string `json:"reason"`
}

type CreateReasonRes struct {
	Id     string `json:"id"`
	Reason string `json:"reason"`
}

func FetchAllReasons(c *gin.Context) {
	userId := c.Request.Header.Get("UserId")

	targetReasons := model.Reasons{}
	if err := targetReasons.GetReasonsByUserId(userId).Error; err != nil {
		//エラー処理を書く
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := make([]FetchAllReasonsRes, 0)
	for i := 0; i < len(targetReasons); i++ {
		res = append(res, FetchAllReasonsRes{
			targetReasons[i].Id,
			targetReasons[i].Reason,
		})
	}

	c.JSON(http.StatusOK, res)
}

func CreateReason(c *gin.Context) {
	userId := c.Request.Header.Get("UserId")

	req := new(CreateReasonReq)
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newReason := model.Reason{
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
