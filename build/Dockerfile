from golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main cmd/main.go

from alpine:latest
ARG VERSION_FILE
ARG CONFIG_FILE
RUN apk --no-cache add ca-certificates
RUN mkdir /app
RUN mkdir /app/config
RUN mkdir /app/logs
COPY ${VERSION_FILE} /app/version.txt
COPY ${CONFIG_FILE} /app/config/portfolio-service-conf.json
COPY --from=builder /app/main /app/main
RUN chmod 777 /app/main
EXPOSE 8080
ENTRYPOINT ["/app/main"]
