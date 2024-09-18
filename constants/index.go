package constants

const DEFAULT_MONGO_CONNECTION_STRING string = "mongodb://localhost:27017"
const DEFAULT_MONGO_DATABASE_NAME string = "ginurls"
const DEFAULT_PORT string = "5454"

var ENV_NAMES = EnvNames{
	MONGO_CONNECTION_STRING: "MONGO_CONNECTION_STRING",
	MONGO_DATABASE_NAME:     "MONGO_DATABASE_NAME",
	PORT:                    "PORT",
}

var INFO = Info{
	BadRequest:          "BAD_REQUEST",
	InternalServerError: "INTERNAL_SERVER_ERROR",
	InvalidData:         "INVALID_DATA",
	MissingData:         "MISSING_DATA",
	Ok:                  "OK",
}
