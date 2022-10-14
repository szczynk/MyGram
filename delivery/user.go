package delivery

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/szczynk/MyGram/helpers"
	"github.com/szczynk/MyGram/middlewares"
	"github.com/szczynk/MyGram/models"
)

type userRoutes struct {
	uuc models.UserUsecase
}

func NewUserRoute(handlers *gin.Engine, uuc models.UserUsecase) {
	route := &userRoutes{uuc}

	handler := handlers.Group("/users")
	{
		handler.POST("/register", route.Register)
		handler.POST("/login", route.Login)
		handler.PUT("/", middlewares.Authentication(), route.Update)
		handler.DELETE("/", middlewares.Authentication(), route.Delete)
	}
}

// Register godoc
// @Summary      Create an user
// @Description  create and store an user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        message  body   models.User   true  "User"
// @Success      201  {object}   models.User
// @Failure      400  {object}	 ErrorResponse
// @Failure      409  {object}	 ErrorResponse
// @Router       /users/register [post]
func (route *userRoutes) Register(c *gin.Context) {
	var (
		user models.User
		err  error
	)

	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = route.uuc.Register(c.Request.Context(), &user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": "You can't use invalid or duplicate emails and/or username",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"age":      user.Age,
	})
}

// Login godoc
// @Summary      Show an user
// @Description  get an user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        message  body   models.User   true  "User"
// @Success      200  {object}  models.User
// @Failure      400  {object}  ErrorResponse
// @Failure      401  {object}	ErrorResponse
// @Router       /users/login	  [get]
func (route *userRoutes) Login(c *gin.Context) {
	var (
		user  models.User
		err   error
		token string
	)

	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = route.uuc.Login(c.Request.Context(), &user)
	if err != nil {
		if strings.Contains(err.Error(), "invalid password") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	token, err = helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Update godoc
// @Summary      Update an user
// @Description  update an user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}	ErrorResponse
// @Failure      401  {object}	ErrorResponse
// @Security     Bearer
// @Router       /users [put]
func (route *userRoutes) Update(c *gin.Context) {
	var (
		user models.User
		err  error
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	updatedUser := models.User{
		Username: user.Username,
		Email:    user.Email,
	}

	user, err = route.uuc.Update(c.Request.Context(), updatedUser, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	})
}

// Delete godoc
// @Summary      Delete an user
// @Description  delete an user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {string}  string
// @Failure      400  {object}	ErrorResponse
// @Failure      401  {object}	ErrorResponse
// @Security     Bearer
// @Router       /users					[delete]
func (route *userRoutes) Delete(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := route.uuc.Delete(c, userID)
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
			"message": "Your account has been successfully deleted",
		},
	)
}
