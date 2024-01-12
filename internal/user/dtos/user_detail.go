package dtos

import (
	"myapp/internal/user/entity"
	"time"
)

type UserDetailResponse struct {
	UserID      int       `json:"id"`
	Email       string    `json:"email"`
	Fullname    string    `json:"fullname"`
	PhoneNumber string    `json:"phone_number"`
	UserType    string    `json:"user_type"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int       `json:"created_by"`
}

func NewUserDetail(data entity.User) UserDetailResponse {
	return UserDetailResponse{
		UserID:      data.UserID,
		Email:       data.Email,
		Fullname:    data.Fullname,
		PhoneNumber: data.PhoneNumber,
		UserType:    data.UserType,
		IsActive:    data.IsActive,
		CreatedAt:   data.CreatedAt,
		CreatedBy:   data.CreatedBy,
	}
}
