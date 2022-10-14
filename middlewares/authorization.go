package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/szczynk/MyGram/models"
)

func PhotoAuthorization(puc models.PhotoUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			photo models.Photo
			err   error
		)

		photoIDStr := c.Param("id")
		photoIDInt, err := strconv.Atoi(photoIDStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Failed to cast photo id to int",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		photoID := uint(photoIDInt)

		err = puc.GetByUserID(c.Request.Context(), &photo, photoID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": fmt.Sprintf("Photo with id %s doesn't exist", photoIDStr),
			})
			return
		}

		if photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You don't have permission to view or edit this photo",
			})
			return
		}
	}
}

func CommentEditDelAuthorization(cuc models.CommentUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			comment models.Comment
			err     error
		)

		commentIDStr := c.Param("id")
		commentIDInt, err := strconv.Atoi(commentIDStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Failed to cast comment id to int",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		commentID := uint(commentIDInt)

		err = cuc.GetByUserID(c.Request.Context(), &comment, commentID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": fmt.Sprintf("Comment with id %s doesn't exist", commentIDStr),
			})
			return
		}

		if comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You don't have permission to view or edit this comment",
			})
			return
		}
	}
}

func SocialMediaAuthorization(cuc models.SocialMediaUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			socialMedia models.SocialMedia
			err         error
		)

		socialMediaIDStr := c.Param("id")
		socialMediaIDInt, err := strconv.Atoi(socialMediaIDStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Failed to cast social media id to int",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		socialMediaID := uint(socialMediaIDInt)

		err = cuc.GetByUserID(c.Request.Context(), &socialMedia, socialMediaID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": fmt.Sprintf("SocialMedia with id %s doesn't exist", socialMediaIDStr),
			})
			return
		}

		if socialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You don't have permission to view or edit this social media",
			})
			return
		}
	}
}
