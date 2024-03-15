package config

import (
    "net/http"

    httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "go-todo-api/docs"
)

func SetupSwagger() http.Handler {
    return httpSwagger.Handler(
        httpSwagger.URL("http://localhost:8080/todo/v1/swagger/doc.json"), // URL para acessar a documentação do Swagger
    )
}