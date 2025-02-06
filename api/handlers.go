package api

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"teca_notifications/utils" // Importa el paquete utils
	"time"

	"github.com/gin-gonic/gin"
)

// GetTasks godoc
// @Summary Obtener todas las tareas
// @Description Obtiene una lista de todas las tareas
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} Task
// @Failure 500 {object} map[string]string
// @Router /getAllTasks [get]
func GetTasks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tasks []Task
		if err := db.Preload("Comments").Find(&tasks).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, tasks)
	}
}

// GetTaskByID godoc
// @Summary Obtener una tarea por ID
// @Description Obtiene una tarea específica por su ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID de la tarea"
// @Success 200 {object} Task
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /getTasksById/{id} [get]
func GetTaskByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var task Task
		if err := db.Preload("Comments").First(&task, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
			return
		}

		c.JSON(http.StatusOK, task)
	}
}

// RegisterUser godoc
// @Summary Registrar un nuevo usuario
// @Description Crea un nuevo usuario en el sistema
// @Tags users
// @Accept json
// @Produce json
// @Param user body RegisterUserRequest true "Datos del usuario"
// @Success 201 {object} User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func RegisterUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request RegisterUserRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Hashear la contraseña
		hashedPassword, err := utils.HashPassword(request.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hashear la contraseña"})
			return
		}

		// Mapear el DTO al modelo User
		user := User{
			Name:         request.Name,
			Email:        request.Email,
			PasswordHash: hashedPassword,
		}

		// Crear el usuario en la base de datos
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}

// LoginUser godoc
// @Summary Iniciar sesión
// @Description Autentica a un usuario y devuelve un token JWT
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Credenciales (email y contraseña)"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /users/login [post]
func LoginUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Buscar el usuario por email
		var user User
		if err := db.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
			return
		}

		// Verificar la contraseña
		if err := utils.CheckPassword(credentials.Password, user.PasswordHash); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
			return
		}

		// Generar un token JWT (opcional)
		// token, err := generateJWT(user.ID)
		// if err != nil {
		//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token"})
		//     return
		// }

		c.JSON(http.StatusOK, gin.H{
			"message": "Inicio de sesión exitoso",
			// "token":   token,
		})
	}
}

// CreateTask godoc
// @Summary Crear una nueva tarea
// @Description Crea una nueva tarea en el sistema
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body CreateTaskRequest true "Datos de la tarea"
// @Success 201 {object} CreateTaskRequest
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /createTask [post]
func CreateTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request CreateTaskRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dueDate, err := time.Parse(time.RFC3339, request.DueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due date format"})
			return
		}

		task := Task{
			Title:       request.Title,
			Description: request.Description,
			Status:      request.Status,
			Priority:    request.Priority,
			CreatedBy:   request.CreatedBy,
			AssignedTo:  request.AssignedTo,
			DueDate:     dueDate,
		}

		if err := db.Create(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, task)
	}
}

// UpdateTask godoc
// @Summary Actualizar una tarea
// @Description Actualiza los datos de una tarea existente
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID de la tarea"
// @Param task body Task true "Datos actualizados de la tarea"
// @Success 200 {object} Task
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /updateTasks/{id} [put]
func UpdateTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var task Task
		if err := db.First(&task, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
			return
		}

		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Save(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, task)
	}
}

// CreateComment godoc
// @Summary Crear un nuevo comentario
// @Description Crea un nuevo comentario en una tarea
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID de la tarea"
// @Param comment body Comment true "Datos del comentario"
// @Success 201 {object} Comment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id}/comments [post]
func CreateComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de tarea inválido"})
			return
		}

		var comment Comment
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		comment.TaskID = uint(taskID)

		if err := db.Create(&comment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, comment)
	}
}

// GetUsers godoc
// @Summary Obtener todos los usuarios
// @Description Obtiene una lista de todos los usuarios
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Failure 500 {object} map[string]string
// @Router /getAllUsers [get]
func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// GetCommentsByTask godoc
// @Summary Obtener comentarios por ID de tarea
// @Description Obtiene una lista de comentarios asociados a una tarea específica
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID de la tarea"
// @Success 200 {array} Comment
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /tasks/{id}/getComments [get]
func GetCommentsByTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de tarea inválido"})
			return
		}

		var comments []Comment
		if err := db.Where("task_id = ?", taskID).Find(&comments).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comentarios no encontrados"})
			return
		}

		c.JSON(http.StatusOK, comments)
	}
}
