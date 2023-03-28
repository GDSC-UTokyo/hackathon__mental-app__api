package router

import (
	"cmd/app/auth"
	"cmd/app/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	g := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"*",
		},
		AllowHeaders: []string{
			"*",
		},
	}
	g.Use(cors.New(corsConfig))
	g.Use(auth.Middleware())

	//g.POST("/signup", controller.Signup)みたいな感じで並べていく
	g.GET("/reasons", controller.FetchAllReasons)
	g.POST("/reasons", controller.CreateReason)
	g.PUT("/reasons/:reasonId", controller.UpdateReason)
	g.DELETE("/reasons/:reasonId", controller.DeleteReason)

	//g.POST("/reports", controller.CreateReport)
	g.PUT("/reports/:mentalPointId", controller.UpdateReport)
	g.GET("/reports/:mentalPointId", controller.FetchDayReport)
	g.POST("/reports", controller.FetchPeriodReports)

	g.Run(":8080")
}
