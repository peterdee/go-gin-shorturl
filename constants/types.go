package constants

type EnvNames struct {
	MONGO_CONNECTION_STRING string
	MONGO_DATABASE_NAME     string
	PORT                    string
}

type Info struct {
	BadRequest          string
	InternalServerError string
	InvalidData         string
	MissingData         string
	Ok                  string
}
