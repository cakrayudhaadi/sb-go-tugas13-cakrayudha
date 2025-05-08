package controllers

import (
	"net/http"
	"strconv"
	"tugas13/commons"
	"tugas13/database"
	"tugas13/repositories"
	"tugas13/structs"

	"github.com/gin-gonic/gin"
)

func CreateBioskop(ctx *gin.Context) {

	var newBioskop structs.Bioskop
	var result gin.H

	if err := ctx.ShouldBindJSON(&newBioskop); err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	if commons.IsStringEmpty(newBioskop.Nama) {
		result = gin.H{
			"result": "parameter nama harus diisi",
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	if commons.IsStringEmpty(newBioskop.Lokasi) {
		result = gin.H{
			"result": "parameter lokasi harus diisi",
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	err := repositories.CreateBioskop(database.DBConnection, newBioskop)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": "data bioskop berhasil dibuat",
		}
	}
	ctx.JSON(http.StatusCreated, result)
}

func GetAllBioskop(ctx *gin.Context) {
	var result gin.H

	bioskops, err := repositories.GetAllBioskop(database.DBConnection)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": bioskops,
		}
	}
	ctx.JSON(http.StatusOK, result)
}

func GetBioskop(ctx *gin.Context) {

	var result gin.H
	id, _ := strconv.Atoi(ctx.Param("id"))

	count, err := repositories.GetBioskop(database.DBConnection, id)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": count,
		}
	}
	ctx.JSON(http.StatusOK, result)
}

func UpdateBioskop(ctx *gin.Context) {

	var newBioskop structs.Bioskop
	var result gin.H
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&newBioskop); err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	if commons.IsStringEmpty(newBioskop.Nama) {
		result = gin.H{
			"result": "parameter nama harus diisi",
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	if commons.IsStringEmpty(newBioskop.Lokasi) {
		result = gin.H{
			"result": "parameter lokasi harus diisi",
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	newBioskop.ID = id

	err := repositories.UpdateBioskop(database.DBConnection, newBioskop)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": "data bioskop berhasil diubah",
		}
	}
	ctx.JSON(http.StatusOK, result)
}

func DeleteBioskop(ctx *gin.Context) {

	var result gin.H
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := repositories.DeleteBioskop(database.DBConnection, id)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": "data bioskop berhasil dihapus",
		}
	}
	ctx.JSON(http.StatusOK, result)
}
