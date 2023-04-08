package controllers

import (
	"fmt"
	"net/http"
	"project_final/database"
	"project_final/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetAllSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMedia := models.SocialMedia{}

	socialMedia.UserID = uint(userData["id"].(float64))

	err := db.Find(&socialMedia).Error

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func GetSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMedia := models.SocialMedia{}
	socialMediaID, _ := strconv.Atoi(ctx.Param("socialMediaID"))

	socialMedia.UserID = uint(userData["id"].(float64))

	err := db.First(&socialMedia, "id=?", socialMediaID).Error

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMedia := models.SocialMedia{}

	err := ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(socialMedia)

	socialMedia.UserID = uint(userData["id"].(float64))

	err = db.Create(&socialMedia).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, socialMedia)
}

func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMedia := models.SocialMedia{}
	socialMediaID, _ := strconv.Atoi(ctx.Param("socialMediaID"))

	err := ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	socialMedia.UserID = uint(userData["id"].(float64))

	err = db.Model(&socialMedia).Where("id=?", socialMediaID).Updates(models.SocialMedia{Name: socialMedia.Name, SosialMediaURL: socialMedia.SosialMediaURL}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func DeleteSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	socialMedia := models.SocialMedia{}
	socialMediaID, _ := strconv.Atoi(ctx.Param("socialMediaID"))

	err := ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	socialMedia.UserID = uint(userData["id"].(float64))

	err = db.Model(&socialMedia).Where("id=?", socialMediaID).Delete(&socialMedia).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Success Delete")
}
