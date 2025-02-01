package domain

import "context"

// Task representa uma tarefa no sistema
// type Task struct {
// 	ID          int64  `json:"id"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Completed   bool   `json:"completed"`
// }

// TaskRepository define as operações para o repositório de tarefas
type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	GetAll(ctx context.Context) ([]Task, error)
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id int64) error
}
