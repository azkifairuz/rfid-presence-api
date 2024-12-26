package controllers

import (

	"time"

	"github.com/azkifairuz/rfid-presence-api/helper"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/models"
	"github.com/gin-gonic/gin"
)

func CreateJadwal(c *gin.Context)  {
	var body struct {
		Jam string `json:"jam" binding:"required"`
		Matkul string `json:"matkul" binding:"required"`
		Hari string `json:"hari" binding:"required"`
		KelasID int `json:"kelasId" binding:"required"`
		DosenID int `json:"dosenId" binding:"required"`

	}

	if err := c.Bind(&body); err != nil {
		helper.ResponseDefault(c,400,nil,"All Field must be filled")
		return
	}
	var existingMatkul models.Jadwal
	if err := initializers.DB.Where("matkul = ?", body.Matkul).First(&existingMatkul).Error; err == nil {
		// If record exists, return error
		helper.ResponseDefault(c, 400, nil, "matkul name already exists")
		return
	}	
	var kelas models.Kelas
	if err := initializers.DB.Where("id = ?", body.KelasID).First(&kelas).Error; err != nil {
		// If Fakultas doesn't exist, return error
		helper.ResponseDefault(c, 400, nil, "Kelas with the provided ID does not exist")
		return
	}

	var dosen models.Dosen
	if err := initializers.DB.Where("id = ?", body.DosenID).First(&dosen).Error; err != nil {
		// If Fakultas doesn't exist, return error
		helper.ResponseDefault(c, 400, nil, "Dosen with the provided ID does not exist")
		return
	}
	_, err := time.Parse("15:04", body.Jam)
	if err != nil {
		helper.ResponseDefault(c, 400, nil, "worng hour format")
		return
	}

	//create a post
	jadwal := models.Jadwal{
		Hour: body.Jam,
		Days: body.Hari,
		Matkul: body.Matkul,
		KelasID: body.KelasID,
		DosenID: dosen.ID,
	}

	result := initializers.DB.Create(&jadwal)
	if result.Error !=nil {
		helper.ResponseDefault(c, 400, nil, "Error binding data")

		return
	}

	//return it
	helper.ResponseDefault(c, 200, nil, "success create data")


}

func GetAllJadwal(c *gin.Context)  {

	
	var jadwal []models.JadwalModel
	initializers.DB.Table("jadwals").
	Select("jadwals.id as jadwalId,jadwals.matkul,dosens.name as dosen kelas.id as kelas_id, kelas.name as kelas, prodis.name AS prodi, prodis.id AS prodi_id,  fakultas.name AS fakultas,fakultas.id AS fakultas_id").
    Joins("JOIN dosens ON jadwals.dosen_id = dosens.id").
    Joins("JOIN prodis ON kelas.prodi_id = prodis.id").
	Joins("JOIN prodis ON kelas.prodi_id = prodis.id").
    Joins("JOIN fakultas ON prodis.fakultas_id = fakultas.id").
    Scan(&jadwal)

	helper.ResponseDefault(c, 200, jadwal, "success get all data jadwal")

}

func GetJadwal(c *gin.Context)  {

	id := c.Param("id")
	
	var jadwal []models.JadwalModel
	initializers.DB.Table("jadwals").
	Select("jadwals.id as jadwalId,jadwals.matkul,dosens.name as dosen kelas.id as kelas_id, kelas.name as kelas, prodis.name AS prodi, prodis.id AS prodi_id,  fakultas.name AS fakultas,fakultas.id AS fakultas_id").
    Joins("JOIN dosens ON jadwals.dosen_id = dosens.id").
    Joins("JOIN prodis ON kelas.prodi_id = prodis.id").
	Joins("JOIN prodis ON kelas.prodi_id = prodis.id").
    Joins("JOIN fakultas ON prodis.fakultas_id = fakultas.id").
	Where("jadwals.id =?",id).
    Scan(&jadwal)

	helper.ResponseDefault(c, 200, jadwal, "success get  data jadwal")
}