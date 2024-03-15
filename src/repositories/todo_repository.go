package repositories

import (
    "database/sql"
    "log"

    "go-todo-api/src/models"
)

type TodoRepository struct {
    db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
    return TodoRepository{
        db: db,
    }
}

func (r *TodoRepository) Create(task *models.Task) error {
    query := "INSERT INTO tasks (title, description) VALUES (?, ?)"
    _, err := r.db.Exec(query, task.Title, task.Description)
    if err != nil {
        log.Printf("Erro ao inserir tarefa no banco de dados: %v", err)
        return err
    }
    return nil
}

func (r *TodoRepository) Update(id int, task *models.Task) error {
    query := "UPDATE tasks SET title = ?, description = ? WHERE id = ?"
    _, err := r.db.Exec(query, task.Title, task.Description, id)
    if err != nil {
        log.Printf("Erro ao atualizar tarefa no banco de dados: %v", err)
        return err
    }
    return nil
}

func (r *TodoRepository) Delete(id int) error {
    query := "DELETE FROM tasks WHERE id = ?"
    _, err := r.db.Exec(query, id)
    if err != nil {
        log.Printf("Erro ao excluir tarefa no banco de dados: %v", err)
        return err
    }
    return nil
}

func (r *TodoRepository) GetByID(id int) (*models.Task, error) {
    query := "SELECT id, title, description FROM tasks WHERE id = ?"
    row := r.db.QueryRow(query, id)

    var task models.Task
    err := row.Scan(&task.ID, &task.Title, &task.Description)
    if err != nil {
        log.Printf("Erro ao obter tarefa do banco de dados: %v", err)
        return nil, err
    }
    return &task, nil
}

func (r *TodoRepository) GetAll() ([]*models.Task, error) {
    query := "SELECT id, title, description FROM tasks"
    rows, err := r.db.Query(query)
    if err != nil {
        log.Printf("Erro ao obter tarefas do banco de dados: %v", err)
        return nil, err
    }
    defer rows.Close()

    var tasks []*models.Task
    for rows.Next() {
        var task models.Task
        err := rows.Scan(&task.ID, &task.Title, &task.Description)
        if err != nil {
            log.Printf("Erro ao ler linha de tarefa do banco de dados: %v", err)
            return nil, err
        }
        tasks = append(tasks, &task)
    }

    return tasks, nil
}