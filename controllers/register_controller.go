package controllers

import (
	"github.com/MuhammadSuryono1997/module-go/base/database"
	_http "github.com/MuhammadSuryono1997/module-go/base/http"
	"github.com/MuhammadSuryono1997/module-go/otp"
	"github.com/MuhammadSuryono1997/module-go/register/controllers"
	"github.com/MuhammadSuryono1997/module-go/register/models"
	"github.com/MuhammadSuryono1997/module-go/register/services"
	"github.com/MuhammadSuryono1997/module-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context)  {
	var registerservice services.RegisterService
	registercontroller := controllers.RegisterHandler(registerservice)

	no_hp, err := registercontroller.RegisterUser(c)

	if err != "" {
		failed := _http.ErrorCode{
			Code: _http.NumberIsRegistered,
			Message: err,
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, failed.AsInvalidResponse())
		return
	}

	success := _http.ErrorCode{
		Code:    _http.SuccessRegister,
		Message: _http.MessageSuccessRegister,
	}

	c.JSON(http.StatusCreated, success.AsValidResponse(utils.MaskedNumber(no_hp)))
}

func ResendOtp(c *gin.Context)  {
	var credential *models.TMerchant

	if err:= c.ShouldBindJSON(&credential); err != nil {
		failed := _http.ErrorCode{
			Code:    _http.BadRequest,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, failed.AsInvalidResponse())
		return
	}
	c.ShouldBindJSON(&credential)

	merchant := database.GetDb().Where("phone_number = ?", credential.PhoneNumber).Find(&credential)
	if merchant.RowsAffected == 0 {
		failed := _http.ErrorCode{
			Code:    _http.BadRequest,
			Message: _http.MessagePhoneNumberNotFound,
		}
		c.JSON(http.StatusNotFound, failed.AsInvalidResponse())
		return
	}

	send, err := otp.ResendOTP(credential.PhoneNumber)
	if err != nil {
		failed := _http.ErrorCode{
			Code:    _http.BadRequest,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, failed.AsInvalidResponse())
		return
	}

	success := _http.ErrorCode{
		Code:    20000,
		Message: _http.MessageSuccessRequest,
	}

	c.JSON(http.StatusOK, success.AsValidResponse(utils.MaskedNumber(send)))
}
