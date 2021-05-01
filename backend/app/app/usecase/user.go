package usecase

import (
	"gorm.io/gorm"
	"todo-list/app/domain"
	"todo-list/app/domain/repository"
	"todo-list/app/infrastructure/database"
)

type UserUseCase interface {
	GetByAuthID(authID string) (*domain.User, error)
	Create(authID string, name string, email string) (*domain.User, error)
}

type userUseCase struct {
	conn           *database.Connection
	userRepository repository.UserRepository
}

func NewUserUseCase(conn *database.Connection, userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{conn: conn, userRepository: userRepository}
}

func (u *userUseCase) GetByAuthID(authID string) (*domain.User, error) {
	return u.userRepository.GetByAuthID(authID)
}

func (u *userUseCase) Create(authID string, name string, email string) (user *domain.User, err error) {
	err = u.conn.DB.Transaction(func(tx *gorm.DB) error {
		user, err = u.userRepository.Create(authID, name, email)
		return err
	})
	return
}
