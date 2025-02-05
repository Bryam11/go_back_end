package api

// RegisterUserRequest es el DTO para el registro de usuarios.
type RegisterUserRequest struct {
	Name     string `json:"name" binding:"required"`     // Nombre del usuario
	Email    string `json:"email" binding:"required"`    // Correo electrónico
	Password string `json:"password" binding:"required"` // Contraseña en texto plano
}
