package entity

import (
	gStruct "myapp/pkg/globalStruct"
)

type TaskStatus string

const (
	DoingTaskStatus   TaskStatus = "doing"
	DoneTaskStatus    TaskStatus = "done"
	DeletedTaskStatus TaskStatus = "deleted"
)

type Task struct {
	ID          uint       `gorm:"primarykey" json:"id"`
	WorkerID    int        `gorm:"column:worker_id" json:"worker_id"`
	Title       string     `gorm:"column:title" json:"title"`
	Description string     `gorm:"column:description" json:"description"`
	Status      TaskStatus `gorm:"column:status" json:"status"`

	gStruct.Tracing
}

func (u *Task) TableName() string {
	return "task"
}

func (m *Task) IsExists() (ok bool) {
	if m.ID != 0 {
		ok = true
	}
	return
}
