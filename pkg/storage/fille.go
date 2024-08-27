package storage

import (
	"io"
	"os"
	"path/filepath"
)

const uploadDir = "./uploads"

func SaveFile(file io.Reader, filename string) (string, error) {
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	dst, err := os.Create(filepath.Join(uploadDir, filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return "", err
	}

	return filepath.Join(uploadDir, filename), nil
}