package link

import "github.com/gin-gonic/gin"

func CreateRouter(app *gin.Engine) {
	router := app.Group("/api/link")

	router.POST("/create", createLinkController)
	router.POST("/delete", deleteLinkController)
}
