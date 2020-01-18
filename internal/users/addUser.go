package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//AddUser - Adds user to the user database
func AddUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "")
}
