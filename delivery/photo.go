package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/szczynk/MyGram/middlewares"
	"github.com/szczynk/MyGram/models"
)

type photoRoutes struct {
	puc models.PhotoUsecase
}

func NewPhotoRoute(handlers *gin.Engine, puc models.PhotoUsecase) {
	route := &photoRoutes{puc}

	handler := handlers.Group("/photos")
	{
		handler.Use(middlewares.Authentication())
		handler.GET("/", route.Fetch)
		handler.POST("/", route.Store)
		handler.PUT("/:id", middlewares.PhotoAuthorization(route.puc), route.Update)
		handler.DELETE("/:id", middlewares.PhotoAuthorization(route.puc), route.Delete)
	}
}

// Fetch godoc
// @Summary      Fetch photos
// @Description  get photos
// @Tags         photos
// @Accept       json
// @Produce      json
// @Success      200	{object}	[]models.Photo
// @Failure      400	{object}	ErrorResponse
// @Failure      401	{object}	ErrorResponse
// @Security     Bearer
// @Router       /photos        [get]
func (route *photoRoutes) Fetch(c *gin.Context) {
	var (
		photos []models.Photo
		err    error
	)

	err = route.puc.Fetch(c.Request.Context(), &photos)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, photos)
}

// Store godoc
// @Summary      Create an photo
// @Description  create and store an photo
// @Tags         photos
// @Accept       json
// @Produce      json
// @Param        message  body  models.Photo true  "Photo"
// @Success      201  {object}  models.Photo
// @Failure      400  {object}	ErrorResponse
// @Failure      401  {object}	ErrorResponse
// @Security     Bearer
// @Router       /photos        [post]
func (route *photoRoutes) Store(c *gin.Context) {
	var (
		photo models.Photo
		err   error
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err = c.ShouldBindJSON(&photo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	photo.UserID = userID

	err = route.puc.Store(c.Request.Context(), &photo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         photo.ID,
		"user_id":    photo.UserID,
		"title":      photo.Title,
		"photo_url":  photo.PhotoUrl,
		"caption":    photo.Caption,
		"created_at": photo.CreatedAt,
	})
}

// Update godoc
// @Summary      Update an photo
// @Description  update an photo by ID
// @Tags         photos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Photo ID"
// @Success      200  {string}  string
// @Failure      400  {object}	ErrorResponse
// @Failure      401  {object}	ErrorResponse
// @Failure      404  {object}	ErrorResponse
// @Security     Bearer
// @Router       /photos/{id}   [put]
func (route *photoRoutes) Update(c *gin.Context) {
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

	err = c.ShouldBindJSON(&photo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	photoID := uint(photoIDInt)

	updatedPhoto := models.Photo{
		UserID:   userID,
		Title:    photo.Title,
		PhotoUrl: photo.PhotoUrl,
		Caption:  photo.Caption,
	}

	photo, err = route.puc.Update(c.Request.Context(), updatedPhoto, photoID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         photo.ID,
		"user_id":    photo.UserID,
		"title":      photo.Title,
		"photo_url":  photo.PhotoUrl,
		"caption":    photo.Caption,
		"updated_at": photo.UpdatedAt,
	})
}

// Delete godoc
// @Summary      Delete an photo
// @Description  delete an photo by ID
// @Tags         photos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Photo ID"
// @Success      200  {string}  string
// @Failure      400  {object}	ErrorResponse
// @Failure      401  {object}	ErrorResponse
// @Failure      404  {object}	ErrorResponse
// @Security     Bearer
// @Router       /photos/{id}   [delete]
func (route *photoRoutes) Delete(c *gin.Context) {
	photoIDStr := c.Param("id")
	photoIDInt, err := strconv.Atoi(photoIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Failed to cast photo id to int",
		})
		return
	}

	photoID := uint(photoIDInt)

	err = route.puc.Delete(c.Request.Context(), photoID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Your photo has been successfully deleted",
		},
	)
}
