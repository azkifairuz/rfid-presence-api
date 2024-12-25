package controllers

import (

	"github.com/azkifairuz/rfid-presence-api/helper"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/models"
	"github.com/gin-gonic/gin"
)

func CreateProdi(c *gin.Context)  {
	var body struct {
		Name string `json:"name" binding:"required"`
		FakultasId int `json:"fakultasId" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		helper.ResponseDefault(c,400,nil,"All Field must be filled")
		return
	}
	var existingprodi models.Prodi
	if err := initializers.DB.Where("name = ?", body.Name).First(&existingprodi).Error; err == nil {
		// If record exists, return error
		helper.ResponseDefault(c, 400, nil, "Prodi name already exists")
		return
	}

	var fakultas models.Fakultas
	if err := initializers.DB.Where("id = ?", body.FakultasId).First(&fakultas).Error; err != nil {
		// If Fakultas doesn't exist, return error
		helper.ResponseDefault(c, 400, nil, "Fakultas with the provided ID does not exist")
		return
	}
	//create a post
	prodi := models.Prodi{Name: body.Name,FakultasID: body.FakultasId}

	result := initializers.DB.Create(&prodi)
	if result.Error !=nil {
		helper.ResponseDefault(c, 400, nil, "Error binding data")

		return
	}

	
	//return it
	c.JSON(200, gin.H{
		"data": prodi,
		"message":"success created data",
	})

}


func GetAllProdi(c *gin.Context)  {
	var prodi []models.Prodi
	initializers.DB.Joins("Fakultas").Find(&prodi)

	helper.ResponseDefault(c, 200, prodi, "success get all data prodi")

}

func GetProdi(c *gin.Context)  {
	var prodi models.Prodi
	initializers.DB.Joins("Fakultas").First(&prodi)

	helper.ResponseDefault(c, 200, prodi, "success get all data prodi")

}