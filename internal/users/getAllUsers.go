package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(ctx *gin.Context) {
	ctx.String(http.StatusOK, "")
}
