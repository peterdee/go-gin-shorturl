package utilities

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin-url/constants"
)

type ResponseOptions struct {
	Context *gin.Context
	Data    gin.H
	Info    string
	Status  int
}

type ResponseObject struct {
	Data    interface{}
	Info    string
	Request string
	Status  int
}

func Response(options ResponseOptions) {
	info := options.Info
	if info == "" {
		info = constants.INFO.Ok
	}
	status := options.Status
	if status == 0 {
		status = http.StatusOK
	}

	responseObject := gin.H{
		"info": info,
		"request": fmt.Sprintf(
			"%s [%s]",
			options.Context.Request.URL,
			options.Context.Request.Method,
		),
		"status": status,
	}

	if options.Data != nil {
		responseObject["data"] = options.Data
	}

	options.Context.Status(status)
	options.Context.JSON(status, responseObject)
}
