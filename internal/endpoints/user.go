package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/portfolioService/internal/users"
)

const UserGroup = "/user"
const User = "/"

//BuildGetUsersEndpoint -Builds get all users endpoint
func BuildGetUsersEndpoint(handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      User,
		Handlers: append(handlers, users.GetAllUsers),
	}
}

//BuildGetUserEndpoint - Build get user endpoint. Gets user based on ID
func BuildGetUserEndpoint(handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      User,
		Handlers: append(handlers, users.GetUser),
	}
}

//BuildAddUserEndpoint - Build add user endpoint. Adds a new user to the database
func BuildAddUserEndpoint(handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      User,
		Handlers: append(handlers, users.AddUser),
	}
}
