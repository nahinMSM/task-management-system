package domain

import "context"

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	GetAll(ctx context.Context) ([]Task, error)
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id int64) error
}
