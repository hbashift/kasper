package kasper

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type (
	StudentHandler interface {
		AllToReview(ctx *gin.Context)
		GetStudentStatus(ctx *gin.Context)

		GetDissertationPage(ctx *gin.Context)
		UpsertSemesterProgress(ctx *gin.Context)
		DownloadDissertation(ctx *gin.Context)
		UploadDissertation(ctx *gin.Context)
		DissertationTitleToReview(ctx *gin.Context)
		UpsertDissertationTitle(ctx *gin.Context)
		DissertationToReview(ctx *gin.Context)

		GetTeachingLoad(ctx *gin.Context)
		UpsertAdditionalLoads(ctx *gin.Context)
		UpsertClassroomLoads(ctx *gin.Context)
		UpsertIndividualLoads(ctx *gin.Context)
		DeleteAdditionalLoads(ctx *gin.Context)
		DeleteClassroomLoads(ctx *gin.Context)
		DeleteIndividualLoads(ctx *gin.Context)
		TeachingLoadToReview(ctx *gin.Context)

		GetScientificWorks(ctx *gin.Context)
		UpsertResearchProjects(ctx *gin.Context)
		UpsertPublications(ctx *gin.Context)
		UpsertConferences(ctx *gin.Context)
		DeleteProjects(ctx *gin.Context)
		DeletePublications(ctx *gin.Context)
		DeleteConferences(ctx *gin.Context)
		ScientificWorksToReview(ctx *gin.Context)

		GetSpecializations(ctx *gin.Context)
		GetGroups(ctx *gin.Context)
	}

	SupervisorHandler interface {
		GetStudentsList(ctx *gin.Context)

		AllToStatus(ctx *gin.Context)

		GetDissertationPage(ctx *gin.Context)
		UpsertFeedback(ctx *gin.Context)
		DownloadDissertation(ctx *gin.Context)

		GetTeachingLoad(ctx *gin.Context)
		GetScientificWorks(ctx *gin.Context)

		GetStudentStatus(ctx *gin.Context)
	}

	AdministratorHandler interface {
		ChangeSupervisor(ctx *gin.Context)
		GetPairs(ctx *gin.Context)
		SetStudentStudyingStatus(ctx *gin.Context)
		GetSupervisors(ctx *gin.Context)

		GetSpecializations(ctx *gin.Context)
		GetGroups(ctx *gin.Context)

		AddSpecializations(ctx *gin.Context)
		AddGroups(ctx *gin.Context)
	}

	AuthenticationHandler interface {
		Authorize(ctx *gin.Context)
		FirstStudentRegistry(ctx *gin.Context)
	}
)

type HTTPServer struct {
	student        StudentHandler
	supervisor     SupervisorHandler
	administrator  AdministratorHandler
	authentication AuthenticationHandler
}

func NewHTTPServer(studentHandler StudentHandler, supervisorHandler SupervisorHandler, adminHandler AdministratorHandler, authenticationHandler AuthenticationHandler) *HTTPServer {
	return &HTTPServer{
		student:        studentHandler,
		supervisor:     supervisorHandler,
		administrator:  adminHandler,
		authentication: authenticationHandler,
	}
}

func (h *HTTPServer) InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding", "StudentID"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Content-Disposition"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// StudentHandlers init
	r.POST("/students/review/:token", h.student.AllToReview)
	r.GET("/students/info/:token", h.student.GetStudentStatus)

	r.GET("/students/dissertation/:token", h.student.GetDissertationPage)
	r.POST("/students/dissertation/progress/:token", h.student.UpsertSemesterProgress)
	r.PUT("/students/dissertation/file/:token", h.student.DownloadDissertation)
	r.POST("/students/dissertation/file/:token", h.student.UploadDissertation)
	r.POST("/students/dissertation_title/review/:token", h.student.DissertationTitleToReview)
	r.POST("/students/dissertation_title/:token", h.student.UpsertDissertationTitle)
	r.POST("/students/dissertation/review/:token", h.student.DissertationToReview)

	r.GET("/students/load/:token", h.student.GetTeachingLoad)
	r.POST("/students/load/classroom/:token", h.student.UpsertClassroomLoads)
	r.PUT("/students/load/classroom/:token", h.student.DeleteClassroomLoads)
	r.POST("/students/load/individual/:token", h.student.UpsertIndividualLoads)
	r.PUT("/students/load/individual/:token", h.student.DeleteIndividualLoads)
	r.POST("/students/load/additional/:token", h.student.UpsertAdditionalLoads)
	r.PUT("/students/load/additional/:token", h.student.DeleteAdditionalLoads)
	r.POST("/student/load/review/:token", h.student.TeachingLoadToReview)

	r.GET("/students/works/:token", h.student.GetScientificWorks)
	r.POST("/students/works/publications/:token", h.student.UpsertPublications)
	r.PUT("/students/works/publications/:token", h.student.DeletePublications)
	r.POST("/students/works/conferences/:token", h.student.UpsertConferences)
	r.PUT("/students/works/conferences/:token", h.student.DeleteConferences)
	r.POST("/students/works/projects/:token", h.student.UpsertResearchProjects)
	r.PUT("/students/works/projects/:token", h.student.DeleteProjects)
	r.POST("/students/works/review/:token", h.student.ScientificWorksToReview)

	r.GET("/student/enum/specializations/:token", h.student.GetSpecializations)
	r.GET("/student/enum/groups/:token", h.student.GetGroups)

	// SupervisorHandler init
	r.PUT("/supervisors/student/list/:token", h.supervisor.GetStudentsList)

	r.POST("/supervisors/student/review/:token", h.supervisor.AllToStatus)

	r.PUT("/supervisors/student/dissertation/:token", h.supervisor.GetDissertationPage)
	r.POST("/supervisors/student/dissertation/feedback/:token", h.supervisor.UpsertFeedback)
	r.PUT("/supervisors/student/dissertation/file/:token", h.supervisor.DownloadDissertation)

	r.PUT("/supervisors/student/load/:token", h.supervisor.GetTeachingLoad)
	r.PUT("/supervisors/student/works/:token", h.supervisor.GetScientificWorks)

	r.PUT("/supervisors/student/info/:token", h.supervisor.GetStudentStatus)

	// AdministratorHandler init
	r.POST("/administrator/student/change/:token", h.administrator.ChangeSupervisor)

	r.GET("/administrator/pairs/:token", h.administrator.GetPairs)
	r.POST("/administrator/student/status/:token", h.administrator.SetStudentStudyingStatus)

	r.GET("/administrator/supervisors/list/:token", h.administrator.GetSupervisors)

	r.GET("/administrator/enum/specializations/:token", h.administrator.GetSpecializations)
	r.GET("/administrator/enum/groups/:token", h.administrator.GetGroups)

	r.POST("/administrator/enum/specializations/:token", h.administrator.AddSpecializations)
	r.POST("/administrator/enum/groups/:token", h.administrator.AddGroups)

	// AuthenticationHandler init
	r.POST("/authorize", h.authentication.Authorize)

	r.POST("/student/registry/:token", h.authentication.FirstStudentRegistry)

	return r
}
