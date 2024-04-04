package email

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
)

type (
	UsersRepository interface {
		GetUserByKasperIDTx(ctx context.Context, tx pgx.Tx, kasperID uuid.UUID) (model.Users, error)
	}

	ClientRepository interface {
		GetStudentTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) (model.Students, error)
		GetStudentsActualSupervisorTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) (models.Supervisor, error)
		GetSupervisorTx(ctx context.Context, tx pgx.Tx, supervisorID uuid.UUID) (models.Supervisor, error)
	}
)

type Service struct {
	sender     string
	password   string
	host       string
	db         *pgxpool.Pool
	userRepo   UsersRepository
	clientRepo ClientRepository
}

func NewService(sender string, password string, host string, db *pgxpool.Pool, userRepo UsersRepository, clientRepo ClientRepository) *Service {
	return &Service{sender: sender, password: password, host: host, db: db, userRepo: userRepo, clientRepo: clientRepo}
}

type SupervisorData struct {
	SupervisorName string
	StudentName    string
	Status         string
	Type           string
}

type StudentData struct {
	SupervisorName string
	StudentName    string
	Type           string
}

func (s *Service) SendSupervisorEmail(ctx context.Context, studentID, supervisorID uuid.UUID, templatePath, tt, status string) error {
	var data SupervisorData
	var (
		receiver string
	)

	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		uStud, err := s.userRepo.GetUserByKasperIDTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		stud, err := s.clientRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		sup, err := s.clientRepo.GetSupervisorTx(ctx, tx, supervisorID)
		if err != nil {
			return err
		}

		data = SupervisorData{
			SupervisorName: sup.FullName,
			StudentName:    stud.FullName,
			Status:         status,
			Type:           tt,
		}

		receiver = uStud.Email

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "SendSupervisorEmail()")
	}

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return errors.Wrap(err, "parsing template")
	}

	err = t.Execute(&body, data)

	m := gomail.NewMessage()
	m.SetHeader("From", s.sender)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", fmt.Sprintf("Уведомление от Научного руководителя %s", data.SupervisorName))
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(s.host, 587, s.sender, s.password)

	if err = d.DialAndSend(m); err != nil {
		return errors.Wrap(err, "sending email")
	}

	return nil
}

func (s *Service) SendStudentEmail(ctx context.Context, studentID uuid.UUID, templatePath, tt string) error {
	var data StudentData
	var (
		receiver string
	)

	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		uStud, err := s.userRepo.GetUserByKasperIDTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		stud, err := s.clientRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		sup, err := s.clientRepo.GetStudentsActualSupervisorTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		data = StudentData{
			SupervisorName: sup.FullName,
			StudentName:    stud.FullName,
			Type:           tt,
		}

		receiver = uStud.Email

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "SendSupervisorEmail()")
	}
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return errors.Wrap(err, "parsing template")
	}

	err = t.Execute(&body, data)

	m := gomail.NewMessage()
	m.SetHeader("From", s.sender)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", fmt.Sprintf("Уведомление от Студента %s", data.StudentName))
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(s.host, 587, s.sender, s.password)

	if err = d.DialAndSend(m); err != nil {
		return errors.Wrap(err, "sending email")
	}

	return nil
}
