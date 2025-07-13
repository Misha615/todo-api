package sqlite

import (
	"database/sql"
	"fmt"
	"todo-api/internal/domain/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	res, err := r.db.Exec("INSERT INTO users (username) VALUES (?)", user.Username)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	user.ID = int(id)
	return nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	row := r.db.QueryRow("SELECT id, username FROM users WHERE id = ?", id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	_, err := r.db.Exec(
		`UPDATE users 
		SET username = ?
		WHERE id = ?`,
		user.Username,
		user.ID,
	)
	return err
}

func (r *UserRepository) Delete(user *models.User) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", user.ID)
	return err
}
