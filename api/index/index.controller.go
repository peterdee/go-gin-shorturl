package index

import (
	"github.com/gin-gonic/gin"

	"go-gin-url/utilities"
)

// indexController godoc
// @Summary      Handle index request
// @Tags         index
// @Produce      json
// @Success      200 {object} utilities.ResponseObject{data=nil} "OK"
// @Router       / [get]
func indexController(ginContext *gin.Context) {
	utilities.Response(utilities.ResponseOptions{Context: ginContext})
}
