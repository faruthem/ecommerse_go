// interfaces
package user

import "github.com/farithem/ecommerse_go/model"

// interface sabe qué hacer no como
// Puerto de comunicación para datos usuario
type UseCase interface {
	Create(m *model.User) error                   // Crea usuario
	GetByEmail(email string) (*model.User, error) // Obtiene usuario por email
	GetAll() (model.Users, error)
}

// Datos de salida, dónde lo voy a almacenar
type Storage interface {
	Create(m *model.User) error                  // Crea usuario
	GetByEmail(email string) (model.User, error) // Obtiene usuario por email
	GetAll() (model.Users, error)
}
