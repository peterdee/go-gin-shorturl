package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	indexRouter "go-gin-url/api/index"
	linkRouter "go-gin-url/api/link"
	"go-gin-url/constants"
	"go-gin-url/docs"
	"go-gin-url/mongodb"
	"go-gin-url/utilities"
)

// @title         URL Shortener APIs
// @version       1.0
// @license.name  MIT
// @host					localhost:5454
// @BasePath  		/
func main() {
	if envError := godotenv.Load(); envError != nil {
		log.Fatal(envError)
	}

	mongodb.Connect()

	mode := utilities.GetEnv(gin.EnvGinMode, gin.DebugMode)
	gin.SetMode(mode)

	app := gin.Default()
	app.StaticFile("/favicon.ico", "./assets/favicon.ico")

	// Swagger
	if value := utilities.GetEnv(constants.ENV_NAMES.ENABLE_SWAGGER); value == "true" {
		docs.SwaggerInfo.BasePath = "/"
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	// APIs
	indexRouter.CreateRouter(app)
	linkRouter.CreateRouter(app)

	// Handle 404 error
	app.NoRoute(func(ginContext *gin.Context) {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.NotFound,
			Status:  http.StatusNotFound,
		})
	})

	port := utilities.GetEnv(constants.ENV_NAMES.PORT, constants.DEFAULT_PORT)
	app.Run(fmt.Sprintf(":%s", port))
}
