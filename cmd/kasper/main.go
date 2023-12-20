package main

import (
	"context"

	"uir_draft/internal/app/kasper"
	"uir_draft/internal/handler/authorization_handler"
	"uir_draft/internal/handler/student_handler"
	"uir_draft/internal/pkg/repositories"
	"uir_draft/internal/pkg/service/authorization"
	"uir_draft/internal/pkg/service/student"

	"github.com/spf13/viper"
)

func main() {
	err := initConfig()
	ctx := context.Background()

	db, err := repositories.InitPostgresDB(ctx, repositories.Config{
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

	studRepo := repositories.NewStudentRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)
	dRepo := repositories.NewDissertationRepository(db)
	semesterRepo := repositories.NewSemesterRepository(db)
	scientificRepo := repositories.NewScientificWork(db)
	loadRepo := repositories.NewTeachingLoadRepository()
	clientRepo := repositories.NewClientUserRepository()

	studService := student.NewService(studRepo, tokenRepo, dRepo, semesterRepo, scientificRepo, loadRepo, db)
	studHandler := student_handler.NewStudentHandler(studService)

	authorizeService := authorization.NewService(clientRepo, tokenRepo, db)
	authorizeHandler := authorization_handler.NewAuthorizationHandler(authorizeService)

	server := kasper.InitRoutes(studHandler, authorizeHandler)

	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
