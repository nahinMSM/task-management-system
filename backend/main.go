package main

import (
	"backend/internal/api"
	"backend/internal/database"
	"backend/internal/usecase"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=tasks password=1234 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := database.NewPostgresTaskRepository(db)

	useCase := usecase.NewTaskUseCase(repository)

	taskHandler := &api.TaskHandler{
		UseCase: useCase,
	}

	http.HandleFunc("/tasks", taskHandler.GetTasks)
	http.HandleFunc("/task", taskHandler.CreateTask)
	http.HandleFunc("/task/update", taskHandler.UpdateTask)
	http.HandleFunc("/task/delete", taskHandler.DeleteTask)

	log.Println("Servidor iniciado na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
