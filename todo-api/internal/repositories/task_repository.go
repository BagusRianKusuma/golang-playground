package repositories

import (
	"database/sql"
	"todo-api/internal/entities"
)

type TaskRepository struct {
	db *sql.DB //menyimpan koneksi db
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db} // Mengembalikan objek TaskRepository yang siap dipakai.
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
