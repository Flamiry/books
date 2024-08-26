package models

import "github.com/go-playground/validator/v10"

/*const (
	TaskStatusNew = "new"
	TaskStatusInprogress = "in_progress"
	TaskStatusCompleted = "completed"

)*/
type Task struct {
	TID string `json:"id"`
	Title string `json: "title" validate:"required"`
	Description string `json: "description" validate:"required"`
	Status string `json: "status" validate:"required"`
}