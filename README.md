## go-gin-shorturl

URL shortener written in Golang with [Gin](https://github.com/gin-gonic/gin) & [MongoDB](https://github.com/mongodb/mongo-go-driver)

### Deploy

Golang v1.22 is required

```shell script
git clone https://github.com/peterdee/go-gin-shorturl
cd ./go-gin-shorturl
gvm use go1.22
go mod download
```

### Environment variables

This project uses `.env` file, see [.env.example](./.env.example) for details

### Launch

##### Without Docker

```shell script
go run ./
```

Alternatively launch with [Air](https://github.com/air-verse/air)

##### With Docker

```shell script
docker run -p 5454:5454 --env-file .env -it $(docker build -q .)
```

##### With Docker Compose

```shell script
docker compose up -d
```

### Swagger

Install [swag](https://github.com/swaggo/swag) and generate documentation

```shell script
go install github.com/swaggo/swag/cmd/swag@latest
cd ./go-gin-shorturl
swag init
```

Set `ENABLE_SWAGGER` environment variable to `true`

Launch the server and open http://localhost:5454/swagger/index.html

### License

[MIT](./LICENSE.md)
