package usecase

import (
	"errors"
	"gorm.io/gorm"
	"todo-list/app/domain"
	"todo-list/app/infrastructure/database"
	"todo-list/app/infrastructure/repository"
)

type UserUseCase interface {
	GetByAuthID(authID string) (*domain.User, error)
	Save(authID string, name string, email string) (*domain.User, error)
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

func (u *userUseCase) Save(authID string, name string, email string) (user *domain.User, err error) {
	user, err = u.userRepository.GetByAuthID(authID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if user != nil {
		return
	}

	err = u.conn.DB.Transaction(func(tx *gorm.DB) error {
		user, err = u.userRepository.Save(authID, name, email)
		return err
	})
	return
}
