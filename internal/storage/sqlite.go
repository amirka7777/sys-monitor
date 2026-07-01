package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func NewSqliteStorage(path string) (*Storage, error) {

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть базу данных: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка при попытке связаться с базой данных: %v", err)
	}

	s := &Storage{db: db}

	if err := s.createTable(); err != nil {
		return nil, err
	}

	return s, nil

}

func (s *Storage) createTable() error {

	query := `CREATE TABLE IF NOT EXISTS metrics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		server_id TEXT NOT NULL,
		cpu_usage REAL NOT NULL,
		free_disk_space INTEGER NOT NULL,
		timestamp INTEGER NOT NULL);`

	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("Ошибка при создании Базы данных: %v", err)
	}

	return nil
}

func (s *Storage) SaveMetrics(serverID string, cpu float32, disk uint64, timestamp int64) error {

	query := `
	INSERT INTO metrics (server_id, cpu_usage, free_disk_space, timestamp) VALUES (?, ?, ?, ?)
	`

	if _, err := s.db.Exec(query, serverID, cpu, disk, timestamp); err != nil {
		return fmt.Errorf("Ошибка при вставки данных в базу данных: %v", err)
	}

	return nil

}

func (s *Storage) Close() error {
	return s.db.Close()
}
