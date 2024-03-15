package main

import (
    "log"
    "net/http"

	db "go-todo-api/src/config"
	swagger "go-todo-api/src/config"
    "go-todo-api/src/handlers"
    "go-todo-api/src/repositories"
    "go-todo-api/src/router"

)

// @title           TODO API
// @version         1.0
// @description     This is a sample todo api.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /todo/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Inicializa conexão com o banco de dados
    db.InitDB()
    defer db.CloseDB()

    // Configuração do repositório
    repo := repositories.NewTodoRepository(db.DB)

    // Configuração do handler
    handler := &handlers.TodoHandler{
        Repository: repo,
    }

    // Configuração do roteamento
    r := router.NewRouter(handler)

    // Configuração do Swagger
	r.PathPrefix("/swagger/").Handler(swagger.SetupSwagger()).Methods(http.MethodGet)


    // Inicialização do servidor
    log.Println("Servidor rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}