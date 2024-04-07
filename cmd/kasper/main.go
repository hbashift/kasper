package main

import (
	"context"

	"uir_draft/internal/app/kasper"
	"uir_draft/internal/handlers/administator_handler"
	"uir_draft/internal/handlers/authorization_handler"
	"uir_draft/internal/handlers/student_handler"
	"uir_draft/internal/handlers/supervisor_handler"
	"uir_draft/internal/pkg/configs"
	"uir_draft/internal/pkg/repository"
	"uir_draft/internal/pkg/service/admin"
	"uir_draft/internal/pkg/service/authentication"
	"uir_draft/internal/pkg/service/email"
	"uir_draft/internal/pkg/service/enum"
	"uir_draft/internal/pkg/service/student"
	"uir_draft/internal/pkg/service/supervisor"

	"github.com/spf13/viper"
)

//	@title		Сервис Kasper
//	@version	0.1.2
// description Серверная часть Системы учета деятельности аспирантов

//	@host		localhost:8080

func main() {
	err := initConfig()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	db, err := configs.InitPostgresDB(ctx, configs.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		panic(err)
	}

	clientRepo := repository.NewClientRepository()
	dissertationRepo := repository.NewDissertationRepository()
	marksRepo := repository.NewMarksRepository()
	scienceRepo := repository.NewScientificRepository()
	loadRepo := repository.NewTeachingLoadRepository()
	enumRepo := repository.NewEnumRepository()
	usersRepo := repository.NewUsersRepository()
	tokenRepo := repository.NewTokenRepository()

	studentService := student.NewService(dissertationRepo, loadRepo, scienceRepo, marksRepo, clientRepo, tokenRepo, usersRepo, db)
	adminService := admin.NewService(dissertationRepo, loadRepo, scienceRepo, marksRepo, clientRepo, tokenRepo, usersRepo, db)
	supervisorService := supervisor.NewService(dissertationRepo, tokenRepo, usersRepo, clientRepo, db)
	authenticationService := authentication.NewService(tokenRepo, usersRepo, db)
	emailService := email.NewService("SENDER", "PASSWORD", "smtp.gmail.com", db, usersRepo, clientRepo)
	enumService := enum.NewService(enumRepo, db)

	studentHandler := student_handler.NewHandler(studentService, studentService, studentService, studentService, authenticationService, emailService, enumService, adminService)
	supervisorHandler := supervisor_handler.NewHandler(studentService, studentService, studentService, authenticationService, supervisorService, studentService, emailService)
	adminHandler := administator_handler.NewHandler(adminService, authenticationService, enumService)
	authenticationHandler := authorization_handler.NewHandler(authenticationService, studentService)

	server := kasper.NewHTTPServer(studentHandler, supervisorHandler, adminHandler, authenticationHandler)
	r := server.InitRouter()

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
