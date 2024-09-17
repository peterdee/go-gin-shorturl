package index

import (
	"github.com/gin-gonic/gin"

	"go-gin-url/utilities"
)

func indexController(context *gin.Context) {
	utilities.Response(utilities.ResponseOptions{Context: context})
}
