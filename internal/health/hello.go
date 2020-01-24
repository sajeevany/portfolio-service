package health

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Ping struct {
	Response string `json:"response" required:"true" description:"Server hello response" example:"hello"`
}

func Hello(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Println("Hello from within hello")

		//Set response
		ctx.JSON(http.StatusOK, Ping{Response: "hello"})
	}
}
