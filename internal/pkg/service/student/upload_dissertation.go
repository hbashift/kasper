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

	pz, err := ctx.FormFile("pz")
	if err != nil {
		err = errors.Wrap(err, "Here")
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil
	}

	titul, err := ctx.FormFile("titul")
	if err != nil {
		err = errors.Wrap(err, "Here")
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil
	}

	log.Println(pz.Filename)

	dst := fmt.Sprintf("./dissertations/%s/semester%d/%s",
		session.KasperID.String(), semester.Semester.SemesterNumber, pz.Filename)

	err = ctx.SaveUploadedFile(pz, dst)
	if err != nil {
		return errors.Wrap(err, "UploadDissertation()")
	}

	dst = fmt.Sprintf("./dissertations/%s/semester%d/%s",
		session.KasperID.String(), semester.Semester.SemesterNumber, titul.Filename)

	err = ctx.SaveUploadedFile(titul, dst)
	if err != nil {
		return errors.Wrap(err, "UploadDissertation()")
	}

	return nil
}
