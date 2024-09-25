package link

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julyskies/gohelpers"

	"go-gin-url/constants"
	"go-gin-url/mongodb"
	"go-gin-url/utilities"
)

// createLinkController godoc
// @Summary      	Create new short URL
// @Tags         	link
// @Param 				request body link.createLinkRequestPayload true "Request body"
// @Produce      	json
// @Success      	200 {object} utilities.ResponseObject{data=link.createLinkResponsePayload} "OK"
// @Failure				400 {object} utilities.ResponseObject{data=nil} "Missing required data"
// @Failure				500 {object} utilities.ResponseObject{data=nil} "Internal server error"
// @Router       	/api/link/create [post]
func createLinkController(ginContext *gin.Context) {
	var payload createLinkRequestPayload
	bindError := ginContext.ShouldBind(&payload)
	if bindError != nil {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.InternalServerError,
			Status:  http.StatusInternalServerError,
		})
		return
	}

	originalURL := strings.Trim(payload.OriginalURL, " ")
	if originalURL == "" {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.MissingData,
			Status:  http.StatusBadRequest,
		})
		return
	}

	_, parsingError := url.ParseRequestURI(originalURL)
	if parsingError != nil {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.InvalidData,
			Status:  http.StatusBadRequest,
		})
		return
	}

	shortID, cuidError := utilities.GenerateCUID()
	if cuidError != nil {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.InternalServerError,
			Status:  http.StatusInternalServerError,
		})
		return
	}
	newLink := mongodb.Link{
		CreatedAt:   int(gohelpers.MakeTimestampSeconds()),
		OriginalURL: originalURL,
		ShortID:     shortID,
	}

	password := strings.Trim(payload.Password, " ")
	if password != "" {
		hash, hashError := utilities.CreateHash(password)
		if hashError != nil {
			utilities.Response(utilities.ResponseOptions{
				Context: ginContext,
				Info:    constants.INFO.InternalServerError,
				Status:  http.StatusInternalServerError,
			})
			return
		}
		newLink.PasswordHash = hash
	}

	queryContext, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, insertError := mongodb.Links.InsertOne(queryContext, newLink)
	if insertError != nil {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.InternalServerError,
			Status:  http.StatusInternalServerError,
		})
		return
	}

	utilities.Response(utilities.ResponseOptions{
		Context: ginContext,
		Data:    gin.H{"shortID": shortID},
	})
}
