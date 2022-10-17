package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/szczynk/MyGram/database"
	"github.com/szczynk/MyGram/delivery"
	docs "github.com/szczynk/MyGram/docs"
	"github.com/szczynk/MyGram/repository"
	"github.com/szczynk/MyGram/usecase"
)

func init() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if viper.GetBool(`debug`) {
		fmt.Println("service RUN on DEBUG mode")
	}
}

// @title MyGram API
// @version 1.0
// @description MyGram is a free photo sharing app written in Go. People can share, view, and comment photos by everyone. Anyone can create an account by registering an email address and selecting a username.
// @termOfService http://swagger.io/terms/
// @contact.name szczynk
// @contact.email szczynk@gmail.com
// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description					        Description for what is this security definition being used

func main() {
	db := database.StartDB()

	// HTTP Server
	port := fmt.Sprintf(":%s", viper.GetString("server.port"))
	routers := gin.Default()

	routers.GET("/health", CheckHealth)

	userRepo := repository.NewUserRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	delivery.NewUserRoute(routers, userUsecase)

	photoRepo := repository.NewPhotoRepo(db)
	photoUsecase := usecase.NewPhotoUsecase(photoRepo)
	delivery.NewPhotoRoute(routers, photoUsecase)

	commentRepo := repository.NewCommentRepo(db)
	commentUsecase := usecase.NewCommentUsecase(commentRepo)
	delivery.NewCommentRoute(routers, commentUsecase, photoUsecase)

	socialMediaRepo := repository.NewSocialMediaRepo(db)
	socialMediaUsecase := usecase.NewSocialMediaUsecase(socialMediaRepo)
	delivery.NewSocialMediaRoute(routers, socialMediaUsecase)

	docs.SwaggerInfo.BasePath = "/"
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	srv := &http.Server{
		Addr:    port,
		Handler: routers,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		// server connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

// CheckHealth godoc
// @Summary check health
// @Description do ping
// @Accept json
// @Produce json
// @Success 200
// @Router /health [get]
func CheckHealth(c *gin.Context) { c.Status(http.StatusOK) }
