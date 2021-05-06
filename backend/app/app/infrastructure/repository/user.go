package repository

import (
	"todo-list/app/domain"
	"todo-list/app/infrastructure/database"
	"todo-list/app/infrastructure/model"
)

type UserRepository interface {
	GetByAuthID(authID string) (*domain.User, error)
	Save(authID string, name string, email string) (*domain.User, error)
}

type userRepository struct {
	conn *database.Connection
}

func NewUserRepository(conn *database.Connection) UserRepository {
	return &userRepository{conn: conn}
}

func (r *userRepository) GetByAuthID(authID string) (*domain.User, error) {
	conn := r.conn.DB
	user := &model.User{}
	result := conn.Where(&model.User{AuthID: authID}).First(user)
	return model.ToUserDomain(user), result.Error
}

func (r *userRepository) Save(authID string, name string, email string) (*domain.User, error) {
	conn := r.conn.DB
	user := &model.User{
		UserName: name,
		AuthID:   authID,
		Email:    email,
	}
	err := conn.Create(user).Error
	if err != nil {
		return nil, err
	}
	return model.ToUserDomain(user), nil
}
