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

func (t *TodoModel) Delete(id int) error {
	sqlQuery := `DELETE FROM public.todo WHERE id = $1`

	_, err := t.DB.Exec(sqlQuery, id)

	if err != nil {
		return err
	}

	return nil
}

func (t *TodoModel) Update(id int, title string) (int, string, error) {
	sqlQuery := `
		UPDATE public.todo
		SET title = $1
		WHERE id = $2
		returning id, title`
	var todoId int
	var todoTitle string
	err := t.DB.QueryRow(sqlQuery, title, id).Scan(&todoId, &todoTitle)

	if err != nil {
		return 0, "", err
	}
	return todoId, todoTitle, nil
}
