package constants

const DEFAULT_MONGO_CONNECTION_STRING string = "mongodb://localhost:27017"
const DEFAULT_MONGO_DATABASE_NAME string = "ginurls"
const DEFAULT_PORT string = "5454"

var ENV_NAMES = EnvNames{
	ENABLE_SWAGGER:          "ENABLE_SWAGGER",
	MONGO_CONNECTION_STRING: "MONGO_CONNECTION_STRING",
	MONGO_DATABASE_NAME:     "MONGO_DATABASE_NAME",
	PORT:                    "PORT",
}

var INFO = Info{
	BadRequest:          "BAD_REQUEST",
	Forbidden:           "FORBIDDEN",
	InternalServerError: "INTERNAL_SERVER_ERROR",
	InvalidData:         "INVALID_DATA",
	InvalidPassword:     "INVALID_PASSWORD",
	MissingData:         "MISSING_DATA",
	NotFound:            "NOT_FOUND",
	Ok:                  "OK",
}
