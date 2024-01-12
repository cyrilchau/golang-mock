package dtos

type CreateWorkerRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Gender      string `json:"gender"`
	CreatedBy   int    `json:"created_by" validate:"required"`
}
