from golang:1.13.8-alpine3.11 as builder
WORKDIR /app
ARG GIT_COMMIT
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD CGO_ENABLED=0 go test -v ./...