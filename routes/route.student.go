package route

import (
	createStudent "github.com/fstar-dev/restapi-gin/controllers/student-controllers/create"
	deleteStudent "github.com/fstar-dev/restapi-gin/controllers/student-controllers/delete"
	resultStudent "github.com/fstar-dev/restapi-gin/controllers/student-controllers/result"
	resultsStudent "github.com/fstar-dev/restapi-gin/controllers/student-controllers/results"
	updateStudent "github.com/fstar-dev/restapi-gin/controllers/student-controllers/update"
	handlerCreateStudent "github.com/fstar-dev/restapi-gin/handlers/student-handlers/create"
	handlerDeleteStudent "github.com/fstar-dev/restapi-gin/handlers/student-handlers/delete"
	handlerResultStudent "github.com/fstar-dev/restapi-gin/handlers/student-handlers/result"
	handlerResultsStudent "github.com/fstar-dev/restapi-gin/handlers/student-handlers/results"
	handlerUpdateStudent "github.com/fstar-dev/restapi-gin/handlers/student-handlers/update"
	middleware "github.com/fstar-dev/restapi-gin/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitStudentRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Student
	*/
	createStudentRepository := createStudent.NewRepositoryCreate(db)
	createStudentService := createStudent.NewServiceCreate(createStudentRepository)
	createStudentHandler := handlerCreateStudent.NewHandlerCreateStudent(createStudentService)

	resultsStudentRepository := resultsStudent.NewRepositoryResults(db)
	resultsStudentService := resultsStudent.NewServiceResults(resultsStudentRepository)
	resultsStudentHandler := handlerResultsStudent.NewHandlerResultsStudent(resultsStudentService)

	resultStudentRepository := resultStudent.NewRepositoryResult(db)
	resultStudentService := resultStudent.NewServiceResult(resultStudentRepository)
	resultStudentHandler := handlerResultStudent.NewHandlerResultStudent(resultStudentService)

	deleteStudentRepository := deleteStudent.NewRepositoryDelete(db)
	deleteStudentService := deleteStudent.NewServiceDelete(deleteStudentRepository)
	deleteStudentHandler := handlerDeleteStudent.NewHandlerDeleteStudent(deleteStudentService)

	updateStudentRepository := updateStudent.NewRepositoryUpdate(db)
	updateStudentService := updateStudent.NewServiceUpdate(updateStudentRepository)
	updateStudentHandler := handlerUpdateStudent.NewHandlerUpdateStudent(updateStudentService)

	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/student", createStudentHandler.CreateStudentHandler)
	groupRoute.GET("/student", resultsStudentHandler.ResultsStudentHandler)
	groupRoute.GET("/student/:id", resultStudentHandler.ResultStudentHandler)
	groupRoute.DELETE("/student/:id", deleteStudentHandler.DeleteStudentHandler)
	groupRoute.PUT("/student/:id", updateStudentHandler.UpdateStudentHandler)
}
