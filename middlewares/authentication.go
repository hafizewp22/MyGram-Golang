package middlewares

import (
	"net/http"
	"project_final/database"
	"project_final/helpers"
	"project_final/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helpers.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			return
		}

		c.Set("userData", claims)
		c.Next()
	}
}

func SosialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialMediaID, err := strconv.Atoi(c.Param("socialMediaID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Unauthorized",
				"error":   "Invalid Social Media ID data type",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		socialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&socialMedia, uint(socialMediaID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
				"error":   "Ypu are not allowed to access this social media",
			})
			return
		}

		if socialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "You are not allowed to access this social media",
			})
			return
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		PhotoID, err := strconv.Atoi(c.Param("PhotoID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Unauthorized",
				"error":   "Invalid Photo ID data type",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		photo := models.Photo{}

		err = db.Select("user_id").First(&photo, uint(PhotoID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
				"error":   "Ypu are not allowed to access this photo",
			})
			return
		}

		if photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "You are not allowed to access this photo",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		CommentID, err := strconv.Atoi(c.Param("CommentID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Unauthorized",
				"error":   "Invalid Comment ID data type",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		comment := models.Comment{}

		err = db.Select("user_id").First(&comment, uint(CommentID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
				"error":   "Ypu are not allowed to access this comment",
			})
			return
		}

		if comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "You are not allowed to access this comment",
			})
			return
		}

		c.Next()
	}
}
