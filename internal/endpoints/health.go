package endpoints

import (
	"github.com/gin-gonic/gin"
	health "github.com/sajeevany/portfolio-service/internal/health/handlers"
	"github.com/sirupsen/logrus"
)

const HealthGroup = "/health"
const helloEndpoint = "/hello"

//BuildHelloEndpoint - Builds hello endpoint. Endpoint returns standard 200 and hello message without need for auth
func BuildHelloEndpoint(logger *logrus.Logger, handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      helloEndpoint,
		Handlers: append(handlers, health.Hello(logger)),
	}
}
