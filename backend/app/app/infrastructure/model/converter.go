package model

import "todo-list/app/domain"

func ToUserDomain(user *User) *domain.User {
	return &domain.User{
		UserID:   user.ID,
		UserName: user.UserName,
		AuthID:   user.AuthID,
		Email:    user.Email,
	}
}

func ToTodoDomain(todo *Todo) *domain.Todo {
	return &domain.Todo{
		TodoID:  todo.ID,
		UserID:  todo.UserID,
		Title:   todo.Title,
		Content: todo.Content,
		Checked: todo.Checked,
	}
}

func ToTodoDomainList(todoList []*Todo) []*domain.Todo {
	var dTodoList []*domain.Todo
	for _, todo := range todoList {
		dTodoList = append(dTodoList, ToTodoDomain(todo))
	}
	return dTodoList
}
