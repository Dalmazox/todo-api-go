package dto

type CreateTodoRequest struct {
	Description string `json:"description" example:"Lorem ipsum"`
	Done        bool   `json:"done" example:"false"`
}
