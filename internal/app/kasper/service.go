package kasper

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type StudentHandler interface {
	GetDissertation(ctx *gin.Context)
	UpsertSemesterProgress(ctx *gin.Context)
	GetScientificWorks(ctx *gin.Context)
	InsertScientificWorks(ctx *gin.Context)
	UpdateScientificWorks(ctx *gin.Context)
	DeleteScientificWork(ctx *gin.Context)
	GetTeachingLoad(ctx *gin.Context)
	InsertTeachingLoad(ctx *gin.Context)
	UpdateTeachingLoad(ctx *gin.Context)
}

func InitRoutes(student StudentHandler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/students/dissertation/:id", student.GetDissertation)
	router.POST("/students/dissertation/progress/:id", student.UpsertSemesterProgress)
	router.GET("/students/scientific_works/:id", student.GetScientificWorks)
	router.POST("/students/scientific_works/:id", student.InsertScientificWorks)
	router.PATCH("/students/scientific_works/:id", student.UpdateScientificWorks)
	router.DELETE("/students/scientific_works/:id", student.DeleteScientificWork)
	router.GET("/students/teaching_load/:id", student.GetTeachingLoad)
	router.POST("/students/teaching_load/:id", student.InsertTeachingLoad)
	router.PATCH("/students/teaching_load/:id", student.UpdateTeachingLoad)

	return router
}
