package index

import "github.com/gin-gonic/gin"

func CreateRouter(app *gin.Engine) {
	router := app.Group("/")

	router.GET("/", indexController)
	router.GET("/api", indexController)
	router.GET("/:id", redirectLinkController)
}
