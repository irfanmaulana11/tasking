package server

import (
	"be-tasking/app/handler"
	"be-tasking/app/service"

	"github.com/gin-gonic/gin"
)

var (
	BaseURL = "/api"
)

func InitRoutes(r *gin.Engine, hcs service.HealthCheckServiceInterface, aus service.AuthServiceInterface, tas service.TaskServiceInterface) {

	// init api handler
	handler := handler.NewApiHandler(hcs, aus, tas)

	// route list
	route := r.Group(BaseURL)
	route.GET("/health-check", handler.Check)

	auth := route.Group("/auth")
	auth.POST("/login", handler.UserLogin)
	auth.POST("/register", handler.UserRegister)

	task := route.Group("/tasks")
	task.GET("/", AuthMiddleware(), handler.GetTaskList)
	task.POST("/", AuthMiddleware(), PelaksanaOnlyMiddleware(), handler.CreateTask)
	task.PUT("/:id", AuthMiddleware(), PelaksanaOnlyMiddleware(), handler.UpdateTask)
	task.PATCH("/:id/progress", AuthMiddleware(), handler.TaskInprogress)
	task.PATCH("/:id/progress/overide", AuthMiddleware(), LeaderOnlyMiddleware(), handler.TaskOveride)
	task.PATCH("/:id/revise", AuthMiddleware(), LeaderOnlyMiddleware(), handler.TaskRevise)
	task.PATCH("/:id/approve", AuthMiddleware(), LeaderOnlyMiddleware(), handler.TaskApproved)
	task.PATCH("/:id/complete", AuthMiddleware(), LeaderAndPelaksanaOnlyMiddleware(), handler.TaskCompleted)

}
