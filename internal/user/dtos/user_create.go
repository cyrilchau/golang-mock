package dtos

import (
	"myapp/config"
	"myapp/internal/user/entity"
	"myapp/pkg/constant"
	"myapp/pkg/utils"
	"time"
)

type CreateUserRequest struct {
	FullName    string `json:"fullname" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	UserID      int    `json:"user_id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func NewCreateUser(data CreateUserRequest, cfg config.Config) entity.User {
	return entity.User{
		Fullname:    data.FullName,
		Email:       data.Email,
		Password:    utils.Encrypt(data.Password, cfg),
		PhoneNumber: data.PhoneNumber,
		UserType:    constant.UserTypeRegular,
		IsActive:    true,
		CreatedAt:   time.Now(),
		CreatedBy:   9999,
	}
}
