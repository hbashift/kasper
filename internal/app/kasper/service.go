package kasper

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type StudentHandler interface {
	GetDissertationPage(ctx *gin.Context)
	UpsertSemesterProgress(ctx *gin.Context)
	GetScientificWorks(ctx *gin.Context)
	InsertScientificWorks(ctx *gin.Context)
	UpdateScientificWorks(ctx *gin.Context)
	DeleteScientificWork(ctx *gin.Context)
	GetTeachingLoad(ctx *gin.Context)
	UpsertTeachingLoad(ctx *gin.Context)
	DeleteTeachingLoad(ctx *gin.Context)
	UploadDissertation(ctx *gin.Context)
	DownloadDissertation(ctx *gin.Context)
	GetSupervisors(ctx *gin.Context)
	FirstRegistry(ctx *gin.Context)
}

type SupervisorHandler interface {
	GetListOfStudents(ctx *gin.Context)
	GetStudentsDissertationPage(ctx *gin.Context)
	DownloadDissertation(ctx *gin.Context)
	SetStatus(ctx *gin.Context)
	UpdateFeedback(ctx *gin.Context)
	GetScientificWorks(ctx *gin.Context)
	GetTeachingLoad(ctx *gin.Context)
}

type AuthorizationHandler interface {
	Authorize(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
	HealthCheck(ctx *gin.Context)
}

type AdministratorHandler interface {
	SetAcademicLeave(ctx *gin.Context)
	UpdateStudentCommonInfo(ctx *gin.Context)
	ChangeSupervisor(ctx *gin.Context)
	GetPairs(ctx *gin.Context)
	GetScientificWorks(ctx *gin.Context)
	GetStudentsDissertationPage(ctx *gin.Context)
	GetTeachingLoad(ctx *gin.Context)
}

func InitRoutes(student StudentHandler,
	supervisor SupervisorHandler,
	authorization AuthorizationHandler,
	adminHandler AdministratorHandler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding", "StudentID"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Content-Disposition"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/students/dissertation/:id", student.GetDissertationPage)
	router.POST("/students/dissertation/progress/:id", student.UpsertSemesterProgress)
	router.GET("/students/scientific_works/:id", student.GetScientificWorks)
	router.POST("/students/scientific_works/:id", student.InsertScientificWorks)
	router.PATCH("/students/scientific_works/:id", student.UpdateScientificWorks)
	router.DELETE("/students/scientific_works/:id", student.DeleteScientificWork)
	router.GET("/students/teaching_load/:id", student.GetTeachingLoad)
	router.POST("/students/teaching_load/:id", student.UpsertTeachingLoad)
	router.DELETE("/students/teaching_load/:id", student.DeleteTeachingLoad)
	router.POST("/students/dissertation/file/:id", student.UploadDissertation)
	router.PUT("/students/dissertation/file/:id", student.DownloadDissertation)
	router.GET("/students/supervisors/:id", student.GetSupervisors)
	router.POST("/students/registration/:id", student.FirstRegistry)

	router.GET("/supervisors/list_of_students/:id", supervisor.GetListOfStudents)
	router.PUT("/supervisors/student/:id", supervisor.GetStudentsDissertationPage)
	router.PUT("/supervisor/students/dissertation/:id", supervisor.DownloadDissertation)
	router.POST("/supervisor/students/set_status/:id", supervisor.SetStatus)
	router.POST("/supervisor/students/feedback/:id", supervisor.UpdateFeedback)
	router.PUT("/supervisor/students/scientific_works/:id", supervisor.GetScientificWorks)
	router.PUT("/supervisor/students/teaching_load/:id", supervisor.GetTeachingLoad)

	router.POST("/authorization/authorize", authorization.Authorize)
	router.POST("/authorization/change_password/:id", authorization.ChangePassword)
	router.GET("/authorization/check/:id", authorization.HealthCheck)

	router.POST("/admin/students/set_academic/:id", adminHandler.SetAcademicLeave)
	router.POST("/admin/students/common_info/:id", adminHandler.UpdateStudentCommonInfo)
	router.POST("/admin/pairs/:id", adminHandler.ChangeSupervisor)
	router.GET("/admin/pairs/:id", adminHandler.GetPairs)
	router.GET("/admin/students/scientific/:id", adminHandler.GetScientificWorks)
	router.GET("/admin/students/load/:id", adminHandler.GetTeachingLoad)
	router.GET("/admin/students/dissertation/:id", adminHandler.GetStudentsDissertationPage)

	return router
}
