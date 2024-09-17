package link

import (
	"github.com/gin-gonic/gin"

	"go-gin-url/utilities"
)

func deleteLinkController(context *gin.Context) {
	utilities.Response(utilities.ResponseOptions{Context: context})
}
