package index

import (
	"github.com/gin-gonic/gin"

	"go-gin-url/utilities"
)

func indexController(ginContext *gin.Context) {
	utilities.Response(utilities.ResponseOptions{Context: ginContext})
}
