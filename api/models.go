package api

import (
	"time"
)

type User struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`          // ID único del usuario
	Name          string    `gorm:"type:varchar(100);not null"`        // Nombre del usuario
	Email         string    `gorm:"type:varchar(100);unique;not null"` // Correo electrónico (único)
	PasswordHash  string    `gorm:"type:text;not null"`                // Hash de la contraseña
	CreatedAt     time.Time `gorm:"type:timestamp;default:now()"`      // Fecha de creación
	UpdatedAt     time.Time `gorm:"type:timestamp;default:now()"`      // Fecha de última actualización
	Tasks         []Task    `gorm:"foreignKey:CreatedBy"`              // Tareas creadas por el usuario
	AssignedTasks []Task    `gorm:"foreignKey:AssignedTo"`             // Tareas asignadas al usuario
	Comments      []Comment `gorm:"foreignKey:UserID"`                 // Comentarios hechos por el usuario
}

type Task struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`           // ID único de la tarea
	Title       string    `gorm:"type:varchar(200);not null"`         // Título de la tarea
	Description string    `gorm:"type:text"`                          // Descripción detallada
	Status      string    `gorm:"type:varchar(50);default:'pending'"` // Estado (pending, in_progress, completed)
	Priority    string    `gorm:"type:varchar(50);default:'medium'"`  // Prioridad (low, medium, high)
	CreatedBy   uint      `gorm:"not null"`                           // Usuario que creó la tarea
	AssignedTo  *uint     // Usuario asignado
	CreatedAt   time.Time `gorm:"type:timestamp;default:now()"` // Fecha de creación
	UpdatedAt   time.Time `gorm:"type:timestamp;default:now()"` // Fecha de última actualización
	DueDate     time.Time // Fecha límite de la tarea
	Comments    []Comment `gorm:"foreignKey:TaskID"` // Comentarios asociados a la tarea
}

type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`     // ID único del comentario
	TaskID    uint      `gorm:"not null"`                     // Tarea asociada
	UserID    uint      `gorm:"not null"`                     // Usuario que hizo el comentario
	Content   string    `gorm:"type:text;not null"`           // Contenido del comentario
	CreatedAt time.Time `gorm:"type:timestamp;default:now()"` // Fecha de creación
}

type Activity struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`     // ID único de la actividad
	TaskID        uint      `gorm:"not null"`                     // Tarea asociada
	UserID        uint      `gorm:"not null"`                     // Usuario que realizó la acción
	ActionType    string    `gorm:"type:varchar(100);not null"`   // Tipo de acción
	ActionDetails string    `gorm:"type:text"`                    // Detalles de la acción
	CreatedAt     time.Time `gorm:"type:timestamp;default:now()"` // Fecha de la actividad
}

type Notification struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`     // ID único de la notificación
	UserID    uint      `gorm:"not null"`                     // Usuario destinatario
	Message   string    `gorm:"type:text;not null"`           // Mensaje de la notificación
	IsRead    bool      `gorm:"default:false"`                // Estado de lectura
	CreatedAt time.Time `gorm:"type:timestamp;default:now()"` // Fecha de creación
}
