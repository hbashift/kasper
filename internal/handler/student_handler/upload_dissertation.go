package student_handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (h *studentHandler) UploadDissertation(ctx *gin.Context) {
	token, err := getUUID(ctx)
	if err != nil {
		err = errors.Wrap(err, "Here 2")
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reqBody := mapping.UploadDissertation{}
	err = ctx.ShouldBind(&reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Printf("%+v", reqBody)

	err = h.service.UploadDissertation(ctx, token.String(), &reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
