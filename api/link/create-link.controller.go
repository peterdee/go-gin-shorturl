package link

import (
	"github.com/gin-gonic/gin"

	"go-gin-url/utilities"
)

func createLinkController(context *gin.Context) {
	utilities.Response(utilities.ResponseOptions{Context: context})
}
