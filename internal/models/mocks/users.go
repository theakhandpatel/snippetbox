package mocks

import (
	"fmt"
	"time"

	"github.com/theakhandpatel/snippetbox/internal/models"
)

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}
func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "alice@example.com" && password == "pa$$word" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}
func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) GetUserByID(id int) (*models.User, error) {
	if id == 1 {
		return &models.User{
			Email:   "alice@example.com",
			Created: time.Now(),
			Name:    "Alice",
		}, nil
	}

	return nil, fmt.Errorf("invalid ID")
}

func (m *UserModel) PasswordUpdate(id int, currentPassword, newPassword string) error {
	if id == 1 {
		if currentPassword != "pa$$word" {
			return models.ErrInvalidCredentials
		}
		return nil
	}
	return models.ErrNoRecord
}
