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
		AllowOrigins:     []string{"http://localhost:3000"}, // Permite solicitudes desde el frontend
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
	//r.GET("/users", GetUsers(db))
	//r.GET("/users/:id", GetUserByID(db))
	//r.PUT("/users/:id", UpdateUser(db))
	//r.DELETE("/users/:id", DeleteUser(db))

	// Tareas
	r.POST("/tasks", CreateTask(db))
	r.GET("/tasks", GetTasks(db))
	r.GET("/tasks/:id", GetTaskByID(db))
	r.PUT("/tasks/:id", UpdateTask(db))
	//r.DELETE("/tasks/:id", DeleteTask(db))
	//r.GET("/users/:id/tasks", GetTasksByUser(db))

	// Comentarios
	r.POST("/tasks/:id/comments", CreateComment(db))
	//r.GET("/tasks/:id/comments", GetCommentsByTask(db))
	//r.DELETE("/comments/:id", DeleteComment(db))

	return r

}
