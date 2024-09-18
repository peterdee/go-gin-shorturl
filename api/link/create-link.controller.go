package link

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/julyskies/gohelpers"

	"go-gin-url/constants"
	"go-gin-url/mongodb"
	"go-gin-url/utilities"
)

func createLinkController(ginContext *gin.Context) {
	var payload CreateLinkPayload
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
		UpdatedAt:   int(gohelpers.MakeTimestampSeconds()),
	}

	password := strings.Trim(payload.Password, " ")
	if password != "" {
		hash, hashError := argon2id.CreateHash(password, argon2id.DefaultParams)
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
