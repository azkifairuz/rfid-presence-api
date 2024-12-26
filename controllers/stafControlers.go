package controllers

import (
	"strings"

	"github.com/azkifairuz/rfid-presence-api/helper"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/models"
	"github.com/gin-gonic/gin"
)

func CreateStaf(c *gin.Context)  {
	var body struct {
		Nip string `json:"nip" binding:"required"`
		Name string `json:"name" binding:"required"`
		FakultasId int `json:"fakultasId" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		helper.ResponseDefault(c,400,nil,"All Field must be filled")
		return
	}
	var existingFakultas models.Prodi
	if err := initializers.DB.Where("nip = ?", body.Nip).First(&existingFakultas).Error; err == nil {
		// If record exists, return error
		helper.ResponseDefault(c, 400, nil, "Nip already exists")
		return
	}

	var fakultas models.Fakultas
	if err := initializers.DB.Where("id = ?", body.FakultasId).First(&fakultas).Error; err != nil {
		// If Fakultas doesn't exist, return error
		helper.ResponseDefault(c, 400, nil, "Fakultas with the provided ID does not exist")
		return
	}
	//create a post
	staf := models.Staf{
		Nip: body.Nip,
		Name: body.Name,
		FakultasID: body.FakultasId,
	}

	result := initializers.DB.Create(&staf)
	if result.Error !=nil {
		helper.ResponseDefault(c, 400, nil, "Error binding data")

		return
	}
	firstName := strings.Fields(body.Name)
	emailFormat := firstName[0] + "@uca.ac.id"
	account := models.Account{Email:emailFormat,Password: "stafdefault",AccountType: "staf",UserID:staf.ID }
	createAccount := initializers.DB.Create(&account)
	if createAccount.Error !=nil {
		helper.ResponseDefault(c, 400, nil, "Error binding data")

		return
	}

	
	//return it
	c.JSON(200, gin.H{
		"data": staf,
		"message":"success created data",
	})

}


func GetAllStaf(c *gin.Context)  {
	var staf []models.Staf
	initializers.DB.Joins("Fakultas").Find(&staf)

	helper.ResponseDefault(c, 200, staf, "success get all data staf")

}

func GetStaf(c *gin.Context)  {
	id := c.Param("id")

	var staf models.Prodi
	if result := initializers.DB.Table("stafs").Joins("Fakultas").First(&staf, id); result.Error != nil {
		helper.ResponseDefault(c, 404, nil, "staf not found")
		return
	}
	helper.ResponseDefault(c, 200, staf, "success get all data staf")


}