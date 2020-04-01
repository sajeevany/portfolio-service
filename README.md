# Portfolio Service
Wrapper service for storing user portfolio information

Generate swagger files:

    swag init cmd/porfolio=service/main.go -g cmd/portfolio-service/main.go  -d ./

Access swagger API:

    ${HOST}:{PORT}/swagger/index.html
    ie. http://localhost:80/swagger/index.html

Run tests:

    go test ./...


