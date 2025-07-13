package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required,min=3"`
}

type Task struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"oneof=todo in_progress done"`
}
