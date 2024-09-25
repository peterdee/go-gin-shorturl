package utilities

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julyskies/gohelpers"

	"go-gin-url/constants"
)

type ResponseOptions struct {
	Context *gin.Context
	Data    gin.H
	Info    string
	Status  int
}

// Struct for Swagger
type ResponseObject struct {
	Data     interface{} `json:"data"`
	Datetime int         `json:"datetime"`
	Info     string      `json:"info"`
	Request  string      `json:"request"`
	Status   int         `json:"status"`
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
		"datetime": gohelpers.MakeTimestampSeconds(),
		"info":     info,
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
