package repository

import (
	"fmt"
	"todo-list/app/domain"
	"todo-list/app/infrastructure/database"
	"todo-list/app/infrastructure/model"
)

type TodoRepository interface {
	Get(userID int64, todoID int64) (*domain.Todo, error)
	List(userID int64, keyWards []string) ([]*domain.Todo, error)
	Save(userID int64, title string, content string) (*domain.Todo, error)
	Update(userID int64, todoID int64, title string, content string, check bool) error
	Delete(userID int64, todoID int64) error
}

type todoRepository struct {
	conn *database.Connection
}

func NewTodoRepository(conn *database.Connection) TodoRepository {
	return &todoRepository{conn: conn}
}

func (r *todoRepository) Get(userID int64, todoID int64) (*domain.Todo, error) {
	conn := r.conn.DB
	todo := &model.Todo{ID: todoID, UserID: userID}
	err := conn.First(todo).Error
	if err != nil {
		return nil, err
	}
	return model.ToTodoDomain(todo), nil
}

func (r *todoRepository) List(userID int64, keyWards []string) ([]*domain.Todo, error) {
	conn := r.conn.DB
	var todoList []*model.Todo
	where := conn.Where(&model.Todo{UserID: userID})

	if 0 < len(keyWards) {
		for _, keyWard := range keyWards {
			if keyWard != "" {
				where = where.Where("content LIKE ?", fmt.Sprintf("%%%s%%", keyWard))
			}
		}
	}

	err := where.Find(todoList).Error
	if err != nil {
		return nil, err
	}
	return model.ToTodoDomainList(todoList), nil
}

func (r *todoRepository) Save(userID int64, title string, content string) (*domain.Todo, error) {
	conn := r.conn.DB
	todo := &model.Todo{
		UserID:  userID,
		Title:   title,
		Content: content,
	}
	err := conn.Create(todo).Error
	if err != nil {
		return nil, err
	}
	return model.ToTodoDomain(todo), err
}

func (r *todoRepository) Update(userID int64, todoID int64, title string, content string, check bool) error {
	conn := r.conn.DB
	todo := &model.Todo{
		ID:      todoID,
		UserID:  userID,
		Title:   title,
		Content: content,
		Check:   check,
	}
	return conn.Save(todo).Error
}

func (r *todoRepository) Delete(userID int64, todoID int64) error {
	conn := r.conn.DB
	return conn.Delete(&model.Todo{ID: todoID, UserID: userID}).Error
}
