package repository

import (
	"todo-list/app/domain"
	"todo-list/app/infrastructure/database"
	"todo-list/app/infrastructure/model"
)

type UserRepository interface {
	GetByAuthID(authID string) (*domain.User, error)
	Create(authID string, name string, email string) (*domain.User, error)
}

type userRepository struct {
	conn *database.Connection
}

func NewUserRepository(conn *database.Connection) UserRepository {
	return &userRepository{conn: conn}
}

func (r *userRepository) GetByAuthID(authID string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepository) Create(authID string, name string, email string) (*domain.User, error) {
	user := &model.User{
		UserName: name,
		AuthID:   authID,
		Email:    email,
	}
	err := r.conn.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return model.ToUserDomain(user), nil
}
