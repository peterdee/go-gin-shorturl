package link

import "github.com/gin-gonic/gin"

func CreateRouter(app *gin.Engine) {
	router := app.Group("/api/link")

	router.DELETE("/", deleteLinkController)
	router.POST("/", createLinkController)
}
