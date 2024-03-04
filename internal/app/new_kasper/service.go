package new_kasper

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type StudentHandler interface {
	AllToReview(ctx *gin.Context)
	GetStudentStatus(ctx *gin.Context)

	GetDissertationPage(ctx *gin.Context)
	UpsertSemesterProgress(ctx *gin.Context)
	DownloadDissertation(ctx *gin.Context)
	UploadDissertation(ctx *gin.Context)

	GetTeachingLoad(ctx *gin.Context)
	UpsertAdditionalLoads(ctx *gin.Context)
	UpsertClassroomLoads(ctx *gin.Context)
	UpsertIndividualLoads(ctx *gin.Context)
	DeleteAdditionalLoads(ctx *gin.Context)
	DeleteClassroomLoads(ctx *gin.Context)
	DeleteIndividualLoads(ctx *gin.Context)

	GetScientificWorks(ctx *gin.Context)
	UpsertResearchProjects(ctx *gin.Context)
	UpsertPublications(ctx *gin.Context)
	UpsertConferences(ctx *gin.Context)
	DeleteProjects(ctx *gin.Context)
	DeletePublications(ctx *gin.Context)
	DeleteConferences(ctx *gin.Context)
}

type SupervisorHandler interface {
	GetStudentsList(ctx *gin.Context)

	AllToStatus(ctx *gin.Context)

	GetDissertationPage(ctx *gin.Context)
	UpsertFeedback(ctx *gin.Context)

	GetTeachingLoad(ctx *gin.Context)
	GetScientificWorks(ctx *gin.Context)
}

type HTTPServer struct {
	student    StudentHandler
	supervisor SupervisorHandler
}

func NewHTTPServer(studentHandler StudentHandler, supervisorHandler SupervisorHandler) *HTTPServer {
	return &HTTPServer{
		student:    studentHandler,
		supervisor: supervisorHandler,
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

	r.GET("/students/load/:token", h.student.GetTeachingLoad)
	r.POST("/students/load/classroom/:token", h.student.UpsertClassroomLoads)
	r.DELETE("/students/load/classroom/:token", h.student.DeleteClassroomLoads)
	r.POST("/students/load/individual/:token", h.student.UpsertIndividualLoads)
	r.DELETE("/students/load/individual/:token", h.student.DeleteIndividualLoads)
	r.POST("/students/load/additional/:token", h.student.UpsertAdditionalLoads)
	r.DELETE("/students/load/additional/:token", h.student.DeleteAdditionalLoads)

	r.GET("/students/works/:token", h.student.GetScientificWorks)
	r.POST("/students/works/publications/:token", h.student.UpsertPublications)
	r.DELETE("/students/works/publications/:token", h.student.DeletePublications)
	r.POST("/students/works/conferences/:token", h.student.UpsertConferences)
	r.DELETE("/students/works/conferences/:token", h.student.DeleteConferences)
	r.POST("/students/works/projects/:token", h.student.UpsertResearchProjects)
	r.DELETE("/students/works/projects/:token", h.student.DeleteProjects)

	// SupervisorHandler init
	r.PUT("/supervisors/student/list/:token", h.supervisor.GetStudentsList)

	r.POST("/supervisors/student/review/:token", h.supervisor.AllToStatus)

	r.PUT("/supervisors/student/dissertation/:token", h.supervisor.GetDissertationPage)
	r.POST("/supervisors/student/dissertation/feedback/:token", h.supervisor.UpsertFeedback)

	r.PUT("/supervisors/student/load/:token", h.supervisor.GetTeachingLoad)
	r.PUT("/supervisors/student/works/:token", h.supervisor.GetScientificWorks)

	return r
}
