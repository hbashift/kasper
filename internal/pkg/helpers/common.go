package helpers

import (
	"github.com/gin-gonic/gin"
)

func GetToken(ctx *gin.Context) string {
	return ctx.Param("token")
}
