package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risdatamamal/api-javaprojects/database"
	"github.com/risdatamamal/api-javaprojects/helpers"
	"github.com/risdatamamal/api-javaprojects/models"
)

func GetHeader(ctx *gin.Context) {
	db := database.GetDB()
	Header := models.Header{}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&Header)
	} else {
		ctx.ShouldBind(&Header)
	}

	err := db.Debug().First(&Header).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    Header,
		"message": "Get header Success",
	})
}
