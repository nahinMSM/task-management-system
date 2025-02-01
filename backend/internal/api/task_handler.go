package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"backend/internal/domain"
	"backend/internal/usecase"
)

type TaskHandler struct {
	UseCase *usecase.TaskUseCase
}

// Listar tarefas com filtro
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")             // Pega o filtro da URL
	tasks, err := h.UseCase.GetAll(r.Context()) // Usa contexto para buscar todas as tarefas
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filtro opcional
	if query != "" {
		filteredTasks := []domain.Task{}
		for _, task := range tasks {
			if contains(task.Title, query) || contains(task.Description, query) {
				filteredTasks = append(filteredTasks, task)
			}
		}
		tasks = filteredTasks
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Criar uma nova tarefa
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.UseCase.Create(r.Context(), &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": int64(task.ID)})
}

// Atualizar uma tarefa
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	task.ID = int(id)
	err = h.UseCase.Update(r.Context(), &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Excluir uma tarefa
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.UseCase.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// contains verifica se o texto contém um substrato (case insensitive)
func contains(text, substr string) bool {
	return strings.Contains(strings.ToLower(text), strings.ToLower(substr))
}
