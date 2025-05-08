package routers

import (
	"tugas13/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	router.POST("/bioskop", controllers.CreateBioskop)
	router.GET("/bioskop", controllers.GetAllBioskop)
	router.GET("/bioskop/:id", controllers.GetBioskop)
	router.PUT("/bioskop/:id", controllers.UpdateBioskop)
	router.DELETE("/bioskop/:id", controllers.DeleteBioskop)

	router.Run(":8080")
}
