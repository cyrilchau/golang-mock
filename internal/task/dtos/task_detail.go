package dtos

type DetailTaskRequest struct {
	ID      uint    `json:"id" validate:"required"`
}
