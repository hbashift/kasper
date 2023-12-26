package student

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (s *Service) UploadDissertation(ctx *gin.Context, token string, semester *mapping.UploadDissertation) error {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return errors.Wrap(err, "[Student]")
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		err = errors.Wrap(err, "Here")
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil
	}

	log.Println(file.Filename)

	dst := fmt.Sprintf("./dissertations/%s/semester%d/%s",
		session.KasperID.String(), semester.Semester.SemesterNumber, file.Filename)

	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		return errors.Wrap(err, "UploadDissertation()")
	}

	return nil
}
