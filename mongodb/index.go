package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/julyskies/gohelpers"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"

	"go-gin-url/constants"
	"go-gin-url/utilities"
)

var Client *mongo.Client
var Database *mongo.Database
var Links *mongo.Collection

func scheduledTasks() {
	schedule := cron.New()
	schedule.AddFunc(
		"@midnight",
		func() {
			queryContext, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			_, queryError := Links.DeleteMany(
				queryContext,
				bson.D{{
					Key: "createdAt",
					Value: bson.D{{
						Key:   "$lt",
						Value: gohelpers.MakeTimestampSeconds() - (14 * 24 * 60 * 60),
					}},
				}},
			)
			if queryError != nil {
				log.Fatal(queryError)
			}
		},
	)
	schedule.Start()
}

func Connect() {
	mongoConnectionString := utilities.GetEnv(
		constants.ENV_NAMES.MONGO_CONNECTION_STRING,
		constants.DEFAULT_MONGO_CONNECTION_STRING,
	)

	for s := 1; s <= 5; s += 1 {
		client, connectionError := mongo.Connect(
			options.Client().ApplyURI(mongoConnectionString),
		)
		if connectionError != nil {
			log.Printf("Failed to connect to MongoDB, retry in %d seconds", s)
			time.Sleep(time.Duration(s) * time.Second)
			continue
		}

		context, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		pingError := client.Ping(context, readpref.Primary())
		if pingError != nil {
			log.Printf("Failed to ping MongoDB server, retry in %d seconds", s)
			time.Sleep(time.Duration(s) * time.Second)
			continue
		}

		databaseName := utilities.GetEnv(
			constants.ENV_NAMES.MONGO_DATABASE_NAME,
			constants.DEFAULT_MONGO_DATABASE_NAME,
		)
		Client = client
		Database = client.Database(databaseName)
		Links = Database.Collection("links")

		scheduledTasks()

		log.Println("Connected to MongoDB")
		break
	}
}
