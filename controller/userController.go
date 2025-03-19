package controller

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/risdatamamal/api-javaprojects/database"
	"github.com/risdatamamal/api-javaprojects/helpers"
	"github.com/risdatamamal/api-javaprojects/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	db := database.GetDB()
	User := models.User{}
	userRole := models.Role{}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	roleErr := db.Where("role_name = ?", "User").Where("guard_name = ?", "web").First(&userRole).Error
	if roleErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": "User web role not found",
		})
		return
	}

	existingUser := models.User{}
	emailErr := db.Where("email = ?", User.Email).First(&existingUser).Error
	if emailErr == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": "Email already registered",
		})
		return
	}

	User.Role = &userRole
	userErr := db.Debug().Preload("Role").Create(&User).Error
	if userErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": userErr.Error(),
		})
		return
	}

	accessToken, accessErr := helpers.GenerateToken(User.ID)
	if accessErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   "Internal server error",
			"message": fmt.Sprintf("Error generating token :%s", accessErr.Error()),
		})
		return
	}

	user := gin.H{
		"id":        User.ID,
		"user_name": User.UserName,
		"email":     User.Email,
		"role":      User.Role.RoleName,
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":         http.StatusCreated,
		"data":         user,
		"access_token": accessToken,
		"message":      "Register user success",
	})
}

func LoginUser(ctx *gin.Context) {
	db := database.GetDB()
	User := models.User{}
	userRole := models.Role{}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	password := User.Password
	userErr := db.Debug().Preload("Role").Where("email = ?", User.Email).First(&User).Error
	if userErr != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"error":   "Unauthorized",
			"message": fmt.Sprintf("Email Not registered :%s", userErr.Error()),
		})
		return
	}

	if !User.IsActive {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"error":   "Unauthorized",
			"message": "User is not active",
		})
		return
	}

	roleErr := db.Where("role_name = ?", "User").Where("guard_name = ?", "web").First(&userRole).Error
	if roleErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": "User web role not found",
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

	if User.RoleID != userRole.ID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"error":   "Forbidden",
			"message": "You do not have permission to access authentication",
		})
		return
	}

	accessToken, accessErr := helpers.GenerateToken(User.ID)
	if accessErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   "Internal server error",
			"message": fmt.Sprintf("Error generating token :%s", accessErr.Error()),
		})
		return
	}

	user := gin.H{
		"id":        User.ID,
		"user_name": User.UserName,
		"email":     User.Email,
		"role":      User.Role.RoleName,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"data":         user,
		"access_token": accessToken,
		"message":      "Login success",
	})
}

func LoginAdminUser(ctx *gin.Context) {
	db := database.GetDB()
	User := models.User{}
	adminRole := models.Role{}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	password := User.Password
	userErr := db.Debug().Preload("Role").Where("email = ?", User.Email).First(&User).Error
	if userErr != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"error":   "Unauthorized",
			"message": fmt.Sprintf("Email Not registered :%s", userErr.Error()),
		})
		return
	}

	if !User.IsActive {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"error":   "Unauthorized",
			"message": "User is not active",
		})
		return
	}

	roleErr := db.Where("role_name = ?", "Admin").Where("guard_name = ?", "web").First(&adminRole).Error
	if roleErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": "Admin web role not found",
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

	if User.RoleID != adminRole.ID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"error":   "Forbidden",
			"message": "You do not have permission to access authentication",
		})
		return
	}

	accessToken, accessErr := helpers.GenerateToken(User.ID)
	if accessErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   "Internal server error",
			"message": fmt.Sprintf("Error generating token :%s", accessErr.Error()),
		})
		return
	}

	user := gin.H{
		"id":        User.ID,
		"user_name": User.UserName,
		"email":     User.Email,
		"role":      User.Role.RoleName,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"data":         user,
		"access_token": accessToken,
		"message":      "Login success",
	})
}

func GetProfile(ctx *gin.Context) {
	db := database.GetDB()
	User := models.User{}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	userErr := db.Debug().Preload("Role").Where("id = ?", userID).First(&User).Error
	if userErr != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   "Not Found",
			"message": fmt.Sprintf("User not found :%s", userErr.Error()),
		})
		return
	}

	user := gin.H{
		"id":                User.ID,
		"user_name":         User.UserName,
		"email":             User.Email,
		"is_active":         User.IsActive,
		"email_verified_at": User.EmailVerifiedAt,
		"photo_path":        User.PhotoPath,
		"role_id":           User.RoleID,
		"role":              User.Role.RoleName,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    user,
		"message": "Get profile success",
	})
}
