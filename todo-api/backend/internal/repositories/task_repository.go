package repositories

import (
	"database/sql"
	"todo-api/internal/entities"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *entities.Task) error {
	query := `INSERT INTO tasks(title, description,completed) VALUES($1, $2, $3) RETURNING id`
	return r.db.QueryRow(query, task.Title, task.Description, task.Completed).Scan(&task.ID)
}

func (r *TaskRepository) GetTask() ([]entities.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entities.Task
	for rows.Next() {
		var task entities.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) UpdateTask(id int, task *entities.Task) error {
	query := `UPDATE tasks SET  title=$1, description=$2, completed=$3 WHERE id=$4`
	_, err := r.db.Exec(query, task.Title, task.Description, task.Completed, id)
	return err
}

func (r *TaskRepository) DeleteTask(id int) error {
	query := `DELETE from tasks WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}
