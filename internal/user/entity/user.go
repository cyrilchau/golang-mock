package entity

import (
	"time"
)

type User struct {
	UserID      int       `gorm:"primarykey"`
	Email       string    `gorm:"column:email;uniqueIndex"`
	Password    string    `gorm:"column:password"`
	Fullname    string    `gorm:"column:fullname"`
	PhoneNumber string    `gorm:"column:phone_number"`
	UserType    string    `gorm:"column:user_type"`
	IsActive    bool      `gorm:"column:is_active"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	CreatedBy   int       `gorm:"column:created_by"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	UpdatedBy   string    `gorm:"column:updated_by"`
}

func (u *User) TableName() string {
	return "user"
}

func (m *User) IsExists() (ok bool) {
	if m.UserID != 0 {
		ok = true
	}
	return
}
