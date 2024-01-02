package student

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (s *Service) DownloadDissertation(ctx *gin.Context, token string, dissertationID *mapping.DownloadDissertation) (*model.Dissertation, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, errors.Wrap(ErrNonValidToken, "[Student]")
	}

	dissertation, err := s.dRepo.GetDissertationData(ctx, s.db, session.KasperID, dissertationID.Semester)
	if err != nil {
		return nil, errors.Wrap(err, "DownloadDissertation()")
	}

	return dissertation, nil
}
