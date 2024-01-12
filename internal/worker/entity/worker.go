package entity

import (
	"myapp/config"
	"myapp/internal/worker/dtos"
	gStruct "myapp/pkg/globalStruct"
	"time"
)

type WorkerGender string

const (
	GenderMale    WorkerGender = "male"
	GenderFemale  WorkerGender = "female"
	GenderUnknown WorkerGender = "unknown"
)

type WorkerStatus string

const (
	StatusActive        WorkerStatus = "active"
	StatusPendingVerify WorkerStatus = "waiting_verify"
	StatusBanned        WorkerStatus = "banned"
)

type Worker struct {
	ID          uint         `gorm:"primarykey"`
	FirstName   string       `gorm:"column:first_name"`
	LastName    string       `gorm:"column:last_name"`
	Email       string       `gorm:"column:email"`
	PhoneNumber string       `gorm:"column:phone"`
	Gender      WorkerGender `gorm:"column:gender"`
	Status      WorkerStatus `gorm:"column:status"`

	gStruct.Tracing `gorm:",inline"`
}

func (u *Worker) TableName() string {
	return "worker"
}

func NewCreateWorker(data dtos.CreateWorkerRequest, cfg config.Config) Worker {
	now := time.Now()
	return Worker{
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		Gender:      WorkerGender(data.Gender),
		Status:      StatusPendingVerify,
		Tracing: gStruct.Tracing{
			CreatedAt: &now,
			CreatedBy: data.CreatedBy,
		},
	}
}

func (m *Worker) IsExists() (ok bool) {
	if m.ID != 0 {
		ok = true
	}
	return
}
