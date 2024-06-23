package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func (s *Service) GenerateReportOne(ctx context.Context) ([]models.ReportOne, error) {
	var studentsInfo []models.StudentInfoForReportOne
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		studentIDs, err := s.clientRepo.GetAllStudentIDs(ctx, tx)
		if err != nil {
			return err
		}

		info, err := s.clientRepo.GetDataForReportOne(ctx, tx, studentIDs)
		if err != nil {
			return err
		}

		studentsInfo = info

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "GenerateReportOne()")
	}

	reports := make([]models.ReportOne, 0, len(studentsInfo))

	for _, student := range studentsInfo {
		report := models.ReportOne{
			FullName:        student.Student.FullName,
			SupervisorName:  student.SupervisorName,
			AttestationMark: student.AttestationMark,
			Progressiveness: student.Progressiveness,
			ScientificWorkCount: func() string {
				count := models.ScientificWorkCount{
					WAC:    0,
					Scopus: 0,
					Rinc:   0,
					Wos:    0,
				}

				for _, publication := range student.Publications {
					if publication.Wac {
						count.WAC++
					}

					if publication.Wos {
						count.Wos++
					}

					if publication.Scopus {
						count.Scopus++
					}

					if publication.Rinc {
						count.Rinc++
					}
				}

				for _, conference := range student.Conferences {
					if conference.Wac {
						count.WAC++
					}

					if conference.Wos {
						count.Wos++
					}

					if conference.Scopus {
						count.Scopus++
					}

					if conference.Rinc {
						count.Rinc++
					}
				}

				return count.String()
			}(),
			Qualified: func() string {
				if student.AttestationMark > 2 {
					return "аттестован"
				}

				return "НЕ аттестован"
			}(),
			ActualYear: student.Student.Years / 2,
		}

		reports = append(reports, report)
	}

	jsonReports, err := json.Marshal(reports)
	fmt.Println("json: " + string(jsonReports))
	// Command to run the Python script
	cmd := exec.Command("venv/bin/python", "report_generate.py", "-js", string(jsonReports), "-ds", "./reports/report_1/output.xlsx")

	// Get the output from the Python script
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("OUTPUT: " + string(output))
		return nil, errors.Wrap(err, "calling python script")
	}
	fmt.Println("OUTPUT: " + string(output))

	return reports, nil
}

func (s *Service) GenerateReportTwo(ctx context.Context) ([]models.ReportTwo, error) {
	var studentsInfo []models.StudentInfoForReportTwo
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		studentIDs, err := s.clientRepo.GetAllStudentIDs(ctx, tx)
		if err != nil {
			return err
		}

		info, err := s.clientRepo.GetDataForReportTwo(ctx, tx, studentIDs)
		if err != nil {
			return err
		}

		studentsInfo = info

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "GenerateReportOne()")
	}

	reports := make([]models.ReportTwo, 0, len(studentsInfo))

	for _, student := range studentsInfo {
		report := models.ReportTwo{
			FullName:          student.Student.FullName,
			SupervisorName:    student.SupervisorName,
			DissertationTitle: student.DissertationTitle.Title,
			SupervisorMark:    student.SupervisorMark,
			Progressiveness:   student.Progressiveness,
			TeachingLoad: lo.Reduce(
				student.ClassroomLoad,
				func(agg int32, item model.ClassroomLoad, _ int) int32 {
					return agg + item.Hours
				},
				0,
			),
			ScientificWorkCount: func() string {
				count := models.ScientificWorkCount{
					WAC:    0,
					Scopus: 0,
					Rinc:   0,
					Wos:    0,
				}

				for _, publication := range student.Publications {
					if publication.Wac {
						count.WAC++
					}

					if publication.Wos {
						count.Wos++
					}

					if publication.Scopus {
						count.Scopus++
					}

					if publication.Rinc {
						count.Rinc++
					}
				}
				return count.String()
			}(),
			ActualYear: student.Student.Years / 2,
			Comment:    "",
		}

		reports = append(reports, report)
	}

	jsonReports, err := json.Marshal(reports)
	fmt.Println("json: " + string(jsonReports))
	// Command to run the Python script
	cmd := exec.Command("venv/bin/python", "report_generate.py", "-js", string(jsonReports), "-ds", "./reports/report_2/output.xlsx")

	// Get the output from the Python script
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, "calling python script")
	}
	fmt.Println("OUTPUT: " + string(output))

	return reports, nil
}
