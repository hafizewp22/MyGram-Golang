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

func GetAllPhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := []models.Photo{} // Initialize socialMedia as a slice

	err := db.Find(&photo).Error

	if err != nil {
		panic(err)
	}

	for i := range photo {
		user := models.APIUser{}
		errUser := db.Model(&models.User{}).Find(&user, "id=?", photo[i].UserID).Error
		if errUser != nil {
			panic(errUser)
		}
		photo[i].User = &user
	}

	ctx.JSON(http.StatusOK, photo)
}

func GetPhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	user := models.APIUser{}
	photo := models.Photo{}
	PhotoID, _ := strconv.Atoi(ctx.Param("PhotoID"))

	photo.UserID = uint(userData["id"].(float64))

	err := db.First(&photo, "id=?", PhotoID).Error

	if err != nil {
		panic(err)
	}

	errUser := db.Model(&models.User{}).First(&user, "id=?", photo.UserID).Error
	if errUser != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errUser)
		return
	}

	photo.User = &user

	ctx.JSON(http.StatusOK, photo)
}

func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	user := models.APIUser{}
	photo := models.Photo{}

	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(photo)

	photo.UserID = uint(userData["id"].(float64))

	err = db.Create(&photo).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	errUser := db.Model(&models.User{}).First(&user, "id=?", photo.UserID).Error
	if errUser != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errUser)
		return
	}

	photo.User = &user

	ctx.JSON(http.StatusCreated, photo)
}

func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	user := models.APIUser{}
	photo := models.Photo{}
	PhotoID, _ := strconv.Atoi(ctx.Param("PhotoID"))

	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	photo.UserID = uint(userData["id"].(float64))

	err = db.Model(&photo).Where("id=?", PhotoID).Updates(models.Photo{UserID: photo.UserID, Title: photo.Title, Caption: photo.Caption, PhotoURL: photo.PhotoURL}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	errUser := db.Model(&models.User{}).First(&user, "id=?", photo.UserID).Error
	if errUser != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errUser)
		return
	}

	photo.User = &user

	ctx.JSON(http.StatusOK, photo)
}

func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	photo := models.Photo{}
	PhotoID, _ := strconv.Atoi(ctx.Param("PhotoID"))

	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	photo.UserID = uint(userData["id"].(float64))

	err = db.Model(&photo).Where("id=?", PhotoID).Delete(&photo).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Success Delete")
}
