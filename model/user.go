package model

import (
	"github.com/google/uuid"
)

// Una estructura es como una clase en otros lenguajes de programación. Define un tipo de dato compuesto que puede contener múltiples campos.
type User struct {
	// ¿qué es UUID? UUID es un tipo de dato que representa un identificador único universal (Universally Unique Identifier).
	ID        uuid.UUID `json:"id"`         // User ID de tipo UUID
	Email     string    `json:"email"`      // Email del usuario
	Password  string    `json:"password"`   // Password del usuario
	IsAdmin   bool      `json:"is_admin"`   // Indica si el usuario es administrador
	Details   string    `json:"details"`    // Detalles adicionales del usuario
	CreatedAt int64     `json:"created_at"` // Fecha de creación del usuario
	UpdatedAt int64     `json:"updated_at"` // Fecha de actualización del usuario
}

type Users []User //Creo un Alias o atajo para un array dinamico
