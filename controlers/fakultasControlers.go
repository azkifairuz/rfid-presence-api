package controlers

import (
	"github.com/azkifairuz/rfid-presence-api/helper"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/models"
	"github.com/gin-gonic/gin"
)

func CreateFakultas(c *gin.Context)  {
	// get req body
	var body struct {
		Name string
	}

	c.Bind(&body)
	//create a post
	fakultas := models.Fakultas{Name: body.Name}

	result := initializers.DB.Create(&fakultas)
	if result.Error !=nil {
		helper.ResponseDefault(c, 400, nil, "Error binding data")

		return
	}

	
	//return it
	c.JSON(200, gin.H{
		"data": fakultas,
		"message":"success created data",
	})
}

func GetAllFakultas(c *gin.Context)  {
	//get data
	var fakultas []models.Fakultas
	initializers.DB.Find(&fakultas)

	helper.ResponseDefault(c, 200, fakultas, "success get all data fakultas")

}


func GetFakultas(c *gin.Context)  {
	// get param
	id := c.Param("id")
	//get data
	var fakultas models.Fakultas
	if result := initializers.DB.First(&fakultas, id); result.Error != nil {
		helper.ResponseDefault(c, 404, nil, "Fakultas not found")
		return
	}
	helper.ResponseDefault(c, 200, fakultas, "success get  data fakultas")

}