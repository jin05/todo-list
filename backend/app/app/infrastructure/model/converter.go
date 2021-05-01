package model

import "todo-list/app/domain"

func ToUserDomain(user *User) *domain.User {
	return &domain.User{
		ID:       user.ID,
		UserName: user.UserName,
		AuthID:   user.AuthID,
		Email:    user.Email,
	}
}
