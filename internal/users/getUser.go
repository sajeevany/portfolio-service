package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "")
}
