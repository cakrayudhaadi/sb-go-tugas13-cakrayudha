package controllers

import (
	"fmt"
	"net/http"
	"tugas13/database"

	"github.com/gin-gonic/gin"
)

type Bioskop struct {
	ID     int     `json:"id"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float32 `json:"rating"`
}

func CreateBioskop(ctx *gin.Context) {

	var newBioskop Bioskop

	if err := ctx.ShouldBindJSON(&newBioskop); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
	INSERT INTO tugas_13.bioskop(nama, lokasi, rating)
	VALUES ($1, $2, $3)
	`

	res, err := database.DBConnection.Exec(sqlStatement, newBioskop.Nama, newBioskop.Lokasi, newBioskop.Rating)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Created data amount:", count)
}
