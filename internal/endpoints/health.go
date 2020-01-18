package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/portfolioService/internal/health"
)

const HealthGroup = "/health"
const helloEndpoint = "/hello"

//BuildHelloEndpoint - Builds hello endpoint. Endpoint returns standard 200 and hello message without need for auth
func BuildHelloEndpoint(handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      helloEndpoint,
		Handlers: append(handlers, health.Hello),
	}
}
