package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type CreateTaskResponse struct {
	Task *Task `json:"task"`
}

type UpdateTaskRequest struct {
	Completed bool `json:"completed"`
}

type UpdateTaskResponse struct {
	Completed bool `json:"completed"`
}

type GetTasksResponse struct {
	Tasks []*Task `json:"tasks"`
}
