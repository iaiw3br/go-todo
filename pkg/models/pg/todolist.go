package pg

import (
	"database/sql"
	"go-todo/pkg/models"
)

type TodoModel struct {
	DB *sql.DB
}

func (t *TodoModel) GetAll() ([]*models.TodoList, error) {
	sqlQuery := `
		SELECT id, title, isCompleted, created 
		FROM public.todo`

	rows, err := t.DB.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	var todos []*models.TodoList

	for rows.Next() {
		todo := &models.TodoList{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.IsCompleted, &todo.Created)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *TodoModel) Create(title string, isCompleted bool) (int, error) {
	sqlQuery := `
		INSERT INTO public.todo (title, isCompleted) 
		VALUES($1, $2) 
		returning id`
	var id int
	err := t.DB.QueryRow(sqlQuery, title, isCompleted).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
