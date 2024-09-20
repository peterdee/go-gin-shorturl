## go-gin-shorturl

URL shortener written in Go with Gin & MongoDB

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

### License

[MIT](./LICENSE.md)
