package controllers

import (

	"github.com/azkifairuz/rfid-presence-api/helper"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/models"
	"github.com/gin-gonic/gin"
)

func CreateKelas(c *gin.Context)  {
	var body struct {
		Name string `json:"name" binding:"required"`
		ProdiId int `json:"prodiId" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		helper.ResponseDefault(c,400,nil,"All Field must be filled")
		return
	}
	var existingKelas models.Mahasiswa
	if err := initializers.DB.Where("name = ?", body.Name).First(&existingKelas).Error; err == nil {
		// If record exists, return error
		helper.ResponseDefault(c, 400, nil, "kelas already exists")
		return
	}

	var prodi models.Prodi
	if err := initializers.DB.Where("id = ?", body.ProdiId).First(&prodi).Error; err != nil {
		// If Fakultas doesn't exist, return error
		helper.ResponseDefault(c, 400, nil, "Prodi with the provided ID does not exist")
		return
	}
	//create a post
	kelas := models.Kelas{Name: body.Name,ProdiID: body.ProdiId}

	result := initializers.DB.Create(&kelas)
	if result.Error !=nil {
		helper.ResponseDefault(c, 400, nil, "Error binding data")

		return
	}

	
	//return it
	c.JSON(200, gin.H{
		"data": kelas,
		"message":"success created data",
	})

}


func GetAllKelas(c *gin.Context)  {

	
	var kelas []models.MhsModel
	initializers.DB.Table("kelas").
	Select("kelas.id as kelas_id, kelas.name as kelas_name, prodis.name AS prodi_name, prodis.id AS prodi_id,  fakultas.name AS fakultas_name,fakultas.id AS fakultas_id").
    Joins("JOIN prodis ON kelas.prodi_id = prodis.id").
    Joins("JOIN fakultas ON prodis.fakultas_id = fakultas.id").
    Scan(&kelas)

	helper.ResponseDefault(c, 200, kelas, "success get all data kelas")

}

func GetKelas(c *gin.Context)  {
	id := c.Param("id")
	
	var kelas models.MhsModel
	
	if result := initializers.DB.Table("kelas").
	Select("kelas.id as kelas_id, kelas.name as kelas_name, prodis.name AS prodi_name, prodis.id AS prodi_id,  fakultas.name AS fakultas_name,fakultas.id AS fakultas_id").
    Joins("JOIN prodis ON kelas.prodi_id = prodis.id").
    Joins("JOIN fakultas ON prodis.fakultas_id = fakultas.id").
	Where("kelas.id = ?", id).
	Scan(&kelas); result.Error != nil {
		helper.ResponseDefault(c, 404, nil, "Kelas not found")
		return
	}
	helper.ResponseDefault(c, 200, kelas, "success get all data kelas")

}

