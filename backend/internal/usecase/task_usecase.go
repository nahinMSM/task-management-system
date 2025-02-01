package usecase

import (
	"backend/internal/domain"
	"context"
	"errors"
)

type TaskUseCase struct {
	TaskRepo domain.TaskRepository
}

func NewTaskUseCase(taskRepo domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{TaskRepo: taskRepo}
}

func (u *TaskUseCase) Create(ctx context.Context, task *domain.Task) error {
	if task.Title == "" {
		return errors.New("o título da tarefa não pode ser vazio")
	}
	return u.TaskRepo.Create(ctx, task)
}

func (u *TaskUseCase) GetAll(ctx context.Context) ([]domain.Task, error) {
	return u.TaskRepo.GetAll(ctx)
}

func (u *TaskUseCase) Update(ctx context.Context, task *domain.Task) error {
	if task.ID == 0 {
		return errors.New("o ID da tarefa é obrigatório")
	}
	return u.TaskRepo.Update(ctx, task)
}

func (u *TaskUseCase) Delete(ctx context.Context, id int64) error {
	return u.TaskRepo.Delete(ctx, id)
}
