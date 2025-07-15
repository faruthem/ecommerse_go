// Services
package user

import (
	"fmt"
	"time"

	"github.com/farithem/ecommerse_go/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{storage: s}
}

// Crear usuario
func (u User) Create(m *model.User) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("uuid.NewUUID(): %w", err)
	}

	m.ID = ID
	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt.GenerateFromPassword(): %w", err)
	}
	m.Password = string(password)
	if m.Details == "" {
		m.Details = "{}"
	}
	m.CreatedAt = time.Now().Unix()
	err = u.storage.Create(m)
	if err != nil {
		return fmt.Errorf("storage.Create(): %w", err)
	}
	m.Password = ""
	return nil
}
