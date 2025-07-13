package sqlite

import (
	"database/sql"
	"fmt"
	"todo-api/internal/domain/models"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) Create(task *models.Task) error {
	res, err := r.db.Exec(
		"INSERT INTO tasks (user_id, title, description, status) VALUES (?, ?, ?, ?)",
		task.UserID, task.Title, task.Description, task.Status,
	)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	task.ID = int(id)
	return nil
}

func (r *TaskRepo) GetAll() ([]*models.Task, error) {
	res, err := r.db.Query("SELECT id, user_id, title, description, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var tasks []*models.Task
	for res.Next() {
		var task models.Task
		if err := res.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *TaskRepo) GetByUserID(id int) ([]*models.Task, error) {
	rows, err := r.db.Query(`
		SELECT id, user_id, title, description, status
		FROM tasks
		WHERE user_id = ?`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *TaskRepo) Update(task *models.Task) error {
	if task.ID <= 0 {
		return fmt.Errorf("invalid ID")
	}
	_, err := r.db.Exec(
		`UPDATE tasks
	     SET title = ?, description = ?, status = ?
		 WHERE id = ? AND user_id = ?
		`,
		task.Title,
		task.Description,
		task.Status,
		task.ID,
		task.UserID,
	)
	return err
}

func (r *TaskRepo) Delete(task *models.Task) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?",
		task.ID,
	)
	return err
}
