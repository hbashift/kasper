package main

import (
	"fmt"

	"uir_draft/internal/generated/new_kasper/new_uir/public/table"

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
	stmt, args := table.ScientificWorksStatus.
		SELECT(
			table.ScientificWorksStatus.WorksID,
			table.ScientificWorksStatus.Semester,
			table.ScientificWorksStatus.StudentID,
			table.ScientificWorksStatus.Status.AS("scientific_works.approval_status"),
			table.ScientificWorksStatus.UpdatedAt,
			table.ScientificWorksStatus.AcceptedAt,
			table.Publications.AllColumns.Except(table.Publications.WorksID),
			table.Conferences.AllColumns.Except(table.Conferences.WorksID),
			table.ResearchProjects.AllColumns.Except(table.ResearchProjects.WorksID),
		).
		FROM(table.ScientificWorksStatus.
			INNER_JOIN(table.Publications, table.ScientificWorksStatus.WorksID.EQ(table.Publications.WorksID)).
			INNER_JOIN(table.Conferences, table.ScientificWorksStatus.WorksID.EQ(table.Conferences.WorksID)).
			INNER_JOIN(table.ResearchProjects, table.ScientificWorksStatus.WorksID.EQ(table.ResearchProjects.WorksID)),
		).
		WHERE(table.ScientificWorksStatus.StudentID.EQ(postgres.UUID(uuid.New()))).
		Sql()

	fmt.Println(stmt, args)

	//err := initConfig()
	//ctx := context.Background()
	//
	//db, err := configs.InitPostgresDB(ctx, configs.Config{
	//	Host:     viper.GetString("db.host"),
	//	Port:     viper.GetString("db.port"),
	//	Username: viper.GetString("db.username"),
	//	Password: viper.GetString("db.password"),
	//	DBName:   viper.GetString("db.dbname"),
	//	SSLMode:  viper.GetString("db.sslmode"),
	//})
	//
	//if err != nil {
	//	panic(err)
	//}
	//
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
