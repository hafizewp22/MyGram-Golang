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
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	comment := models.Comment{}

	comment.UserID = uint(userData["id"].(float64))

	err := db.Find(&comment).Error

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, comment)
}

func GetComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	comment := models.Comment{}
	CommentID, _ := strconv.Atoi(ctx.Param("CommentID"))

	comment.UserID = uint(userData["id"].(float64))

	err := db.First(&comment, "id=?", CommentID).Error

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, comment)
}

func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
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

	ctx.JSON(http.StatusCreated, comment)
}

func UpdateComment(ctx *gin.Context) {
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

	err = db.Model(&comment).Where("id=?", CommentID).Updates(models.Comment{UserID: comment.UserID, PhotoID: comment.PhotoID, Name: comment.Name, Message: comment.Message}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

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
