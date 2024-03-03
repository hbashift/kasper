package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUUID(ctx *gin.Context) (*uuid.UUID, error) {
	stringId := ctx.Param("id")

	id, err := uuid.Parse(stringId)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func GetToken(ctx *gin.Context) string {
	return ctx.Param("token")
}
