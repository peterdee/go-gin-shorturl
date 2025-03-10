package index

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"go-gin-url/constants"
	"go-gin-url/mongodb"
	"go-gin-url/utilities"
)

// redirectLinkController godoc
// @Summary      	Redirect to the original URL
// @Tags         	index
// @Param 				short_id path string false "Short link ID"
// @Success      	301
// @Failure				400 {object} utilities.ResponseObject{data=nil} "Missing required data"
// @Failure				404 {object} utilities.ResponseObject{data=nil} "Record not found"
// @Failure				500 {object} utilities.ResponseObject{data=nil} "Internal server error"
// @Router       	/{short_id} [get]
func redirectLinkController(ginContext *gin.Context) {
	shortID, ok := ginContext.Params.Get("id")
	if !ok {
		utilities.Response(utilities.ResponseOptions{
			Context: ginContext,
			Info:    constants.INFO.MissingData,
			Status:  http.StatusBadRequest,
		})
		return
	}

	queryContext, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var linkRecord mongodb.Link
	queryError := mongodb.Links.FindOne(
		queryContext,
		bson.D{{Key: "shortID", Value: shortID}},
	).Decode(&linkRecord)
	if queryError != nil {
		if queryError == mongo.ErrNoDocuments {
			utilities.Response(utilities.ResponseOptions{
				Context: ginContext,
				Info:    constants.INFO.NotFound,
				Status:  http.StatusNotFound,
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

	ginContext.Redirect(http.StatusMovedPermanently, linkRecord.OriginalURL)
}
