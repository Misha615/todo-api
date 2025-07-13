package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func InitSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./todo.db")
	if err != nil {
		return nil, err
	}

	return db, err
}

func InitSchema(db *sql.DB) error {
	if err := createUsersTable(db); err != nil {
		return fmt.Errorf("users table creation failed: %w", err)
	}

	if err := createTasksTable(db); err != nil {
		return fmt.Errorf("task table creation failed: %w", err)
	}

	return nil
}

func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS users (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				username TEXT NOT NULL
			);
	`)
	return err
}

func createTasksTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT DEFAULT 'todo',
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
			CHECK(status IN ('todo', 'in_progress', 'done', 'archived'))
		);
	`)
	return err
}
