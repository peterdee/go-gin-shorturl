package index

import (
	"github.com/gin-gonic/gin"

	"go-gin-url/utilities"
)

func redirectController(context *gin.Context) {
	link, _ := context.Params.Get("link")

	// TODO: get original link and redirect instead of regular response
	utilities.Response(utilities.ResponseOptions{
		Context: context,
		Data: gin.H{
			"link": link,
		},
	})
}
