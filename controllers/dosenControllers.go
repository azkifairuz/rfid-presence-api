package controllers

import (

	"github.com/azkifairuz/rfid-presence-api/helper"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/models"
	"github.com/gin-gonic/gin"
)

func CreateDosen(c *gin.Context)  {
	var body struct {
		NIP string `json:"nip" binding:"required"`
		Name string `json:"name" binding:"required"`
		ProdiId int `json:"prodiId" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		helper.ResponseDefault(c,400,nil,"All Field must be filled")
		return
	}
	var existingDosen models.Dosen
	if err := initializers.DB.Where("nip = ?", body.NIP).First(&existingDosen).Error; err == nil {
		// If record exists, return error
		helper.ResponseDefault(c, 400, nil, "Nip name already exists")
		return
	}

	var prodi models.Prodi
	if err := initializers.DB.Where("id = ?", body.ProdiId).First(&prodi).Error; err != nil {
		// If Fakultas doesn't exist, return error
		helper.ResponseDefault(c, 400, nil, "Prodi with the provided ID does not exist")
		return
	}
	//create a post
	dosen := models.Dosen{Nip: body.NIP,Name: body.Name,ProdiID: body.ProdiId}

	result := initializers.DB.Create(&dosen)
	if result.Error !=nil {
		helper.ResponseDefault(c, 400, nil, "Error binding data")

		return
	}

	
	//return it
	c.JSON(200, gin.H{
		"data": dosen,
		"message":"success created data",
	})

}


func GetAllDosen(c *gin.Context)  {

	
	var dosen []models.DosenModel
	initializers.DB.Table("dosens").
	Select("dosens.id as dosen_id, dosens.name as dosen_name, prodis.name AS prodi_name,prodis.id AS prodis_id, fakultas.name AS fakultas_name,fakultas.id as fakultas_id").
    Joins("JOIN prodis ON dosens.prodi_id = prodis.id").
    Joins("JOIN fakultas ON prodis.fakultas_id = fakultas.id").
    Scan(&dosen)

	helper.ResponseDefault(c, 200, dosen, "success get all data prodi")

}

func GetDosen(c *gin.Context)  {
	id := c.Param("id")
	
	var dosen models.DosenModel
	
	if result := initializers.DB.Table("dosens").
	Select("dosens.id as dosen_id, dosens.name as dosen_name, prodis.name AS prodi_name,prodis.id AS prodis_id, fakultas.name AS fakultas_name,fakultas.id as fakultas_id").
    Joins("JOIN prodis ON dosens.prodi_id = prodis.id").
    Joins("JOIN fakultas ON prodis.fakultas_id = fakultas.id").
	Where("dosens.id = ?", id).
	Scan(&dosen); result.Error != nil {
		helper.ResponseDefault(c, 404, nil, "Dosen not found")
		return
	}
	helper.ResponseDefault(c, 200, dosen, "success get all data dosen")

}

