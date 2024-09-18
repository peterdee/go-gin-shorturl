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

	shortID := strings.Trim(payload.ShortID, " ")
	if shortID == "" {
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
		if queryError == mongo.ErrNoDocuments {
			utilities.Response(utilities.ResponseOptions{
				Context: ginContext,
				Info:    constants.INFO.InvalidData,
				Status:  http.StatusBadRequest,
			})
			return
		}
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.InternalServerError,
			Status:  http.StatusInternalServerError,
		})
		return
	}

	utilities.Response(utilities.ResponseOptions{Context: ginContext})
}
