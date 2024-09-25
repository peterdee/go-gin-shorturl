package constants

type EnvNames struct {
	ENABLE_SWAGGER          string
	ENV_SOURCE              string
	MONGO_CONNECTION_STRING string
	MONGO_DATABASE_NAME     string
	PORT                    string
}

type Info struct {
	BadRequest          string
	Forbidden           string
	InternalServerError string
	InvalidData         string
	InvalidPassword     string
	MissingData         string
	NotFound            string
	Ok                  string
}
