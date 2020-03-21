package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/portfolio-service/internal/users/handlers"
	"github.com/sirupsen/logrus"
)

const UserGroup = "/user"
const GetUsers = "/"
const GetUser = "/:id"
const PostUser = "/:id"

//BuildGetUsersEndpoint -Builds get all users endpoint
func BuildGetUsersEndpoint(logger *logrus.Logger, handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      GetUsers,
		Handlers: append(handlers, users.GetAllUsers(logger)),
	}
}

//BuildGetUserEndpoint - Build get user endpoint. Gets user based on ID
func BuildGetUserEndpoint(logger *logrus.Logger, handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      GetUser,
		Handlers: append(handlers, users.GetUser(logger)),
	}
}

//BuildAddUserEndpoint - Build add user endpoint. Adds a new user to the database
func BuildAddUserEndpoint(logger *logrus.Logger, handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      PostUser,
		Handlers: append(handlers, users.AddUser(logger)),
	}
}
