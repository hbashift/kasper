package student

import (
	"context"
)

type StudentCommonInfo struct {
}

func (s *Service) UpdateDissertationPage(ctx context.Context, token string, info StudentCommonInfo) error {
	// TODO mapping
	/*
		1) Валидация токена
		2) Достать сессию
		3) Достать student_id
		4) Upsert
	*/

	return nil
}

// TODO upsert dissertation plan and only dissertation
