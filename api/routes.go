package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	_ "teca_notifications/docs"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Configura CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Permite solicitudes desde el frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Usuarios
	r.POST("/users", RegisterUser(db))
	r.POST("/users/login", LoginUser(db))
	r.GET("/getAllUsers", GetUsers(db))
	//r.GET("/users/:id", GetUserByID(db))
	//r.PUT("/users/:id", UpdateUser(db))
	//r.DELETE("/users/:id", DeleteUser(db))

	// Tareas
	r.POST("/createTask", CreateTask(db))
	r.GET("/getAllTasks", GetTasks(db))
	r.GET("/getTasksById/:id", GetTaskByID(db))
	r.PUT("/updateTasks/:id", UpdateTask(db))
	//r.DELETE("/tasks/:id", DeleteTask(db))
	//r.GET("/users/:id/tasks", GetTasksByUser(db))

	// Comentarios
	r.POST("/tasks/createComments", CreateComment(db))
	r.GET("/tasks/:id/comments", GetCommentsByTask(db))
	//r.DELETE("/comments/:id", DeleteComment(db))

	// Actividades
	r.POST("/tasks/createActivity", CreateActivity(db))
	r.GET("/tasks/:id/activities", GetActivitiesByTask(db))

	return r

}
