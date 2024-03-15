# Aplicação de TO-DO em GoLang com MySQL e Swagger

Esta é uma aplicação de TO-DO simples desenvolvida em GoLang, utilizando MySQL como banco de dados e Swagger para documentação da API. A aplicação oferece operações CRUD (Create, Read, Update, Delete) para gerenciar tarefas em uma lista de afazeres.

## Pré-requisitos

- GoLang instalado na máquina local
- MySQL instalado e configurado
- Pacote Swagger (`github.com/swaggo/swag`) instalado

## Como executar a aplicação

1. Clone o repositório para a sua máquina local:

   ```sh
   git clone https://github.com/leonardosugahara/go-todo-api

2. Acesse o diretório da aplicação:

   ```sh
   cd go-todo-api

3. Execute o comando para gerar o arquivo de documentação do Swagger:

   ```sh
   swag init

4. Execute o comando para rodar a aplicação:

   ```sh
   go run main.go

5. Acesse a documentação da API Swagger em http://localhost:8080/todo/v1/swagger/index.html

## Endpoints da API
- POST /tasks: Cria uma nova tarefa. Corpo da requisição JSON com os campos title e description.
- PUT /tasks/{id}: Atualiza uma tarefa existente com o ID fornecido. Corpo da requisição JSON com os campos title e description.
- DELETE /tasks/{id}: Exclui uma tarefa existente com o ID fornecido.
- GET /tasks/{id}: Retorna os detalhes de uma tarefa com o ID fornecido.
- GET /tasks: Retorna todas as tarefas cadastradas.

## Referencia
O passo a passo você pode encontrar no post 
https://leonardosugahara.blogspot.com/2024/03/criando-uma-api-em-go-com-swagger-e.html