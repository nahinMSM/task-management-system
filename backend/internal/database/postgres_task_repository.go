package database

import (
	"backend/internal/domain"
	"context"
	"database/sql"
)

type PostgresTaskRepository struct {
	DB *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{DB: db}
}

func (r *PostgresTaskRepository) Create(ctx context.Context, task *domain.Task) error {
	query := "INSERT INTO tasks (title, description, completed) VALUES ($1, $2, $3)"
	_, err := r.DB.ExecContext(ctx, query, task.Title, task.Description, task.Completed)
	return err
}

func (r *PostgresTaskRepository) GetAll(ctx context.Context) ([]domain.Task, error) {
	query := "SELECT id, title, description, completed FROM tasks"
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *PostgresTaskRepository) Update(ctx context.Context, task *domain.Task) error {
	query := "UPDATE tasks SET title = $1, description = $2, completed = $3 WHERE id = $4"
	_, err := r.DB.ExecContext(ctx, query, task.Title, task.Description, task.Completed, task.ID)
	return err
}

func (r *PostgresTaskRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}
