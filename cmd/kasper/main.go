package main

import (
	"context"
	"fmt"

	"uir_draft/internal/app/new_kasper"
	"uir_draft/internal/generated/new_kasper/new_uir/public/table"
	"uir_draft/internal/handlers/administator_handler"
	"uir_draft/internal/handlers/student_handler"
	"uir_draft/internal/handlers/supervisor_handler"
	"uir_draft/internal/pkg/configs"
	"uir_draft/internal/pkg/new_repo"
	"uir_draft/internal/pkg/new_service/admin"
	"uir_draft/internal/pkg/new_service/authentication"
	"uir_draft/internal/pkg/new_service/student"
	"uir_draft/internal/pkg/new_service/supervisor"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

//	@title		Сервис Kasper
//	@version	0.1.2
// description Серверная часть Системы учета деятельности аспирантов

//	@host		localhost:8080
//	@BasePath	/api/v1

func main() {
	stmt, _ := table.Users.
		DELETE().
		WHERE(table.Users.UserID.EQ(postgres.UUID(uuid.New()))).
		Sql()

	fmt.Println(stmt, uuid.New())

	err := initConfig()
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

	clientRepo := new_repo.NewClientRepository()
	dissertationRepo := new_repo.NewDissertationRepository()
	marksRepo := new_repo.NewMarksRepository()
	scienceRepo := new_repo.NewScientificRepository()
	loadRepo := new_repo.NewTeachingLoadRepository()
	//enumRepo := new_repo.NewEnumRepository()
	usersRepo := new_repo.NewUsersRepository()
	tokenRepo := new_repo.NewTokenRepository()

	studentService := student.NewService(dissertationRepo, loadRepo, scienceRepo, marksRepo, clientRepo, tokenRepo, usersRepo, db)
	adminService := admin.NewService(dissertationRepo, loadRepo, scienceRepo, marksRepo, clientRepo, tokenRepo, usersRepo, db)
	supervisorService := supervisor.NewService(dissertationRepo, tokenRepo, usersRepo, clientRepo, db)
	authenticationService := authentication.NewService(tokenRepo, usersRepo, db)

	studentHandler := student_handler.NewHandler(studentService, studentService, studentService, studentService, authenticationService)
	supervisorHandler := supervisor_handler.NewHandler(studentService, studentService, studentService, authenticationService, supervisorService)
	adminHandler := administator_handler.NewHandler(adminService, authenticationService)

	server := new_kasper.NewHTTPServer(studentHandler, supervisorHandler, adminHandler)
	r := server.InitRouter()

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}

	//studRepo := repositories.NewStudentRepository(db)
	//tokenRepo := repositories.NewTokenRepository(db)
	//dRepo := repositories.NewDissertationRepository(db)
	//semesterRepo := repositories.NewSemesterRepository(db)
	//scientificRepo := repositories.NewScientificWork(db)
	//loadRepo := repositories.NewTeachingLoadRepository()
	//clientRepo := repositories.NewClientUserRepository()
	//supRepo := repositories.NewSupervisorRepository()
	//
	//studService := student.NewService(studRepo, tokenRepo, dRepo, semesterRepo, scientificRepo, loadRepo, supRepo, clientRepo, db)
	//studHandler := student_handler.NewStudentHandler(studService)
	//
	//authorizeService := authorization.NewService(clientRepo, tokenRepo, db)
	//authorizeHandler := authorization_handler.NewAuthorizationHandler(authorizeService)
	//
	//supService := supervisor.NewService(studRepo, tokenRepo, semesterRepo, dRepo, scientificRepo, loadRepo, db)
	//supervisorHandler := supervisor_handler.NewSupervisorHandler(supService)
	//
	//adminService := admin.NewService(studRepo, tokenRepo, semesterRepo, dRepo, scientificRepo, loadRepo, supRepo, db)
	//adminHandler := admin_handler.NewAdministratorHandler(adminService)
	//
	//server := kasper.InitRoutes(studHandler, supervisorHandler, authorizeHandler, adminHandler)
	//
	//err = server.Run(":8080")
	//if err != nil {
	//	panic(err)
	//}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
