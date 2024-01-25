package globalstruct

import "time"

type Tracing struct {
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy int        `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updated_by"`
}
