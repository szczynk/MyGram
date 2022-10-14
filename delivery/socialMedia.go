package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/szczynk/MyGram/middlewares"
	"github.com/szczynk/MyGram/models"
)

type socialMediaRoutes struct {
	cuc models.SocialMediaUsecase
}

func NewSocialMediaRoute(handlers *gin.Engine, cuc models.SocialMediaUsecase) {
	route := &socialMediaRoutes{cuc}

	handler := handlers.Group("/socialmedias")
	{
		handler.Use(middlewares.Authentication())
		handler.GET("/", route.Fetch)
		handler.POST("/", route.Store)
		handler.PUT("/:id", middlewares.SocialMediaAuthorization(route.cuc), route.Update)
		handler.DELETE("/:id", middlewares.SocialMediaAuthorization(route.cuc), route.Delete)
	}
}

// Fetch godoc
// @Summary      Fetch socialMedias
// @Description  get socialMedias
// @Tags         socialMedias
// @Accept       json
// @Produce      json
// @Success      200	{object}	[]models.SocialMedia
// @Failure      400	{object}	ErrorResponse
// @Failure      401	{object}	ErrorResponse
// @Security     Bearer
// @Router       /socialMedias  [get]
func (route *socialMediaRoutes) Fetch(c *gin.Context) {
	var (
		socialMedias []models.SocialMedia
		err          error
	)

	err = route.cuc.Fetch(c.Request.Context(), &socialMedias)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, socialMedias)
}

// Store godoc
// @Summary      Create an socialMedia
// @Description  create and store an socialMedia
// @Tags         socialMedias
// @Accept       json
// @Produce      json
// @Param        message  body  models.SocialMedia true  "SocialMedia"
// @Success      201  {object}  models.SocialMedia
// @Failure      400	{object}	ErrorResponse
// @Failure      401	{object}	ErrorResponse
// @Security     Bearer
// @Router       /socialMedias  [post]
func (route *socialMediaRoutes) Store(c *gin.Context) {
	var (
		socialMedia models.SocialMedia
		err         error
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err = c.ShouldBindJSON(&socialMedia)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	socialMedia.UserID = userID

	err = route.cuc.Store(c.Request.Context(), &socialMedia)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               socialMedia.ID,
		"user_id":          socialMedia.UserID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"created_at":       socialMedia.CreatedAt,
	})
}

// Update godoc
// @Summary      Update an socialMedia
// @Description  update an socialMedia by ID
// @Tags         socialMedias
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "SocialMedia ID"
// @Success      200  {string}  string
// @Failure      400  {object}	ErrorResponse
// @Failure      401  {object}	ErrorResponse
// @Failure      404  {object}	ErrorResponse
// @Security     Bearer
// @Router       /socialMedias/{id} [put]
func (route *socialMediaRoutes) Update(c *gin.Context) {
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

	err = c.ShouldBindJSON(&socialMedia)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	socialMediaID := uint(socialMediaIDInt)

	updatedSocialMedia := models.SocialMedia{
		UserID:         userID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}

	socialMedia, err = route.cuc.Update(c.Request.Context(), updatedSocialMedia, socialMediaID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               socialMedia.ID,
		"user_id":          socialMedia.UserID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"updated_at":       socialMedia.UpdatedAt,
	})
}

// Delete godoc
// @Summary      Delete an socialMedia
// @Description  delete an socialMedia by ID
// @Tags         socialMedias
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "SocialMedia ID"
// @Success      200  {string}  string
// @Failure      400  {object}	ErrorResponse
// @Failure      401  {object}	ErrorResponse
// @Failure      404  {object}	ErrorResponse
// @Security     Bearer
// @Router       /socialMedias/{id} [delete]
func (route *socialMediaRoutes) Delete(c *gin.Context) {
	socialMediaIDStr := c.Param("id")
	socialMediaIDInt, err := strconv.Atoi(socialMediaIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Failed to cast social media id to int",
		})
		return
	}

	socialMediaID := uint(socialMediaIDInt)

	err = route.cuc.Delete(c.Request.Context(), socialMediaID)
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
			"message": "Your social media has been successfully deleted",
		},
	)
}
