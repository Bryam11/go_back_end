package api

// RegisterUserRequest es el DTO para el registro de usuarios.
type RegisterUserRequest struct {
	Name     string `json:"name" binding:"required"`     // Nombre del usuario
	Email    string `json:"email" binding:"required"`    // Correo electrónico
	Password string `json:"password" binding:"required"` // Contraseña en texto plano
}

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`      // Título de la tarea
	Description string `json:"description"`                   // Descripción detallada
	Status      string `json:"status" binding:"required"`     // Estado (pending, in_progress, completed)
	Priority    string `json:"priority" binding:"required"`   // Prioridad (low, medium, high)
	CreatedBy   uint   `json:"created_by" binding:"required"` // Usuario que creó la tarea
	AssignedTo  *uint  `json:"assigned_to"`                   // Usuario asignado
	DueDate     string `json:"due_date"`                      // Fecha límite de la tarea
}
