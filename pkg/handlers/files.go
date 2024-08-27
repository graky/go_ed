package handlers

import (
	"database/sql"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go_ed/pkg/models"
	"go_ed/pkg/storage"
)

func Upload(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")

		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
			return
		}
		defer file.Close()

		if header.Size > 10*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 10MB limit"})
			return
		}

		ext := filepath.Ext(header.Filename)
		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Only PNG and JPEG files are allowed"})
			return
		}

		filename := uuid.New().String() + ext
		fileURL, err := storage.SaveFile(file, filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		fileModel := models.File{
			UserID: userID,
			Name:   header.Filename,
			Size:   header.Size,
			URL:    fileURL,
		}

		err = fileModel.Create(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file information"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
	}
}

func GetFiles(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")

		files, err := models.GetFilesByUserID(db, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
			return
		}

		c.JSON(http.StatusOK, files)
	}
}