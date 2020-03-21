from golang:1.13.8-alpine3.11 as builder
WORKDIR /app
ARG GIT_COMMIT
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-X main.GIT_COMMIT=$GIT_COMMIT -X main.CONFIG_FILE=/app/config/portfolio-service-conf.json" -o main cmd/portfolio-service/main.go
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init cmd/porfolio=service/main.go -g cmd/portfolio-service/main.go  -d ./

from alpine:latest
ARG CONFIG_FILE
RUN apk --no-cache add ca-certificates
RUN mkdir /app
RUN mkdir /app/config
RUN mkdir /app/logs
RUN mkdir /app/docs
COPY ${CONFIG_FILE} /app/config/portfolio-service-conf.json
COPY --from=builder /app/main /app/main
COPY --from=builder /app/docs /app/docs
RUN chmod 777 /app/main
EXPOSE 8080
ENTRYPOINT ["/app/main"]
