package link

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"go-gin-url/constants"
	"go-gin-url/mongodb"
	"go-gin-url/utilities"
)

func deleteLinkController(ginContext *gin.Context) {
	var payload DeleteLinkPayload
	bindError := ginContext.ShouldBind(&payload)
	if bindError != nil {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.InternalServerError,
			Status:  http.StatusInternalServerError,
		})
		return
	}

	password := strings.Trim(payload.Password, " ")
	shortID := strings.Trim(payload.ShortID, " ")
	if password == "" || shortID == "" {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.MissingData,
			Status:  http.StatusBadRequest,
		})
		return
	}

	queryContext, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var link mongodb.Link
	queryError := mongodb.Links.FindOne(
		queryContext,
		bson.D{{Key: "shortID", Value: shortID}},
	).Decode(&link)
	if queryError != nil {
		info := constants.INFO.InternalServerError
		status := http.StatusInternalServerError
		if queryError == mongo.ErrNoDocuments {
			info = constants.INFO.InvalidData
			status = http.StatusBadRequest
		}
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    info,
			Status:  status,
		})
		return
	}

	if link.PasswordHash == "" {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.Forbidden,
			Status:  http.StatusForbidden,
		})
		return
	}

	passwordIsValid, hashError := utilities.CompareHashWithPlaintext(
		link.PasswordHash,
		password,
	)
	if hashError != nil {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.InternalServerError,
			Status:  http.StatusInternalServerError,
		})
		return
	}
	if !passwordIsValid {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.InvalidPassword,
			Status:  http.StatusUnauthorized,
		})
		return
	}

	_, queryError = mongodb.Links.DeleteOne(
		queryContext,
		bson.D{{Key: "shortID", Value: shortID}},
	)
	if queryError != nil {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.InternalServerError,
			Status:  http.StatusInternalServerError,
		})
		return
	}

	utilities.Response(utilities.ResponseOptions{Context: ginContext})
}
