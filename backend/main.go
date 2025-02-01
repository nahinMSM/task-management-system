package main

import (
	"backend/internal/api"
	"backend/internal/database"
	"backend/internal/usecase"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Pacote para PostgreSQL
)

func main() {
	// Conecta ao banco de dados (ajuste os parâmetros de conexão conforme necessário)
	connStr := "user=postgres dbname=tasks password=1234 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inicializa o repositório com a conexão do banco de dados
	repository := database.NewPostgresTaskRepository(db)

	// Inicializa o caso de uso com o repositório
	useCase := usecase.NewTaskUseCase(repository)

	// Inicializa o manipulador de tarefas
	taskHandler := &api.TaskHandler{
		UseCase: useCase,
	}

	// Configura as rotas do servidor HTTP
	http.HandleFunc("/tasks", taskHandler.GetTasks)
	http.HandleFunc("/task", taskHandler.CreateTask)
	http.HandleFunc("/task/update", taskHandler.UpdateTask)
	http.HandleFunc("/task/delete", taskHandler.DeleteTask)

	// Inicia o servidor HTTP
	log.Println("Servidor iniciado na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
