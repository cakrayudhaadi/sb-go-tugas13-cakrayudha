package routers

import (
	"tugas13/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	router.POST("/bioskop", controllers.CreateBioskop)

	router.Run(":8080")
}
