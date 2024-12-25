package controllers

import (

	"github.com/azkifairuz/rfid-presence-api/helper"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/models"
	"github.com/gin-gonic/gin"
)

func CreateMahasiswa(c *gin.Context)  {
	var body struct {
		NIM string `json:"nim" binding:"required"`
		Name string `json:"name" binding:"required"`
		ProdiId int `json:"prodiId" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		helper.ResponseDefault(c,400,nil,"All Field must be filled")
		return
	}
	var existingMahasiswa models.Mahasiswa
	if err := initializers.DB.Where("nim = ?", body.NIM).First(&existingMahasiswa).Error; err == nil {
		// If record exists, return error
		helper.ResponseDefault(c, 400, nil, "NIM already exists")
		return
	}

	var prodi models.Prodi
	if err := initializers.DB.Where("id = ?", body.ProdiId).First(&prodi).Error; err != nil {
		// If Fakultas doesn't exist, return error
		helper.ResponseDefault(c, 400, nil, "Prodi with the provided ID does not exist")
		return
	}
	//create a post
	mahasiswa := models.Mahasiswa{Nim: body.NIM,Name: body.Name,ProdiID: body.ProdiId}

	result := initializers.DB.Create(&mahasiswa)
	if result.Error !=nil {
		helper.ResponseDefault(c, 400, nil, "Error binding data")

		return
	}

	
	//return it
	c.JSON(200, gin.H{
		"data": mahasiswa,
		"message":"success created data",
	})

}


func GetAllMahasiswa(c *gin.Context)  {

	
	var mahasiswa []models.MhsModel
	initializers.DB.Table("mahasiswas").
	Select("mahasiswas.id as mahasiswa_id, mahasiswas.name as mahasiswa_name, prodis.name AS prodi_name, prodis.id AS prodi_id,  fakultas.name AS fakultas_name,fakultas.id AS fakultas_id").
    Joins("JOIN prodis ON mahasiswas.prodi_id = prodis.id").
    Joins("JOIN fakultas ON prodis.fakultas_id = fakultas.id").
    Scan(&mahasiswa)

	helper.ResponseDefault(c, 200, mahasiswa, "success get all data prodi")

}

func GetMahasiswa(c *gin.Context)  {
	id := c.Param("id")
	
	var mahasiwa models.MhsModel
	
	if result := initializers.DB.Table("mahasiswas").
	Select("mahasiswas.id as mahasiswa_id, mahasiswas.name as mahasiswa_name, prodis.name AS prodi_name, prodis.id AS prodi_id,  fakultas.name AS fakultas_name,fakultas.id AS fakultas_id").
    Joins("JOIN prodis ON mahasiswas.prodi_id = prodis.id").
    Joins("JOIN fakultas ON prodis.fakultas_id = fakultas.id").
	Where("mahasiswas.id = ?", id).
	Scan(&mahasiwa); result.Error != nil {
		helper.ResponseDefault(c, 404, nil, "Dosen not found")
		return
	}
	helper.ResponseDefault(c, 200, mahasiwa, "success get all data dosen")

}

