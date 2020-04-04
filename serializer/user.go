package serializer

import (
	"goa-shopping/gen/auth"
	"goa-shopping/model"
)

func ModelUserToResult(user model.User) *auth.User {
	resp := auth.User{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Mobile:    user.Mobile,
		Email:     user.Email,
		Type:      user.Type,
		IsActive:  user.IsActive,
		LoginTime: "",
	}
	return &resp
}
