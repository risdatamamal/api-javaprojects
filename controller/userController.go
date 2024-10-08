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
	Role := models.Role{}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	err := db.Debug().Preload("Role").Create(&User).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	roleErr := db.Where("id = ?", User.RoleID).First(&Role).Error
	if roleErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Invalid Role",
			"message": "The specified role does not exist",
		})
		return
	}

	var users []map[string]interface{}
	RoleName := ""
	if User.Role != nil {
		RoleName = User.Role.RoleName
	} else {
		RoleName = Role.RoleName
	}

	users = append(users, gin.H{
		"id":        User.ID,
		"user_name": User.UserName,
		"email":     User.Email,
		"role":      RoleName,
	})

	ctx.JSON(http.StatusCreated, gin.H{
		"users": users,
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
	err := db.Debug().Preload("Role").Where("email = ? ", User.Email).Take(&User).Error
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

	if User.RoleID != 1 {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"error":   "Forbidden",
			"message": "You do not have permission to access authentication",
		})
		return
	}

	access_token, err := helpers.GenerateToken(User.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   "Internal server error",
			"message": fmt.Sprintf("Error generating token :%s", err.Error()),
		})
		return
	}

	var users []map[string]interface{}
	users = append(users, gin.H{
		"id":        User.ID,
		"user_name": User.UserName,
		"email":     User.Email,
		"role":      User.Role.RoleName,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"users":        users,
		"access_token": access_token,
	})
}
