package controllers

import (

	"github.com/azkifairuz/rfid-presence-api/helper"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/models"
	"github.com/gin-gonic/gin"
)

func GetPresenceByClass(c *gin.Context) {
	class := c.Param("class")
	
	var presence []models.PresenceResponse
	if result := initializers.DB.Table("presensis").
	Select("presensis.id as presence_id, mahasiswas.name as mahasiswa_name, mahasiswas.nim as mahasiswa_nim,mahasiswas.id as mahasiswa_id, kelas.name as mahasiswa_class").
	Joins("JOIN mahasiswas on presensis.mahasiswa_id = mahasiswas.id").
	Joins("JOIN kelas on mahasiswas.kelas_id = kelas.id").
	Where("kelas.name = ? ",class).
	Scan(&presence); result.Error != nil {
		helper.ResponseDefault(c, 404, nil, "presence not found")
		return
	}
	helper.ResponseDefault(c, 200, presence, "success get all data presence")


}