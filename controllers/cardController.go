package controllers

import (
	"time"

	"github.com/azkifairuz/rfid-presence-api/helper"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/models"
	"github.com/gin-gonic/gin"
)

func RegisterCard(c *gin.Context)  {
	var body struct {
		CardUid string `json:"carduid" binding:"required"`
		Nim string `json:"nim" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		// Return error if the name field is missing or invalid
		helper.ResponseDefault(c, 400, nil, "all field is required")
		return
	}

	var mahasiswa models.Mahasiswa
	if err := initializers.DB.Where("nim = ?", body.Nim).First(&mahasiswa).Error; err != nil {
		// If Fakultas doesn't exist, return error
		helper.ResponseDefault(c, 400, nil, "mahasiswa with the provided ID does not exist")
		return
	}
	var existingcard models.MhsCard
	if err := initializers.DB.Where("mahasiswa_id = ?", mahasiswa.ID).First(&existingcard).Error; err == nil {
		// If record exists, return error
		helper.ResponseDefault(c, 400, nil, "this mahaiswas already exists")
		return
	}
	card := models.MhsCard{CardUid: body.CardUid,MahasiswaID: mahasiswa.ID}
	result := initializers.DB.Create(&card)
	if result.Error !=nil {
		helper.ResponseDefault(c, 400, nil, "Error binding data")

		return
	}

	helper.ResponseDefault(c, 200, result.RowsAffected, "succes register card")

}

func ReadCard(c *gin.Context) {
	cardUid := c.Param("uid")

	var card models.CardResponse
	result := initializers.DB.Table("mhs_cards").
	Select("mhs_cards.card_uid as uid, mahasiswas.nim as mahasiswa_nim,  mahasiswas.id as mahasiswa_id, mahasiswas.name as mahasiswa_name,kelas.id AS kelas_id,kelas.name AS mahasiswa_class,kelas.id AS kelas_id").
	Joins("JOIN mahasiswas on mhs_cards.mahasiswa_id = mahasiswas.id").
    Joins("JOIN kelas ON mahasiswas.kelas_id = kelas.id").
	Where("mhs_cards.card_uid = ? ",cardUid).
    Scan(&card)

	if result.RowsAffected == 0 {
		helper.ResponseDefault(c, 404, nil, "data not found")
		return
	}
	var existingPresensi models.Presensi
	currentDate := time.Now().Format("2006-01-02") // Format tanggal untuk dibandingkan
	checkResult := initializers.DB.Where("mahasiswa_id = ? AND DATE(date) = ?", card.MahasiswaId, currentDate).
		First(&existingPresensi)

	if checkResult.RowsAffected > 0 {
		helper.ResponseDefault(c, 400, nil, "already checked in for today")
		return
	}

	// Masukkan data presensi
	newPresensi := models.Presensi{
		MahasiswaID: card.MahasiswaId,
		KelasID:     card.KelasId,
		Date:        time.Now(),
	}

	if err := initializers.DB.Create(&newPresensi).Error; err != nil {
		helper.ResponseDefault(c, 500, nil, "failed to save presensi")
		return
	}

	helper.ResponseDefault(c, 200, card, "success  read card and presence record ")

}