package student

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (s *Service) UploadDissertation(ctx *gin.Context, token string, semester *mapping.UploadDissertation) error {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return errors.Wrap(err, "[Student]")
	}

	form, _ := ctx.MultipartForm()
	files := form.File["upload"]

	for _, file := range files {
		dst := fmt.Sprintf("./dissertations/%s/semester%d/%s",
			session.KasperID.String(), semester.Semester.SemesterNumber, file.Filename)

		err = ctx.SaveUploadedFile(file, dst)
		if err != nil {
			return errors.Wrap(err, "UploadDissertation()")
		}

		log.Println(file.Filename)
	}

	return nil
}
