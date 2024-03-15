package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go-todo-api/src/models"
	"go-todo-api/src/repositories"
)

type TodoHandler struct {
	Repository repositories.TodoRepository
}

// @Summary      Create a task
// @Description  create a task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task   body   models.Task true  "Task"
// @Success      201  {object}  models.Task
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /tasks [post]
func (h *TodoHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	err = h.Repository.Create(&task)
	if err != nil {
		http.Error(w, "Erro ao criar tarefa", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// @Summary      Update a task
// @Description  update a task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Param        task   body   models.Task true  "Task"
// @Success      200  {object}  models.Task
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /tasks/{id} [put]
func (h *TodoHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var task models.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	err = h.Repository.Update(id, &task)
	if err != nil {
		http.Error(w, "Erro ao atualizar tarefa", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// @Summary      Delete a task
// @Description  delete a task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      204
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /tasks/{id} [delete]
func (h *TodoHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.Repository.Delete(id)
	if err != nil {
		http.Error(w, "Erro ao excluir tarefa", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary      Show a task
// @Description  get task by ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  models.Task
// @Failure      404  {object}  string
// @Router       /tasks/{id} [get]
func (h *TodoHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	task, err := h.Repository.GetByID(id)
	if err != nil {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// @Summary 	 Get Task List
// @Description  get all tasks
// @Tags         tasks
// @Accept   	 json
// @Produce  	 json
// @Success      200  {object}  []models.Task
// @Failure      500  {object}  string
// @Router  	/tasks [get]
func (h *TodoHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Repository.GetAll()
	if err != nil {
		http.Error(w, "Erro ao obter tarefas", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}
