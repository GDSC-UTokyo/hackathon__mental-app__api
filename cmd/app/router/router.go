package router

import (
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
			"GET",
			"POST",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"authentication",
			"workspace_id",
		},
	}
	g.Use(cors.New(corsConfig))

	//g.POST("/signup", controller.Signup)みたいな感じで並べていく
	g.GET("/reasons", controller.FetchAllReasons)
	g.POST("/reasons", controller.CreateReason)
	g.PUT("/reasons/:reasonId", controller.UpdateReason)
	g.DELETE("/reasons/:reasonId", controller.DeleteReason)

	g.POST("/reports", controller.CreateReport)
	g.PUT("/reports/:mentalPointId", controller.UpdateReport)

	g.Run(":8080")
}
