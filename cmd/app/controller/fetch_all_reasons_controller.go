package controller

import (
	"cmd/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FetchAllReasonsRes struct {
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
