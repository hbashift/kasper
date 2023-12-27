package student

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (s *Service) UploadDissertation(ctx *gin.Context, token string, semester *mapping.UploadDissertation) error {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return errors.Wrap(err, "[Student]")
	}

	dirPath := fmt.Sprintf("./dissertations/%s/semester%d", session.KasperID.String(), semester.Semester.SemesterNumber)
	err = os.RemoveAll(dirPath)
	if err != nil {
		return errors.Wrap(err, "could not clean directory")
	}

	pz, err := ctx.FormFile("upload")
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

	return s.dRepo.UpsertDissertationData(ctx, s.db, &session.KasperID, semester.Semester.SemesterNumber, pz.Filename)
}
