package controller

import (
	"fmt"
	"net/http"

	"github.com/risdatamamal/api-javaprojects/database"
	"github.com/risdatamamal/api-javaprojects/helpers"
	"github.com/risdatamamal/api-javaprojects/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	db := database.GetDB()
	User := models.User{}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"user_name": User.UserName,
	})
}

func LoginUser(ctx *gin.Context) {
	db := database.GetDB()
	User := models.User{}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	password := User.Password
	err := db.Debug().Where("email = ? ", User.Email).Take(&User).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"error":   "Unauthorized",
			"message": fmt.Sprintf("Email Not registered :%s", err.Error()),
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"error":   "Unauthorized",
			"message": "Wrong password",
		})
		return
	}

	token, err := helpers.GenerateToken(User.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   "Internal server error",
			"message": fmt.Sprintf("Error generating token :%s", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
