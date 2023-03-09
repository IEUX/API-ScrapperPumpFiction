package router

import (
	"github.com/gin-gonic/gin"
	"goScrapper/modules/database"
	"net/http"
)

func GetFunctions(router *gin.Engine, database *database.Database) {
	router.GET("/stations", func(request *gin.Context) {
		city := request.Query("city")
		data := database.GetStation("city", city)
		request.JSON(http.StatusOK, gin.H{
			"response": data,
		})
	})
}
