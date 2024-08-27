package models

import (
	"database/sql"
	"time"
)

type File struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Name       string    `json:"name"`
	Size       int64     `json:"size"`
	UploadDate time.Time `json:"upload_date"`
	URL        string    `json:"url"`
}

func (f *File) Create(db *sql.DB) error {
	query := `INSERT INTO files (user_id, name, size, upload_date, url) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return db.QueryRow(query, f.UserID, f.Name, f.Size, time.Now(), f.URL).Scan(&f.ID)
}

func GetFilesByUserID(db *sql.DB, userID int) ([]File, error) {
	query := `SELECT id, user_id, name, size, upload_date, url FROM files WHERE user_id = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []File
	for rows.Next() {
		var file File
		err := rows.Scan(&file.ID, &file.UserID, &file.Name, &file.Size, &file.UploadDate, &file.URL)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}