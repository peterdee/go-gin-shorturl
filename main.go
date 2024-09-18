package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	indexRouter "go-gin-url/api/index"
	linkRouter "go-gin-url/api/link"
	"go-gin-url/constants"
	"go-gin-url/mongodb"
	"go-gin-url/utilities"
)

func main() {
	if envError := godotenv.Load(); envError != nil {
		log.Fatal(envError)
	}

	mongodb.Connect()

	mode := utilities.GetEnv(gin.EnvGinMode, gin.DebugMode)
	gin.SetMode(mode)

	app := gin.Default()

	app.StaticFile("/favicon.ico", "./assets/favicon.ico")

	indexRouter.CreateRouter(app)
	linkRouter.CreateRouter(app)

	port := utilities.GetEnv(constants.ENV_NAMES.PORT, constants.DEFAULT_PORT)
	app.Run(fmt.Sprintf(":%s", port))
}
