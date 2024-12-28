package controllers

import (
	"github.com/azkifairuz/rfid-presence-api/helper"
	"github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context)  {
	var body struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Type string `json:"type" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		helper.ResponseDefault(c,400,nil,"All Field must be filled")
		return
	}
	var isAccountExist models.Account
	if result := initializers.DB.Where("email = ? AND account_type = ?",body.Email,body.Type).First(&isAccountExist); result.Error != nil {
		helper.ResponseDefault(c, 404, nil, "user not found")
		return
	}


	if body.Password != isAccountExist.Password {
		helper.ResponseDefault(c, 400, nil, "wrong email or password")
		return
	}

	helper.ResponseDefault(c, 200, isAccountExist, "login success")
}

func ChangePassword(c *gin.Context) {
		var body struct {
			Email string `json:"email" binding:"required"`
			NewPassword string `json:"new_password" binding:"required"`
			Type string `json:"type" binding:"required"`
		}
	
		if err := c.Bind(&body); err != nil {
			helper.ResponseDefault(c,400,nil,"All Field must be filled")
			return
		}
		var isAccountExist models.Account
		if result := initializers.DB.Where("email = ? AND account_type = ?",body.Email,body.Type).First(&isAccountExist); result.Error != nil {
			helper.ResponseDefault(c, 404, nil, "user not found")
			return
		}
	
	
        initializers.DB.Model(&isAccountExist).Update("password",body.NewPassword)
	 
		helper.ResponseDefault(c, 400, isAccountExist, "change password success")
}