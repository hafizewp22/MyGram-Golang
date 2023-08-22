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

func GetAllComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := []models.Comment{} // Initialize socialMedia as a slice

	err := db.Find(&comment).Error

	if err != nil {
		panic(err)
	}

	for i := range comment {
		user := models.APIUser{}
		errUser := db.Model(&models.User{}).Find(&user, "id=?", comment[i].UserID).Error
		if errUser != nil {
			panic(errUser)
		}
		comment[i].User = &user

		photo := models.Photo{}
		errPhoto := db.Model(&models.Photo{}).Find(&photo, "id=?", comment[i].PhotoID).Error
		if errPhoto != nil {
			panic(errPhoto)
		}
		comment[i].Photo = &photo
	}

	ctx.JSON(http.StatusOK, comment)
}

func GetComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	user := models.APIUser{}
	photo := models.Photo{}
	comment := models.Comment{}
	CommentID, _ := strconv.Atoi(ctx.Param("CommentID"))

	comment.UserID = uint(userData["id"].(float64))

	err := db.First(&comment, "id=?", CommentID).Error

	if err != nil {
		panic(err)
	}

	errUser := db.Model(&models.User{}).First(&user, "id=?", comment.UserID).Error
	if errUser != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errUser)
		return
	}

	comment.User = &user

	errPhoto := db.Model(&models.Photo{}).First(&photo, "id=?", comment.PhotoID).Error
	if errPhoto != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errPhoto)
		return
	}

	comment.Photo = &photo

	ctx.JSON(http.StatusOK, comment)
}

func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	user := models.APIUser{}
	photo := models.Photo{}
	comment := models.Comment{}

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(comment)

	comment.UserID = uint(userData["id"].(float64))

	err = db.Create(&comment).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	errUser := db.Model(&models.User{}).First(&user, "id=?", comment.UserID).Error
	if errUser != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errUser)
		return
	}

	comment.User = &user

	errPhoto := db.Model(&models.Photo{}).First(&photo, "id=?", comment.PhotoID).Error
	if errPhoto != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errPhoto)
		return
	}

	comment.Photo = &photo

	ctx.JSON(http.StatusCreated, comment)
}

func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	user := models.APIUser{}
	photo := models.Photo{}
	comment := models.Comment{}
	CommentID, _ := strconv.Atoi(ctx.Param("CommentID"))

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	comment.UserID = uint(userData["id"].(float64))

	err = db.Model(&comment).Where("id=?", CommentID).Updates(models.Comment{UserID: comment.UserID, PhotoID: comment.PhotoID, Name: comment.Name, Message: comment.Message}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	errUser := db.Model(&models.User{}).First(&user, "id=?", comment.UserID).Error
	if errUser != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errUser)
		return
	}

	comment.User = &user

	errPhoto := db.Model(&models.Photo{}).First(&photo, "id=?", comment.PhotoID).Error
	if errPhoto != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errPhoto)
		return
	}

	comment.Photo = &photo

	ctx.JSON(http.StatusOK, comment)
}

func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	comment := models.Comment{}
	CommentID, _ := strconv.Atoi(ctx.Param("CommentID"))

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	comment.UserID = uint(userData["id"].(float64))

	err = db.Model(&comment).Where("id=?", CommentID).Delete(&comment).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Success Delete")
}
