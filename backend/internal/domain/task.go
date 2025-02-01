package domain

import "time"

// Task representa uma tarefa no sistema
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TaskRepository define os métodos para interação com o repositório de tarefas
// type TaskRepository interface {
// 	Create(task Task) (int, error)       // Criar tarefa
// 	GetAll(query string) ([]Task, error) // Listar tarefas com filtro
// 	GetByID(id int) (Task, error)        // Obter tarefa por ID
// 	Update(task Task) error              // Atualizar tarefa
// 	Delete(id int) error                 // Excluir tarefa
// }
