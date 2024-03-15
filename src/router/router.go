package router

import (
    "github.com/gorilla/mux"
    "go-todo-api/src/handlers"
)

func NewRouter(handler *handlers.TodoHandler) *mux.Router {
    r := mux.NewRouter()
    s := r.PathPrefix("/todo/v1").Subrouter()
    s.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
    s.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
    s.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
    s.HandleFunc("/tasks/{id}", handler.GetTaskByID).Methods("GET")
    s.HandleFunc("/tasks", handler.GetAllTasks).Methods("GET")
    return s
}
