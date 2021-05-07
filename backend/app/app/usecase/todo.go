package usecase

import (
	"strings"
	"todo-list/app/domain"
	"todo-list/app/infrastructure/database"
	"todo-list/app/infrastructure/repository"
)

type TodoUseCase interface {
	Get(authID string, todoID int64) (*domain.Todo, error)
	Create(authID string, title string, content string) (*domain.Todo, error)
	Update(authID string, todoID int64, title string, content string, checked bool) (*domain.Todo, error)
	Delete(authID string, todoID int64) error
	List(authID string, keyWard string) ([]*domain.Todo, error)
}

type todoUseCase struct {
	conn           *database.Connection
	todoRepository repository.TodoRepository
	userRepository repository.UserRepository
}

func NewTodoUseCase(
	conn *database.Connection,
	todoRepository repository.TodoRepository,
	userRepository repository.UserRepository,
) TodoUseCase {
	return &todoUseCase{
		conn:           conn,
		todoRepository: todoRepository,
		userRepository: userRepository,
	}
}

func (u *todoUseCase) Get(authID string, todoID int64) (*domain.Todo, error) {
	user, err := u.userRepository.GetByAuthID(authID)
	if err != nil {
		return nil, err
	}
	return u.todoRepository.Get(user.UserID, todoID)
}

func (u *todoUseCase) Create(authID string, title string, content string) (*domain.Todo, error) {
	user, err := u.userRepository.GetByAuthID(authID)
	if err != nil {
		return nil, err
	}
	return u.todoRepository.Save(user.UserID, title, content)
}

func (u *todoUseCase) Update(authID string, todoID int64, title string, content string, checked bool) (*domain.Todo, error) {
	user, err := u.userRepository.GetByAuthID(authID)
	if err != nil {
		return nil, err
	}

	todo, err := u.todoRepository.Get(user.UserID, todoID)
	if err != nil {
		return nil, err
	}

	if err = u.todoRepository.Update(user.UserID, todoID, title, content, checked); err != nil {
		return nil, err
	}

	todo.Title = title
	todo.Content = content
	todo.Checked = checked
	return todo, nil
}

func (u *todoUseCase) Delete(authID string, todoID int64) error {
	user, err := u.userRepository.GetByAuthID(authID)
	if err != nil {
		return err
	}
	return u.todoRepository.Delete(user.UserID, todoID)
}

func (u *todoUseCase) List(authID string, keyWard string) ([]*domain.Todo, error) {
	user, err := u.userRepository.GetByAuthID(authID)
	if err != nil {
		return nil, err
	}

	keyWards := strings.Fields(keyWard)
	return u.todoRepository.List(user.UserID, keyWards)
}
